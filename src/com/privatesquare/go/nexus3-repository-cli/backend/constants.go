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
	scriptNameRequiredInfo = "script-name is a required parameter"
	repoNameRequiredInfo   = "repo-name is a required parameter"
	repoFormatRequiredInfo = "repo-format is a required parameter"
	proxyRepoRequiredInfo  = "repo-name and remote-url are required parameters"
	groupRequiredInfo      = "repo-name and repo-members are required parameters"
	dockerPortsInfo        = "You need to specify either a http port or a https port or both for creating a docker repository"
	setVerboseInfo         = "There was an error calling the function. Set verbose flag for more information"
)
