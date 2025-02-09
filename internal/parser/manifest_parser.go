package parser

type ImageManifest struct {
	Tar       string `json:"tar"`
	Signature string `json:"signature"`
}
