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
	scriptNameRequiredInfo = "scriptName is a required parameter"
	repoNameRequiredInfo   = "repoName is a required parameter"
	setVerboseInfo         = "There was an error calling the function. Set verbose flag for more information"
)
