package parser

import (
	"encoding/json"
	"fmt"
)

type ImageManifest struct {
	Tar       string `json:"tar"`
	Signature string `json:"signature"`
}

func ParseImageManifest(data []byte) (*ImageManifest, error) {

	// Parse JSON
	var metadata ImageManifest
	if err := json.Unmarshal(data, &metadata); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &metadata, nil
}
