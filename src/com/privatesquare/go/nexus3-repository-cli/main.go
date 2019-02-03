package main

import (
	b "com/privatesquare/go/nexus3-repository-cli/backend"
	m "com/privatesquare/go/nexus3-repository-cli/model"
	"flag"
	"log"
	"os"
)

func main() {

	// subcommands
	confCommand := flag.NewFlagSet(b.ConfCommandFlag, flag.ExitOnError)
	scriptCommand := flag.NewFlagSet(b.ScriptCommandFlag, flag.ExitOnError)
	repoCommand := flag.NewFlagSet(b.RepoCommandFlag, flag.ExitOnError)
	// conf flags
	nexusURL := confCommand.String(b.NexusURLFlag, "", b.NexusURLUsage)
	username := confCommand.String(b.NexusUsernameFlag, "", b.NexusUsernameUsage)
	password := confCommand.String(b.NexusPasswordFlag, "", b.NexusPasswordUsage)
	// script flags
	scriptTask := scriptCommand.String(b.TaskFlag, "", b.ScriptTaskUsage)
	scriptName := scriptCommand.String(b.ScriptNameFlag, "", b.ScriptNameUsage)
	scriptPayload := scriptCommand.String(b.ScriptPayloadFlag, "", b.ScriptPayloadUsage)
	scSkipTLS := scriptCommand.Bool(b.SkipTlsFlag, false, b.SkipTlsUsage)
	scDebug := scriptCommand.Bool(b.DebugFlag, false, b.DebugUsage)
	scVerbose := scriptCommand.Bool(b.VerboseFlag, false, b.VerboseUsage)
	// repo flags
	repoTask := repoCommand.String(b.TaskFlag, "", b.RepoTaskUsage)
	repoName := repoCommand.String(b.RepoNameFlag, "", b.RepoNameUsage)
	repoFormat := repoCommand.String(b.RepoFormatFlag, "", b.RepoFormatUsage)
	remoteURL := repoCommand.String(b.RemoteURLFlag, "", b.RemoteURLUsage)
	repoMembers := repoCommand.String(b.RepoMembersFlag, "", b.RepoMembersUsage)
	proxyUser := repoCommand.String(b.ProxyUserFlag, "", b.ProxyUserUsage)
	proxyPass := repoCommand.String(b.ProxyPassFlag, "", b.ProxyPassUsage)
	dockerHttpPort := repoCommand.Int(b.DockerHttpPortFlag, 0, b.DockerHttpPortUsage)
	dockerHttpsPort := repoCommand.Int(b.DockerHttpsPortFlag, 0, b.DockerHttpsPortUsage)
	blobStoreName := repoCommand.String(b.BlobStoreNameFlag, "", b.BlobStoreNameUsage)
	releases := repoCommand.Bool(b.ReleaseFlag, false, b.ReleaseUsage)
	rcSkipTLS := repoCommand.Bool(b.SkipTlsFlag, false, b.SkipTlsUsage)
	rcDebug := repoCommand.Bool(b.DebugFlag, false, b.DebugUsage)
	rcVerbose := repoCommand.Bool(b.VerboseFlag, false, b.VerboseUsage)

	b.Usage()

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "configure":
		b.PrintConfCommandUsage(confCommand)
		confCommand.Parse(os.Args[2:])
	case "script":
		b.PrintScriptCommandUsage(scriptCommand)
		scriptCommand.Parse(os.Args[2:])
	case "repo":
		b.PrintRepoCommandUsage(repoCommand)
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
		b.StoreConnectionDetails()
	}

	if scriptCommand.Parsed() {
		// Required Flags
		if *scriptTask == "" {
			scriptCommand.Usage()
			log.Printf(b.TaskEmptyInfo, b.ScriptTasks)
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
			log.Printf(b.TaskNotValidInfo, *scriptTask, "script", b.ScriptTasks)
			os.Exit(1)
		}
	}

	if repoCommand.Parsed() {
		// Required Flags
		if *repoTask == "" {
			repoCommand.Usage()
			log.Printf(b.TaskEmptyInfo, b.RepoTasks)
			os.Exit(1)
		}
		//Choice flag
		repoFormatChoice := map[string]bool{"": true}
		for _, repoFormat := range b.RepoFormats {
			repoFormatChoice[repoFormat] = true
		}
		if _, validChoice := repoFormatChoice[*repoFormat]; !validChoice {
			log.Printf(b.RepoFormatNotValidInfo, *repoFormat, b.RepoFormats)
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
		case "create-hosted":
			b.CreateHosted(*repoName, *blobStoreName, *repoFormat, *dockerHttpPort, *dockerHttpsPort, *releases)
		case "create-proxy":
			b.CreateProxy(*repoName, *blobStoreName, *repoFormat, *remoteURL, *proxyUser, *proxyPass, *dockerHttpPort, *dockerHttpsPort, *releases)
		case "create-group":
			b.CreateGroup(*repoName, *blobStoreName, *repoFormat, *repoMembers, *dockerHttpPort, *dockerHttpsPort, *releases)
		case "add-group-members":
			b.AddMembersToGroup(*repoName, *repoFormat, *repoMembers)
		case "delete":
			b.DeleteRepository(*repoName)
		default:
			repoCommand.Usage()
			log.Printf(b.TaskNotValidInfo, *repoTask, "repo", b.RepoTasks)
			os.Exit(1)
		}
	}
}
