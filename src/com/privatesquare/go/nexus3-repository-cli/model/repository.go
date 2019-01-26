package model

type Repository struct {
	Name          string   `json:"name"`
	URL           string   `json:"url"`
	Type          string   `json:"type"`
	Format        string   `json:"format"`
	Recipe        string   `json:"recipe"`
	BlobStoreName string   `json:"blobStoreName"`
	VersionPolicy string   `json:"versionPolicy"`
	RemoteURL     string   `json:"remoteURL"`
	Members       []string `json:"members"`
}
