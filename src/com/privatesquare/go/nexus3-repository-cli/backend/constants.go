package backend

const (
	ConfFileName = "nexus3-repository-cli.json"

	// API Extensions
	apiBase        = "service/rest"
	scriptAPI      = "v1/script"
	repositoryPath = "v1/repositories"

	successStatus   = "200 OK"
	notFoundStatus  = "404 Not Found"
	noContentStatus = "204 No Content"
	foundStatus     = "302 Found"

	// Script Path
	scriptBasePath = "./scripts/groovy"

	// Error Strings
	jsonMarshalError   = "JSON Marshal Error"
	jsonUnmarshalError = "JSON Unmarshal Error"

	// CLI flags and usage
	ConfCommandFlag       = "configure"
	ConfCommandUsage      = "Set nexus connection details"
	ScriptCommandFlag     = "script"
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

	connDetailsSuccessInfo = "Connection details were stored successfully in the file ./%s\n"
	connDetailsEmptyInfo   = "Server connection details are not set...First Run %q to set the connection details\n"

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
	setVerboseInfo   = "There was an error calling the function. Set verbose flag for more information"

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

	scriptNameRequiredInfo = "name is a required parameter"
	scriptAddedInfo        = "The script %q is added to nexus\n"
	scriptUpdatedInfo      = "The script %q is updated in nexus\n"
	scriptDeletedInfo      = "The script %q is deleted from nexus\n"
	scriptRunSuccessInfo   = "The script %q was executed successfully\n"
	scriptRunNotFoundInfo  = "The script %q was not found in nexus. Make sure you add the script to nexus before executing the script\n"
	scriptExistsInfo       = "The script %q already exists in nexus\n"
	scriptNotfoundInfo     = "The script %q was not found in nexus\n"

	//scripts
	getRepoScript            = "get-repo"
	createHostedRepoScript   = "create-hosted-repo"
	createProxyRepoScript    = "create-proxy-repo"
	createGroupRepoScript    = "create-group-repo"
	updateGroupMembersScript = "update-group-members"
	deleteRepoScript         = "delete-repo"
	getPrivilegesScript      = "get-privileges"
	createPrivilegeScript    = "create-privilege"
	updatePrivilegeScript    = "update-privilege"
	deletePrivilegeScript    = "delete-privilege"
	getRoleScript            = "get-roles"
	createRoleScript         = "create-role"
	deleteRoleScript         = "delete-role"

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

	RepoFormatNotValidInfo        = "%q is not a valid repository format. Available repository formats are : %v\n"
	repoNameRequiredInfo          = "name is a required parameter"
	repoFormatRequiredInfo        = "format is a required parameter"
	hostedRepoRequiredInfo        = "name and format are required parameters to create a hosted repository"
	proxyRepoRequiredInfo         = "name, format and remote-url are required parameters to create a proxy repository"
	groupRequiredInfo             = "name, format and members are required parameters"
	dockerPortsInfo               = "You need to specify either a http port or a https port or both for creating a docker repository"
	repositoryNotFoundInfo        = "Repository %q was not found in nexus"
	repoCreatedInfo               = "Repository %q was created in nexus\n"
	repoUpdatedStatus             = "Repository %q was updated in nexus\n"
	repoDeletedInfo               = "Repository %q was deleted from nexus\n"
	repoCreateErrorInfo           = "Error creating repository : %s\n"
	repoUpdateErrorInfo           = "Error updating repository : %s\n"
	repoDeleteErrorInfo           = "Error deleting repository : %s\n"
	repoExistsInfo                = "Repository %q already exists in nexus\n"
	cannotBeSameRepoInfo          = "Member %q == group %q, cannot add a group repository as a member in the same group\n"
	proxyCredsNotValidInfo        = "You need to provide both proxy-user and proxy-pass to set credentials to a proxy repository"
	remoteURLNotValidInfo         = "%q is an invalid url. URL must begin with either http:// or https://"
	notAGroupRepoInfo             = "%q is not a group repository\n"
	groupMemberInvalidFormatInfo  = "Repository %q is not a %q format repository, hence it cannot be added to the group repository\n"
	groupMemberAlreadyExistsInfo  = "Member %q already exists in the group %q, hence not adding the member again\n"
	groupMemberNotFoundInfo       = "Repository %q was not found in Nexus, hence it cannot be added to the group repository\n"
	groupMemberRemoveNotFoundInfo = "Member %q was not found in the group %q, hence cannot remove the member from the group\n"
	groupMemberRequiredInfo       = "At least one valid group member should be provided to add to a group repository"
	groupMemberAddSuccessInfo     = "Member %q is added to the group %q\n"
	groupMemberRemoveSuccessInfo  = "Member %q is removed from the group %q\n"

	//selector
	contentSelectorType = "csel"

	SelectorTaskUsage = "Selector Task (Required)\n\n" +
		"    list 	    List all the content selectors in nexus (Optional: name)\n" +
		"    create  	    Create a content selector in nexus (Required: name and expression) (Optional: description)\n" +
		"    update 	    Update the details of a content selector. (Required: name and expression) (Optional: description)\n" +
		"    delete          Delete a content selector (Required: name)\n"

	SelectorNameFlag        = "name"
	SelectorNameUsage       = "Content Selector name"
	SelectorDescFlag        = "description"
	SelectorDescUsage       = "Content Selector description"
	SelectorExpressionFlag  = "expression"
	SelectorExpressionUsage = "Pattern expression for the content selector"

	defaultContentSelectorDescription = "Custom content-selector created from the CLI"
	selectorNameRequiredInfo          = "name is a required parameter"
	createSelectorRequiredInfo        = "name and expression are required parameters"
	createSelectorSuccessInfo         = "Content selector %q was created\n"
	updateSelectorSuccessInfo         = "Content selector %q was updated\n"
	deleteSelectorSuccessInfo         = "Content selector %q was deleted\n"
	selectorAlreadyExistsInfo         = "Content selector %q already exists in nexus\n"
	selectorNotFoundInfo              = "Content selector %q was not found in nexus\n"

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
	PrivilegeDescFlag  = "description"
	PrivilegeDescUsage = "Privilege description"

	defaultPrivilegeDescription = "Custom privilege created from the CLI"
	privilegeNameRequiredInfo   = "name is a required parameter"
	privilegeNotFoundInfo       = "Privilege %q was not found in nexus\n"
	privilegeExistsInfo         = "Privilege %q already exists\n"
	createPrivilegeRequiredInfo = "name, selector-name and repo-name are required parameters"
	createPrivilegeSuccessInfo  = "Privilege %q is created"
	updatePrivilegeSuccessInfo  = "Privilege %q is updated"
	deletePrivilegeSuccessInfo  = "Privilege %q is deleted"

	//role
	RoleTaskUsage = "Role Task (Required)\n\n" +
		"    list 	    List all the roles in nexus (Optional: id)\n" +
		"    create  	    Create a role in nexus (Required: id) (Optional: description, role-members, role-privileges)\n" +
		"    update 	    Update the details of a role. (Required: id, action) (Optional: description, role-members, role-privileges)\n" +
		"    delete          Delete a Privilege (Required: id)\n"

	RoleIDFlag          = "id"
	RoleIDUsage         = "Role ID"
	RoleDescFlag        = "description"
	RoleDescUsage       = "Role description"
	RoleMembersFlag     = "role-members"
	RoleMembersUsage    = "Comma separated role member id's to be added to a role"
	RolePrivilegesFlag  = "role-privileges"
	RolePrivilegesUsage = "Comma separated privileges to be added to a role"

	UpdateActionFlag         = "action"
	UpdateActionUsage        = "Update Action. Available values = %+q\n"
	UpdateActionRequiredInfo = "Update action is a required parameter. Available values = %+q\n"
	UpdateActionInvalidInfo  = "%s is not a valid update action. Available actions: %+q\n"

	defaultRoleDescription        = "Custom role created from the CLI"
	defaultRoleSource             = "Nexus"
	roleIDRequiredInfo            = "id is a required parameter"
	roleNotFoundInfo              = "Role %q was not found in nexus\n"
	roleExistsInfo                = "Role %q already exists\n"
	createRoleRequiredInfo        = "id, description and source are required parameters"
	createRoleSuccessInfo         = "Role %q is created with role members %v and privileges %+q\n"
	updateRoleSuccessInfo         = "Role %q is updated\n"
	deleteRoleSuccessInfo         = "Role %q is deleted\n"
	roleItemsRequiredInfo         = "%s : You need to provide at least one valid role member or role privilege during role creation\n"
	noRoleMemberProvidedInfo      = "No role members are provided to add to the role"
	noValidRoleMemberInfo         = "No valid role members are provided to add to the role"
	cannotBeSameRoleInfo          = "Role member %q == role id %s, cannot add a role as a member in the same role"
	roleMemberNotFoundInfo        = "Role %q was not found in nexus, hence it cannot be added to the role"
	rolePrivilegeNotFoundInfo     = "Privilege %q was not found in nexus, hence it cannot be added to the role"
	noValidRolePrivilegeInfo      = "No valid privileges are provided to add to the role"
	noRolePrivilegesIProvidedInfo = "No privileges are provided to add to the role"
)
