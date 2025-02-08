package core

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"syscall"

	"github.com/arnaudlcm/container-engine/common/log"
)

const LIB_FS_MERGED_DIR = LIB_FOLDER_PATH + "/containers"
const LIB_FS_LAYERS_DIR = LIB_FOLDER_PATH + "/layers"

type FSManager struct {
	layers map[string]string
}

func NewFSManager() *FSManager {
	fsManager := &FSManager{
		layers: make(map[string]string),
	}

	fsManager.SetupFSDirs()

	return fsManager
}

// The purpose of this function is to setup the overlay FS dirs
func (f *FSManager) SetupFSDirs() {
	lowerDir := LIB_FS_LAYERS_DIR
	upperDir := path.Join(LIB_FOLDER_PATH, "volumes")
	workDir := path.Join(LIB_FOLDER_PATH, "work")
	mergedDir := LIB_FS_MERGED_DIR
	os.MkdirAll(lowerDir, 0755)
	os.MkdirAll(upperDir, 0755)
	os.MkdirAll(workDir, 0755)
	os.MkdirAll(mergedDir, 0755)

	// Mount OverlayFS
	opts := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", lowerDir, upperDir, workDir)
	err := syscall.Mount("overlay", mergedDir, "overlay", 0, opts)
	if err != nil {
		log.Fatal("Failed to mount overlay: %v", err)
	}

}

func (f *FSManager) AddLayer(layerUrl string, containerUUID string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "tmp_layer_cengine")
	if err != nil {
		return "", fmt.Errorf("error creating temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	resp, err := http.Get(layerUrl)
	if err != nil {
		return "", fmt.Errorf("failed to download tarball: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad response status: %d", resp.StatusCode)
	}

	tarballPath := filepath.Join(tmpDir, containerUUID)

	// Create file for saving
	outFile, err := os.Create(tarballPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer outFile.Close()

	// Compute checksum while writing
	hasher := sha256.New()
	multiWriter := io.MultiWriter(outFile, hasher)

	_, err = io.Copy(multiWriter, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save tarball: %w", err)
	}

	checksum := hex.EncodeToString(hasher.Sum(nil))

	layerPath := filepath.Join(LIB_FS_LAYERS_DIR, checksum)

	if _, ok := f.layers[checksum]; !ok {
		if err := extractTarball(tarballPath, layerPath); err != nil {
			return "", err
		}
		f.layers[checksum] = layerPath
	}

	return filepath.Join(LIB_FS_MERGED_DIR, checksum), nil
}

func extractTarball(tarballPath, destDir string) error {
	file, err := os.Open(tarballPath)
	if err != nil {
		return fmt.Errorf("failed to open tarball: %w", err)
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return fmt.Errorf("error reading tarball: %w", err)
		}

		targetPath := filepath.Join(destDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		case tar.TypeReg:
			outFile, err := os.Create(targetPath)
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}
			defer outFile.Close()

			if _, err := io.Copy(outFile, tarReader); err != nil {
				return fmt.Errorf("failed to extract file: %w", err)
			}
		}
	}

	return nil
}

func (f *FSManager) CleanUp() error {

	for _, layerPath := range f.layers {
		syscall.Unmount(layerPath, 0)
	}

	return syscall.Unmount(LIB_FS_MERGED_DIR, 0)
}
