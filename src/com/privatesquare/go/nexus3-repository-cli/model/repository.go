package model

type Repository struct {
	Name      string     `json:"name"`
	URL       string     `json:"url"`
	Type      string     `json:"type"`
	Format    string     `json:"format"`
	Recipe    string     `json:"recipe"`
	Attribute Attributes `json:"attributes"`
}

type Attributes struct {
	Storage       Storage       `json:"storage"`
	Maven         Maven         `json:"maven"`
	Proxy         Proxy         `json:"proxy"`
	Httpclient    HttpClient    `json:"httpclient"`
	Group         Group         `json:"group"`
	NegativeCache NegetiveCache `json:"negativeCache"`
	Docker        Docker        `json:"docker"`
	DockerProxy   DockerProxy   `json:"dockerProxy"`
	Cleanup       Cleanup       `json:"cleanup"`
}

type Storage struct {
	BlobStoreName               string `json:"blobStoreName"`
	WritePolicy                 string `json:"writePolicy"`
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation"`
}

type Maven struct {
	VersionPolicy string `json:"versionPolicy"`
	LayoutPolicy  string `json:"layoutPolicy"`
}

type Proxy struct {
	RemoteURL      string `json:"remoteUrl"`
	ContentMaxAge  int    `json:"contentMaxAge"`
	MetadataMaxAge int    `json:"metadataMaxAge"`
}

type HttpClient struct {
	Blocked        bool           `json:"blocked"`
	AutoBlock      bool           `json:"autoBlock"`
	Authentication HttpClientAuth `json:"authentication"`
}

type HttpClientAuth struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Docker struct {
	HTTPPort       int  `json:"httpPort"`
	HTTPSPort      int  `json:"httpsPort"`
	ForceBasicAuth bool `json:"forceBasicAuth"`
	V1Enabled      bool `json:"v1Enabled"`
}

type DockerProxy struct {
	IndexType string `json:"indexType"`
}

type Group struct {
	MemberNames []string `json:"memberNames"`
}

type NegetiveCache struct {
	Enabled    bool    `json:"enabled"`
	TimeToLive float64 `json:"timeToLive"`
}

type Cleanup struct {
	PolicyName string `json:"policyName"`
}
