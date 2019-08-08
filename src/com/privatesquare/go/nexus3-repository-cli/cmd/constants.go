package cmd

const (

	// CLI flags and usage
	ConfCommandFlag       = "configure"
	ConfCommandUsage      = "Set nexus connection details"
	ScriptCommandFlag     = "scripts"
	ScriptCommandUsage    = "Nexus script operations"
	RepoCommandFlag       = "repo"
	RepoCommandUsage      = "Nexus repository operations"
	SelectorCommandFlag   = "selector"
	SelectorCommandUsage  = "Nexus content selector operations"
	PrivilegeCommandFlag  = "privilege"
	PrivilegeCommandUsage = "Nexus privilege operations"
	RoleCommandFlag       = "role"
	RoleCommandUsage      = "Nexus role operations"

	//configure
	NexusURLFlag       = "nexus-url"
	NexusURLUsage      = "Nexus 3 server URL. (Required)"
	NexusUsernameFlag  = "username"
	NexusUsernameUsage = "Nexus 3 server login user. (Required)"
	NexusPasswordFlag  = "password"
	NexusPasswordUsage = "Nexus 3 server login password. (Required)"

	//common
	SkipTlsFlag  = "skip-tls"
	SkipTlsUsage = "Skip TLS verification for the nexus server instance"
	DebugFlag    = "debug"
	DebugUsage   = "Set Default for more information on the nexus script execution"
	VerboseFlag  = "verbose"
	VerboseUsage = "Set Verbose for detailed http request and response logs"

	DescFlag = "description"

	ListTask        = "list"
	CreateTask      = "create"
	UpdateTask      = "update"
	DeleteTask      = "delete"
	AddTask         = "add"
	AddOrUpdateTask = "add-or-update"
	RunTask         = "run"

	//script
	ScriptNameFlag     = "name"
	ScriptNameUsage    = "Name of the script"
	ScriptPayloadFlag  = "payload"
	ScriptPayloadUsage = "Arguments can be passed to a nexus script as a payload during script execution"

	//repo
	RepoNameFlag         = "name"
	RepoNameUsage        = "Nexus repository name"
	RepoFormatFlag       = "format"
	RepoFormatUsage      = "Repository format. Available formats : %+q"
	RepoTypeFlag = "type"
	RepoTypeUsage = "Repository Type. Available types : %+q"
	RemoteURLFlag        = "remote-url"
	RemoteURLUsage       = "Remote URL to be proxied in nexus"
	RepoMembersFlag      = "members"
	RepoMembersUsage     = "Comma-separated repository names that should be added to a group repo"
	ProxyUserFlag        = "proxy-user"
	ProxyUserUsage       = "Username for accessing the proxy repository"
	ProxyPassFlag        = "proxy-pass"
	ProxyPassUsage       = "Password for accessing the proxy repository"
	DockerHttpPortFlag   = "docker-http-port"
	DockerHttpPortUsage  = "Docker HTTP port"
	DockerHttpsPortFlag  = "docker-https-port"
	DockerHttpsPortUsage = "Docker HTTPs port"
	BlobStoreNameFlag    = "blob-store-name"
	BlobStoreNameUsage   = "Blob store name"
	ReleaseFlag          = "releases"
	ReleaseUsage         = "Set this flag to create a releases repository"

	//selector
	SelectorNameFlag        = "name"
	SelectorNameUsage       = "Content Selector name"
	SelectorDescUsage       = "Content Selector description"
	SelectorExpressionFlag  = "expression"
	SelectorExpressionUsage = "Pattern expression for the content selector"

	//privilege
	PrivilegeNameFlag  = "name"
	PrivilegeNameUsage = "Privilege name"
	PSelectorNameFlag  = "selector-name"
	PRepoNameFlag      = "repo-name"
	ActionFlag         = "action"
	ActionUsage        = "Privilege Action. Available actions %+q"
	PrivilegeDescUsage = "Privilege description"

	//role
	RoleIDFlag          = "id"
	RoleIDUsage         = "Role ID"
	RoleDescUsage       = "Role description"
	RoleMembersFlag     = "role-members"
	RoleMembersUsage    = "Comma separated role member id's to be added to a role"
	RolePrivilegesFlag  = "role-privileges"
	RolePrivilegesUsage = "Comma separated privileges to be added to a role"

	UpdateActionFlag         = "action"
	UpdateActionUsage        = "Update Action. Available values = %+q\n"
)
