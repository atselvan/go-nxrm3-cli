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
	TaskFlag     = "task"
	SkipTlsFlag  = "skip-tls"
	SkipTlsUsage = "Skip TLS verification for the nexus server instance"
	DebugFlag    = "debug"
	DebugUsage   = "Set Default for more information on the nexus script execution"
	VerboseFlag  = "verbose"
	VerboseUsage = "Set Verbose for detailed http request and response logs"

	TaskEmptyInfo    = "You need to select a task to be performed. Available tasks : %+q\n"
	TaskNotValidInfo = "%q is not a valid %s task. Available tasks : %+q\n"

	DescFlag = "description"

	ListTask        = "list"
	CreateTask      = "create"
	UpdateTask      = "update"
	DeleteTask      = "delete"
	AddTask         = "add"
	AddOrUpdateTask = "add-or-update"
	RunTask         = "run"

	//script
	ScriptTaskUsage = "Script Task (Required)  (For all tasks the script(s) should exist under the path ./scripts/groovy)\n\n" +
		"    list 	    List all the scripts available in Nexus. (Optional: name) If script name is passed the contents of the script will be printed\n" +
		"    add  	    Add a new script to nexus. (Required: name)\n" +
		"    update 	    Update a script that is available in nexus. (Required: name)\n" +
		"    add-or-update   Add or Update a script in nexus. (Required: name)\n" +
		"    delete          Delete a script from nexus. (Required: name)\n" +
		"    run 	    Run/Execute a script in nexus. (Required: name)(Optional: payload)\n"

	ScriptNameFlag     = "name"
	ScriptNameUsage    = "Name of the script"
	ScriptPayloadFlag  = "payload"
	ScriptPayloadUsage = "Arguments can be passed to a nexus script as a payload during script execution"

	//repo
	RepoTaskUsage = "Repo Task (Required)\n\n" +
		"    list   		  List all the repositories in nexus.\n" +
		"			  (Optional: name) If repo-name is passed the list command will get the details of the repository.\n" +
		"			  (Optional: format) If repo-format is passed the list command will list the repositories as per the format\n" +
		"    create-hosted	  Create a hosted repository in nexus. (Required: name and format)\n" +
		"    create-proxy	  Create a proxy repository in nexus. (Required: name, repo-format and remote-url ) (Optional: proxy-user and proxy-pass)\n" +
		"    create-group	  Create a group repository in nexus. (Required: name,repo-format and repo-members)\n" +
		"    add-group-members	  Add new members to an existing group repository. Comma-separated values (Required: name, format and members)\n" +
		"    remove-group-members  Remove members from an existing group repository. Comma-separated values (Required: name, format and members)\n" +
		"    delete		  Delete a repository from nexus (Required: name)\n\n" +
		"    If you are creating a docker repository it is necessary to also provide either a docker-http-port or a docker-https-port or both.\n"

	RepoNameFlag         = "name"
	RepoNameUsage        = "Nexus repository name"
	RepoFormatFlag       = "format"
	RepoFormatUsage      = "Repository format. Available formats : %+q"
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
	SelectorTaskUsage = "Selector Task (Required)\n\n" +
		"    list 	    List all the content selectors in nexus (Optional: name)\n" +
		"    create  	    Create a content selector in nexus (Required: name and expression) (Optional: description)\n" +
		"    update 	    Update the details of a content selector. (Required: name and expression) (Optional: description)\n" +
		"    delete          Delete a content selector (Required: name)\n"

	SelectorNameFlag        = "name"
	SelectorNameUsage       = "Content Selector name"
	SelectorDescUsage       = "Content Selector description"
	SelectorExpressionFlag  = "expression"
	SelectorExpressionUsage = "Pattern expression for the content selector"

	//privilege
	PrivilegeTaskUsage = "Privilege Task (Required)\n\n" +
		"    list 	    List all the privileges in nexus (Optional: name)\n" +
		"    create  	    Create a Privilege in nexus (Required: name, selector-name and repo-name) (Optional: description and action)\n" +
		"    update 	    Update the details of a Privilege. (Required: name)(Optional: selector-name, repo-name, description and action)\n" +
		"    delete          Delete a Privilege (Required: name)\n"

	PrivilegeNameFlag  = "name"
	PrivilegeNameUsage = "Privilege name"
	PSelectorNameFlag  = "selector-name"
	PRepoNameFlag      = "repo-name"
	ActionFlag         = "action"
	ActionUsage        = "Privilege Action. Available actions %+q"
	PrivilegeDescUsage = "Privilege description"

	//role
	RoleTaskUsage = "Role Task (Required)\n\n" +
		"    list 	    List all the roles in nexus (Optional: id)\n" +
		"    create  	    Create a role in nexus (Required: id) (Optional: description, role-members, role-privileges)\n" +
		"    update 	    Update the details of a role. (Required: id, action) (Optional: description, role-members, role-privileges)\n" +
		"    delete          Delete a Privilege (Required: id)\n"

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
