package models

type FileEntry struct {
	Path   string `json:"path"`
	Size   int64  `json:"size"`
	SHA256 string `json:"sha256"`
	URL    string `json:"url,omitempty"`
}
type Manifest struct {
	App     string      `json:"app"`
	Version string      `json:"version"`
	BaseURL string      `json:"base_url"`
	Files   []FileEntry `json:"files"`
}
