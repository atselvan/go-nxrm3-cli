package cmd

func something() {
	/*

		// repo flags
		rcSkipTLS := repoCommand.Bool(SkipTlsFlag, false, SkipTlsUsage)
		rcDebug := repoCommand.Bool(DebugFlag, false, DebugUsage)
		rcVerbose := repoCommand.Bool(VerboseFlag, false, VerboseUsage)
		// selector flags
		selectorTask := selectorCommand.String(TaskFlag, "", "")
		selectorName := selectorCommand.String(SelectorNameFlag, "", SelectorNameUsage)
		selectorDesc := selectorCommand.String(DescFlag, "", SelectorDescUsage)
		expression := selectorCommand.String(SelectorExpressionFlag, "", SelectorExpressionUsage)
		csSkipTLS := selectorCommand.Bool(SkipTlsFlag, false, SkipTlsUsage)
		csDebug := selectorCommand.Bool(DebugFlag, false, DebugUsage)
		csVerbose := selectorCommand.Bool(VerboseFlag, false, VerboseUsage)
		//privilege flags
		privilegeTask := privilegeCommand.String(TaskFlag, "", PrivilegeTaskUsage)
		privilegeName := privilegeCommand.String(PrivilegeNameFlag, "", PrivilegeNameUsage)
		pSelectorName := privilegeCommand.String(PSelectorNameFlag, "", SelectorNameUsage)
		privilegeDesc := privilegeCommand.String(DescFlag, "", PrivilegeDescUsage)
		pRepoName := privilegeCommand.String(PRepoNameFlag, "", RepoNameUsage)
		action := privilegeCommand.String(ActionFlag, "", ActionUsage)
		pSkipTLS := privilegeCommand.Bool(SkipTlsFlag, false, SkipTlsUsage)
		pDebug := privilegeCommand.Bool(DebugFlag, false, DebugUsage)
		pVerbose := privilegeCommand.Bool(VerboseFlag, false, VerboseUsage)
		//role flags
		roleTask := roleCommand.String(TaskFlag, "", RoleTaskUsage)
		roleID := roleCommand.String(RoleIDFlag, "", RoleIDUsage)
		roleDesc := roleCommand.String(DescFlag, "", RoleDescUsage)
		roleMembers := roleCommand.String(RoleMembersFlag, "", RoleMembersUsage)
		rolePrivileges := roleCommand.String(RolePrivilegesFlag, "", RolePrivilegesUsage)
		updateAction := roleCommand.String(UpdateActionFlag, "", UpdateActionUsage)
		rSkipTLS := roleCommand.Bool(SkipTlsFlag, false, SkipTlsUsage)
		rDebug := roleCommand.Bool(DebugFlag, false, DebugUsage)
		rVerbose := roleCommand.Bool(VerboseFlag, false, VerboseUsage)


		if repoCommand.Parsed() {
			// Required flags
			if *repoTask == "" {
				repoCommand.Usage()
				log.Printf(TaskEmptyInfo, RepoTasks)
				os.Exit(1)
			}
			//Choice flag
			repoFormatChoice := map[string]bool{"": true}
			for _, repoFormat := range RepoFormats {
				repoFormatChoice[repoFormat] = true
			}
			if _, validChoice := repoFormatChoice[*repoFormat]; !validChoice {
				log.Printf(RepoFormatNotValidInfo, *repoFormat, RepoFormats)
				os.Exit(1)
			}
			// set global variables
			SetConnectionDetails()
			SkipTLSVerification = *rcSkipTLS
			Debug = *rcDebug
			Verbose = *rcVerbose
			// run tasks
			switch *repoTask {
			case ListTask:
				ListRepositories(*repoName, *repoFormat)
			case "create-hosted":
				CreateHosted(*repoName, *blobStoreName, *repoFormat, *dockerHttpPort, *dockerHttpsPort, *releases)
			case "create-proxy":
				CreateProxy(*repoName, *blobStoreName, *repoFormat, *remoteURL, *proxyUser, *proxyPass, *dockerHttpPort, *dockerHttpsPort, *releases)
			case "create-group":
				CreateGroup(*repoName, *blobStoreName, *repoFormat, *repoMembers, *dockerHttpPort, *dockerHttpsPort, *releases)
			case "add-group-members":
				AddMembersToGroup(*repoName, *repoFormat, *repoMembers)
			case "remove-group-members":
				RemoveMembersFromGroup(*repoName, *repoFormat, *repoMembers)
			case DeleteTask:
				DeleteRepository(*repoName)
			default:
				repoCommand.Usage()
				log.Printf(TaskNotValidInfo, *repoTask, "repo", RepoTasks)
				os.Exit(1)
			}
		}

		if selectorCommand.Parsed() {
			// Required flags
			if *selectorTask == "" {
				selectorCommand.Usage()
				log.Printf(TaskEmptyInfo, SelectorTasks)
				os.Exit(1)
			}
			// set global variables
			SetConnectionDetails()
			SkipTLSVerification = *csSkipTLS
			Debug = *csDebug
			Verbose = *csVerbose
			// run tasks
			switch *selectorTask {
			case ListTask:
				ListSelectors(*selectorName)
			case CreateTask:
				CreateSelector(*selectorName, *selectorDesc, *expression)
			case UpdateTask:
				UpdateSelector(*selectorName, *selectorDesc, *expression)
			case DeleteTask:
				DeleteSelector(*selectorName)
			default:
				selectorCommand.Usage()
				log.Printf(TaskNotValidInfo, *selectorTask, "selector", SelectorTasks)
				os.Exit(1)
			}
		}

		if privilegeCommand.Parsed() {
			// Required flags
			if *privilegeTask == "" {
				privilegeCommand.Usage()
				log.Printf(TaskEmptyInfo, PrivilegeTasks)
				os.Exit(1)
			}
			// set global variables
			SetConnectionDetails()
			SkipTLSVerification = *pSkipTLS
			Debug = *pDebug
			Verbose = *pVerbose
			// run tasks
			switch *privilegeTask {
			case ListTask:
				ListPrivileges(*privilegeName)
			case CreateTask:
				CreatePrivilege(*privilegeName, *privilegeDesc, *pSelectorName, *pRepoName, *action)
			case UpdateTask:
				UpdatePrivilege(*privilegeName, *privilegeDesc, *pSelectorName, *pRepoName, *action)
			case DeleteTask:
				DeletePrivilege(*privilegeName)
			default:
				privilegeCommand.Usage()
				log.Printf(TaskNotValidInfo, *privilegeTask, "privilege", PrivilegeTasks)
				os.Exit(1)
			}
		}

		if roleCommand.Parsed() {
			// Required flags
			if *roleTask == "" {
				roleCommand.Usage()
				log.Printf(TaskEmptyInfo, RoleTasks)
				os.Exit(1)
			}
			// set global variables
			SetConnectionDetails()
			SkipTLSVerification = *rSkipTLS
			Debug = *rDebug
			Verbose = *rVerbose
			// run tasks
			switch *roleTask {
			case ListTask:
				ListRoles(*roleID)
			case CreateTask:
				CreateRole(*roleID, *roleDesc, *roleMembers, *rolePrivileges)
			case UpdateTask:
				UpdateRole(*roleID, *roleDesc, *roleMembers, *rolePrivileges, *updateAction)
			case DeleteTask:
				DeleteRole(*roleID)
			default:
				roleCommand.Usage()
				log.Printf(TaskNotValidInfo, *roleTask, "role", RoleTasks)
				os.Exit(1)
			}
		}
	*/
}
