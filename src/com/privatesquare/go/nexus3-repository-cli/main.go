package main

import (
	"com/privatesquare/go/nexus3-repository-cli/backend"
	"com/privatesquare/go/nexus3-repository-cli/model"
	"fmt"
)

func main() {

	backend.NexusURL = "https://localhost:8443"
	backend.AuthUser = model.AuthUser{Username: "admin", Password: "admin123"}
	backend.SkipTLSVerification = true
	backend.Verbose = true

	backend.ListScripts()
	fmt.Println(backend.GetScript("create-maven-hosted"))
 	backend.AddOrUpdateScript("create-maven-hosted")
	backend.RunScript("create-maven-hosted")
}





