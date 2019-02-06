package model

type ContentSelector struct {
	Name        string                    `json:"name"`
	Type        string                    `json:"type"`
	Description string                    `json:"description"`
	Attributes  ContentSelectorAttributes `json:"attributes"`
}

type ContentSelectorAttributes struct {
	Expression string `json:"expression"`
}

type Privilege struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Type        string              `json:"type"`
	Properties  PrivilegeProperties `json:"properties"`
	ReadOnly    bool                `json:"readOnly"`
}

type PrivilegeProperties struct {
	ContentSelector string `json:"contentSelector"`
	Repository      string `json:"repository"`
	Actions         string `json:"actions"`
}
