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
