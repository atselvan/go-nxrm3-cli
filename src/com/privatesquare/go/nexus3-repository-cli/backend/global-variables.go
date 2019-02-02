package backend

import m "com/privatesquare/go/nexus3-repository-cli/model"

var (
	NexusURL            string
	AuthUser            m.AuthUser
	Verbose             bool
	Debug               bool
	SkipTLSVerification bool
	RepoFormats         = []string{"maven", "npm", "nuget", "bower", "pypi", "raw", "rubygems", "yum", "docker"}
)
