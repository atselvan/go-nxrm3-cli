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
	selectorCommand := flag.NewFlagSet(b.SelectorCommandFlag, flag.ExitOnError)
	privilegeCommand := flag.NewFlagSet(b.PrivilegeCommandFlag, flag.ExitOnError)
	roleCommand := flag.NewFlagSet(b.RoleCommandFlag, flag.ExitOnError)
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
	dockerHttpPort := repoCommand.Float64(b.DockerHttpPortFlag, 0, b.DockerHttpPortUsage)
	dockerHttpsPort := repoCommand.Float64(b.DockerHttpsPortFlag, 0, b.DockerHttpsPortUsage)
	blobStoreName := repoCommand.String(b.BlobStoreNameFlag, "", b.BlobStoreNameUsage)
	releases := repoCommand.Bool(b.ReleaseFlag, false, b.ReleaseUsage)
	rcSkipTLS := repoCommand.Bool(b.SkipTlsFlag, false, b.SkipTlsUsage)
	rcDebug := repoCommand.Bool(b.DebugFlag, false, b.DebugUsage)
	rcVerbose := repoCommand.Bool(b.VerboseFlag, false, b.VerboseUsage)
	// selector flags
	selectorTask := selectorCommand.String(b.TaskFlag, "", "")
	selectorName := selectorCommand.String(b.SelectorNameFlag, "", b.SelectorNameUsage)
	selectorDesc := selectorCommand.String(b.DescFlag, "", b.SelectorDescUsage)
	expression := selectorCommand.String(b.SelectorExpressionFlag, "", b.SelectorExpressionUsage)
	csSkipTLS := selectorCommand.Bool(b.SkipTlsFlag, false, b.SkipTlsUsage)
	csDebug := selectorCommand.Bool(b.DebugFlag, false, b.DebugUsage)
	csVerbose := selectorCommand.Bool(b.VerboseFlag, false, b.VerboseUsage)
	//privilege flags
	privilegeTask := privilegeCommand.String(b.TaskFlag, "", b.PrivilegeTaskUsage)
	privilegeName := privilegeCommand.String(b.PrivilegeNameFlag, "", b.PrivilegeNameUsage)
	pSelectorName := privilegeCommand.String(b.PSelectorNameFlag, "", b.SelectorNameUsage)
	privilegeDesc := privilegeCommand.String(b.DescFlag, "", b.PrivilegeDescUsage)
	pRepoName := privilegeCommand.String(b.PRepoNameFlag, "", b.RepoNameUsage)
	action := privilegeCommand.String(b.ActionFlag, "", b.ActionUsage)
	pSkipTLS := privilegeCommand.Bool(b.SkipTlsFlag, false, b.SkipTlsUsage)
	pDebug := privilegeCommand.Bool(b.DebugFlag, false, b.DebugUsage)
	pVerbose := privilegeCommand.Bool(b.VerboseFlag, false, b.VerboseUsage)
	//role flags
	roleTask := roleCommand.String(b.TaskFlag, "", b.RoleTaskUsage)
	roleID := roleCommand.String(b.RoleIDFlag, "", b.RoleIDUsage)
	roleDesc := roleCommand.String(b.DescFlag, "", b.RoleDescUsage)
	roleMembers := roleCommand.String(b.RoleMembersFlag, "", b.RoleMembersUsage)
	rolePrivileges := roleCommand.String(b.RolePrivilegesFlag, "", b.RolePrivilegesUsage)
	updateAction := roleCommand.String(b.UpdateActionFlag, "", b.UpdateActionUsage)
	rSkipTLS := roleCommand.Bool(b.SkipTlsFlag, false, b.SkipTlsUsage)
	rDebug := roleCommand.Bool(b.DebugFlag, false, b.DebugUsage)
	rVerbose := roleCommand.Bool(b.VerboseFlag, false, b.VerboseUsage)

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
	case "selector":
		b.PrintSelectorCommandUsage(selectorCommand)
		selectorCommand.Parse(os.Args[2:])
	case "privilege":
		b.PrintPrivilegeCommandUsage(privilegeCommand)
		privilegeCommand.Parse(os.Args[2:])
	case "role":
		b.PrintRoleCommandUsage(roleCommand)
		roleCommand.Parse(os.Args[2:])
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
		// Required flags
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
		case b.ListTask:
			b.ListScripts(*scriptName)
		case b.AddTask:
			b.Debug = true
			b.AddScript(*scriptName)
		case b.UpdateTask:
			b.Debug = true
			b.UpdateScript(*scriptName)
		case b.AddOrUpdateTask:
			b.Debug = true
			b.AddOrUpdateScript(*scriptName)
		case b.DeleteTask:
			b.Debug = true
			b.DeleteScript(*scriptName)
		case b.RunTask:
			b.Debug = true
			b.RunScript(*scriptName, *scriptPayload)
		default:
			scriptCommand.Usage()
			log.Printf(b.TaskNotValidInfo, *scriptTask, "script", b.ScriptTasks)
			os.Exit(1)
		}
	}

	if repoCommand.Parsed() {
		// Required flags
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
		case b.ListTask:
			b.ListRepositories(*repoName, *repoFormat)
		case "create-hosted":
			b.CreateHosted(*repoName, *blobStoreName, *repoFormat, *dockerHttpPort, *dockerHttpsPort, *releases)
		case "create-proxy":
			b.CreateProxy(*repoName, *blobStoreName, *repoFormat, *remoteURL, *proxyUser, *proxyPass, *dockerHttpPort, *dockerHttpsPort, *releases)
		case "create-group":
			b.CreateGroup(*repoName, *blobStoreName, *repoFormat, *repoMembers, *dockerHttpPort, *dockerHttpsPort, *releases)
		case "add-group-members":
			b.AddMembersToGroup(*repoName, *repoFormat, *repoMembers)
		case "remove-group-members":
			b.RemoveMembersFromGroup(*repoName, *repoFormat, *repoMembers)
		case b.DeleteTask:
			b.DeleteRepository(*repoName)
		default:
			repoCommand.Usage()
			log.Printf(b.TaskNotValidInfo, *repoTask, "repo", b.RepoTasks)
			os.Exit(1)
		}
	}

	if selectorCommand.Parsed() {
		// Required flags
		if *selectorTask == "" {
			selectorCommand.Usage()
			log.Printf(b.TaskEmptyInfo, b.SelectorTasks)
			os.Exit(1)
		}
		// set global variables
		b.SetConnectionDetails()
		b.SkipTLSVerification = *csSkipTLS
		b.Debug = *csDebug
		b.Verbose = *csVerbose
		// run tasks
		switch *selectorTask {
		case b.ListTask:
			b.ListSelectors(*selectorName)
		case b.CreateTask:
			b.CreateSelector(*selectorName, *selectorDesc, *expression)
		case b.UpdateTask:
			b.UpdateSelector(*selectorName, *selectorDesc, *expression)
		case b.DeleteTask:
			b.DeleteSelector(*selectorName)
		default:
			selectorCommand.Usage()
			log.Printf(b.TaskNotValidInfo, *selectorTask, "selector", b.SelectorTasks)
			os.Exit(1)
		}
	}

	if privilegeCommand.Parsed() {
		// Required flags
		if *privilegeTask == "" {
			privilegeCommand.Usage()
			log.Printf(b.TaskEmptyInfo, b.PrivilegeTasks)
			os.Exit(1)
		}
		// set global variables
		b.SetConnectionDetails()
		b.SkipTLSVerification = *pSkipTLS
		b.Debug = *pDebug
		b.Verbose = *pVerbose
		// run tasks
		switch *privilegeTask {
		case b.ListTask:
			b.ListPrivileges(*privilegeName)
		case b.CreateTask:
			b.CreatePrivilege(*privilegeName, *privilegeDesc, *pSelectorName, *pRepoName, *action)
		case b.UpdateTask:
			b.UpdatePrivilege(*privilegeName, *privilegeDesc, *pSelectorName, *pRepoName, *action)
		case b.DeleteTask:
			b.DeletePrivilege(*privilegeName)
		default:
			privilegeCommand.Usage()
			log.Printf(b.TaskNotValidInfo, *privilegeTask, "privilege", b.PrivilegeTasks)
			os.Exit(1)
		}
	}

	if roleCommand.Parsed() {
		// Required flags
		if *roleTask == "" {
			roleCommand.Usage()
			log.Printf(b.TaskEmptyInfo, b.RoleTasks)
			os.Exit(1)
		}
		// set global variables
		b.SetConnectionDetails()
		b.SkipTLSVerification = *rSkipTLS
		b.Debug = *rDebug
		b.Verbose = *rVerbose
		// run tasks
		switch *roleTask {
		case b.ListTask:
			b.ListRoles(*roleID)
		case b.CreateTask:
			b.CreateRole(*roleID, *roleDesc, *roleMembers, *rolePrivileges)
		case b.UpdateTask:
			b.UpdateRole(*roleID, *roleDesc, *roleMembers, *rolePrivileges, *updateAction)
		case b.DeleteTask:
			b.DeleteRole(*roleID)
		default:
			roleCommand.Usage()
			log.Printf(b.TaskNotValidInfo, *roleTask, "role", b.RoleTasks)
			os.Exit(1)
		}
	}
}
