package backend

import m "com/privatesquare/go/nexus3-repository-cli/model"

var (
	NexusURL            string
	AuthUser            m.AuthUser
	Verbose             bool
	Debug               bool
	SkipTLSVerification bool

	InitialRepoList   = []string{"maven-public", "maven-central", "maven-snapshots", "maven-releases", "nuget-group", "nuget-hosted", "nuget.org-proxy"}
	NexusScripts      = []string{"get-repo", "create-hosted-repo", "create-proxy-repo", "create-group-repo", "update-group-members", "delete-repo", "get-content-selectors", "create-content-selector", "update-content-selector", "delete-content-selector", "get-privileges", "create-privilege", "update-privilege", "delete-privilege", "get-roles", "create-role", "delete-role"}
	RepoFormats       = []string{"maven", "npm", "nuget", "bower", "pypi", "raw", "rubygems", "yum", "docker"}
	ScriptTasks       = []string{ListTask, AddTask, UpdateTask, AddOrUpdateTask, DeleteTask, RunTask}
	RepoTasks         = []string{ListTask, "create-hosted", "create-proxy", "create-group", "add-group-members", "remove-group-members", DeleteTask}
	SelectorTasks     = []string{ListTask, CreateTask, UpdateTask, DeleteTask}
	PrivilegeTasks    = []string{ListTask, CreateTask, UpdateTask, DeleteTask}
	PrivilegeActions  = []string{"all", "read", "write"}
	RoleTasks         = []string{ListTask, CreateTask, UpdateTask, DeleteTask}
	UpdateActions     = []string{"add", "remove"}
)
