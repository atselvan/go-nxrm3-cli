package main

import (
	b "com/privatesquare/go/nexus3-repository-cli/backend"
	m "com/privatesquare/go/nexus3-repository-cli/model"
	"flag"
	"log"
	"os"
)

func main() {

	//TODO : Explain the script and repo tasks

	// subcommands
	confCommand := flag.NewFlagSet("configure", flag.ExitOnError)
	scriptCommand := flag.NewFlagSet("script", flag.ExitOnError)
	repoCommand := flag.NewFlagSet("repo", flag.ExitOnError)
	// conf flags
	nexusURL := confCommand.String("nexus-url", "", "Nexus 3 server URL. (Required)")
	username := confCommand.String("username", "", "Nexus 3 server login user. (Required)")
	password := confCommand.String("password", "", "Nexus 3 server login password. (Required)")
	// script flags
	scriptTask := scriptCommand.String("task", "", "Script Task (list|add|update|add-or-update|delete|run). (Required)")
	scriptName := scriptCommand.String("script-name", "", "Name of the script to be executed in nexus. \nThe script should exist under the path ./scripts/groovy")
	scriptPayload := scriptCommand.String("payload", "", "Arguments to be passed to a nexus script can be sent as a payload during script execution.")
	scSkipTLS := scriptCommand.Bool("skip-tls", false, "Skip TLS verification for the nexus server instance.")
	scDebug := scriptCommand.Bool("debug", false, "Set Default for more information on the nexus script execution.")
	scVerbose := scriptCommand.Bool("verbose", false, "Set Verbose for detailed http request and response logs.")
	// repo flags
	repoTask := repoCommand.String("task", "", "Script Task (list|create-maven-hosted|create-maven-proxy|create-maven-group|delete). (Required)")
	repoName := repoCommand.String("repo-name", "", "Nexus repository name")
	repoFormat := repoCommand.String("repo-format", "", "Repository format (maven|npm|nuget|docker).")
	remoteURL := repoCommand.String("remote-url", "", "Remote URL to be proxied in nexus.")
	proxyUser := repoCommand.String("proxy-user", "", "Username for accessing the proxy repository")
	proxyPass := repoCommand.String("proxy-pass", "", "Password for accessing the proxy repository")
	dockerHttpPort := repoCommand.Int("docker-http-port", 0, "Docker HTTP port.")
	dockerHttpsPort := repoCommand.Int("docker-https-port", 0, "Docker HTTPs port")
	blobStoreName := repoCommand.String("blob-store-name", "", "Blob store name.")
	releases := repoCommand.Bool("releases", false, "Set this flag to create a releases repository.")
	//repoMembers := repoCommand.String("repo-members", "", "Comma-separated repository names that should be added to a group repo.")
	rcSkipTLS := repoCommand.Bool("skip-tls", false, "Skip TLS verification for the nexus server instance.")
	rcDebug := repoCommand.Bool("debug", false, "Set Default for more information on the nexus script execution.")
	rcVerbose := repoCommand.Bool("verbose", false, "Set Verbose for detailed http request and response logs.")

	b.Usage()

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "configure":
		b.ConfigureCommandUsage(confCommand)
		confCommand.Parse(os.Args[2:])
	case "script":
		b.ScriptCommandUsage(scriptCommand)
		scriptCommand.Parse(os.Args[2:])
	case "repo":
		b.RepoCommandUsage(repoCommand)
		repoCommand.Parse(os.Args[2:])
	default:
		flag.Usage()
		os.Exit(1)
	}

	if confCommand.Parsed() {
		// Required Flags
		if *nexusURL == "" || *username == "" || *password == "" {
			confCommand.Usage()
			os.Exit(1)
		}
		b.NexusURL = *nexusURL
		b.AuthUser = m.AuthUser{Username: *username, Password: *password}
		b.SetCLIConfiguration()
	}

	if scriptCommand.Parsed() {
		// Required Flags
		if *scriptTask == "" {
			scriptCommand.Usage()
			log.Println("You need to select a task to be performed.")
			os.Exit(1)
		}
		// set global variables
		b.SetConnectionDetails()
		b.SkipTLSVerification = *scSkipTLS
		b.Debug = *scDebug
		b.Verbose = *scVerbose
		// run tasks
		switch *scriptTask {
		case "list":
			b.ListScripts(*scriptName)
		case "add":
			b.Debug = true
			b.AddScript(*scriptName)
		case "update":
			b.Debug = true
			b.UpdateScript(*scriptName)
		case "add-or-update":
			b.Debug = true
			b.AddOrUpdateScript(*scriptName)
		case "delete":
			b.Debug = true
			b.DeleteScript(*scriptName)
		case "run":
			b.Debug = true
			b.RunScript(*scriptName, *scriptPayload)
		default:
			scriptCommand.Usage()
			log.Printf("%q is not a valid task.\n\n", *scriptTask)
			os.Exit(1)
		}
	}

	if repoCommand.Parsed() {
		// Required Flags
		if *repoTask == "" {
			repoCommand.Usage()
			log.Println("You need to select a task to be performed.")
			os.Exit(1)
		}
		//Choice flag
		repoFormatChoice := map[string]bool{"": true}
		for _, repoFormat := range b.RepoFormats {
			repoFormatChoice[repoFormat] = true
		}
		if _, validChoice := repoFormatChoice[*repoFormat]; !validChoice {
			log.Printf("%q is not a valid repository format. Available repository formats are : %v\n", *repoFormat, b.RepoFormats)
			os.Exit(1)
		}
		// set global variables
		b.SetConnectionDetails()
		b.SkipTLSVerification = *rcSkipTLS
		b.Debug = *rcDebug
		b.Verbose = *rcVerbose
		// run tasks
		switch *repoTask {
		case "list":
			b.ListRepositories(*repoName, *repoFormat)
		case "get-attributes":
			b.GetRepositoryAttributes(*repoName)
		case "create-hosted":
			b.CreateHosted(*repoName, *blobStoreName, *repoFormat, *dockerHttpPort, *dockerHttpsPort, *releases)
		case "create-proxy":
			b.CreateProxy(*repoName, *blobStoreName, *repoFormat, *remoteURL, *proxyUser, *proxyPass, *dockerHttpPort, *dockerHttpsPort, *releases)
		case "delete":
			b.DeleteRepository(*repoName)
		default:
			repoCommand.Usage()
			log.Printf("%q is not a valid task.\n\n", *repoTask)
			os.Exit(1)
		}
	}
}
