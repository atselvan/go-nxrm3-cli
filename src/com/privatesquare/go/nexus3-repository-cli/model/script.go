package model

type Script struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type ScriptOutput struct {
	Name   string `json:"name"`
	Result string `json:"result"`
}

type ScriptResult struct {
	Status  string `json:"status"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Type    string `json:"type"`
	Format  string `json:"format"`
	Recipe  string `json:"recipe"`
	Message string `json:"message"`
}
