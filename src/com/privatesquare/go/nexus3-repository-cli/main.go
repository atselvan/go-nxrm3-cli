package main

import (
	b "com/privatesquare/go/nexus3-repository-cli/backend"
	m "com/privatesquare/go/nexus3-repository-cli/model"
)

func main() {

	// test parameters
	b.NexusURL = "https://localhost:8443"
	b.AuthUser = m.AuthUser{Username: "admin", Password: "admin123"}
	b.SkipTLSVerification = true
	b.Verbose = false
	b.Debug = false

	b.ListScripts("")
	b.ListRepositories("", "")

	b.CreateMavenHostedRepository("maven-releases-test", "", true)
	b.CreateMavenHostedRepository("maven-snapshots-test", "", false)
	b.CreateMavenProxyRepository("maven-proxy-test", "", "https://repo1.maven.org/maven2/")
	b.CreateMavenGroupRepository("maven-group-test", "", []string{"maven-releases-test", "maven-snapshots-test", "maven-proxy-test"})

	deleteRepo := false

	if deleteRepo {
		b.DeleteRepository("maven-releases-test")
		b.DeleteRepository("maven-snapshots-test")
		b.DeleteRepository("maven-proxy-test")
		b.DeleteRepository("maven-group-test")
	}
}
