package backend

import m "com/privatesquare/go/nexus3-repository-cli/model"

var (
	NexusURL            string
	AuthUser            m.AuthUser
	Verbose             bool
	Debug               bool
	SkipTLSVerification bool
	RepoFormats         = []string{"maven", "npm", "nuget", "bower", "pypi", "raw", "rubygems", "yum", "docker"}
	ScriptTasks         = []string{"list", "add", "update", "add-or-update", "delete", "run"}
	RepoTasks           = []string{"list", "create-hosted", "create-proxy", "create-group", "add-group-member", "delete"}
	SelectorTasks       = []string{"list", "create", "update", "delete"}
	PrivilegeTasks      = []string{"list", "create", "update", "delete"}
	PrivilegeActions    = []string{"all", "read", "write"}
)
