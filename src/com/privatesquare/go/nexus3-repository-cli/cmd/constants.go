package cmd

const (

	// sub-commands
	confCommandFlag       = "configure"
	confCommandUsage      = "Set nexus connection details"
	scriptCommandFlag     = "scripts"
	scriptCommandUsage    = "Nexus script operations"
	repoCommandFlag       = "repo"
	repoCommandUsage      = "Nexus repository operations"
	selectorCommandFlag   = "selector"
	selectorCommandUsage  = "Nexus content selector operations"
	privilegeCommandFlag  = "privilege"
	privilegeCommandUsage = "Nexus privilege operations"
	roleCommandFlag       = "role"
	roleCommandUsage      = "Nexus role operations"

	//configure
	nexusURLFlag       = "nexus-url"
	nexusURLUsage      = "Nexus 3 server URL. (Required)"
	nexusUsernameFlag  = "username"
	nexusUsernameUsage = "Nexus 3 server login user. (Required)"
	nexusPasswordFlag  = "password"
	nexusPasswordUsage = "Nexus 3 server login password. (Required)"

	//common
	skipTlsFlag  = "skip-tls"
	skipTlsUsage = "Skip TLS verification for the nexus server instance"
	debugFlag    = "debug"
	debugUsage   = "Set Default for more information on the nexus script execution"
	verboseFlag  = "verbose"
	verboseUsage = "Set Verbose for detailed http request and response logs"

	descFlag = "description"

	initTask        = "init"
	listTask        = "list"
	createTask      = "create"
	updateTask      = "update"
	deleteTask      = "delete"
	addTask         = "add"
	addOrUpdateTask = "add-or-update"
	runTask         = "run"

	//script
	scriptNameFlag     = "name"
	scriptNameUsage    = "Name of the script"
	scriptPayloadFlag  = "payload"
	scriptPayloadUsage = "Arguments can be passed to a nexus script as a payload during script execution"

	//repo
	repoNameFlag         = "name"
	repoNameUsage        = "Nexus repository name"
	repoFormatFlag       = "format"
	repoFormatUsage      = "Repository format. Available formats : %+q"
	repoTypeFlag         = "type"
	repoTypeUsage        = "Repository Type. Available types : %+q"
	remoteURLFlag        = "remote-url"
	remoteURLUsage       = "Remote URL to be proxied in nexus"
	repoMembersFlag      = "members"
	repoMembersUsage     = "Comma-separated repository names that should be added to a group repo"
	proxyUserFlag        = "proxy-user"
	proxyUserUsage       = "Username for accessing the proxy repository"
	proxyPassFlag        = "proxy-pass"
	proxyPassUsage       = "Password for accessing the proxy repository"
	dockerHttpPortFlag   = "docker-http-port"
	dockerHttpPortUsage  = "Docker HTTP port"
	dockerHttpsPortFlag  = "docker-https-port"
	dockerHttpsPortUsage = "Docker HTTPs port"
	blobStoreNameFlag    = "blob-store-name"
	blobStoreNameUsage   = "Blob store name"
	releaseFlag          = "releases"
	releaseUsage         = "Set this flag to create a releases repository"

	//selector
	selectorNameFlag        = "name"
	selectorNameUsage       = "Content Selector name"
	selectorDescUsage       = "Content Selector description"
	selectorExpressionFlag  = "expression"
	selectorExpressionUsage = "Pattern expression for the content selector"

	//privilege
	privilegeNameFlag  = "name"
	privilegeNameUsage = "Privilege name"
	pSelectorNameFlag  = "selector-name"
	pRepoNameFlag      = "repo-name"
	actionFlag         = "action"
	actionUsage        = "Privilege Action. Available actions %+q"
	privilegeDescUsage = "Privilege description"

	//role
	roleIDFlag          = "id"
	roleIDUsage         = "Role ID"
	roleDescUsage       = "Role description"
	roleMembersFlag     = "role-members"
	roleMembersUsage    = "Comma separated role member id's to be added to a role"
	rolePrivilegesFlag  = "role-privileges"
	rolePrivilegesUsage = "Comma separated privileges to be added to a role"

	updateActionFlag  = "action"
	updateActionUsage = "Update Action. Available values = %+q\n"
)
