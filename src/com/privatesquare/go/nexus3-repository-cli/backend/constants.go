package backend

const (

	// API Extensions
	apiBase        = "service/rest"
	scriptAPI      = "v1/script"
	repositoryPath = "v1/repositories"

	// Script Path
	scriptBasePath = "./scripts/groovy"

	// Error Strings
	jsonMarshalError   = "JSON Marshal Error"
	jsonUnmarshalError = "JSON Unmarshal Error"

	// Info Strings
	scriptNameRequiredInfo      = "script-name is a required parameter"
	repoNameRequiredInfo        = "repo-name is a required parameter"
	mavenHostedRepoRequiredInfo = "repo-name is a required parameter"
	proxyRepoRequiredInfo       = "repo-name and remote-url are required parameters"
	groupRequiredInfo           = "repo-name and repo-members are required parameters"
	setVerboseInfo              = "There was an error calling the function. Set verbose flag for more information"
)
