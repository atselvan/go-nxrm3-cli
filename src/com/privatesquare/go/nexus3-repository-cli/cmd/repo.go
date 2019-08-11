/*
Copyright © 2019 atselvan
*/
package cmd

import (
	"fmt"
	"github.com/atselvan/go-nxrm-lib"
	"github.com/spf13/cobra"
	"log"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   repoCommandFlag,
	Short: repoCommandUsage,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		nxrm.SetConnectionDetails()
		nxrm.SkipTLSVerification, _ = cmd.Flags().GetBool(skipTlsFlag)
		nxrm.Debug, _ = cmd.Flags().GetBool(debugFlag)
		nxrm.Verbose, _ = cmd.Flags().GetBool(verboseFlag)
	},
}

// lostRepoCmd represents the repo list command
var listRepoCmd = &cobra.Command{
	Use:   listTask,
	Short: "List Repositories",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(repoNameFlag)
		format, _ := cmd.Flags().GetString(repoFormatFlag)
		nxrm.ListRepositories(name, format)
	},
}

// createRepoCmd represents the repo create command
var createRepoCmd = &cobra.Command{
	Use:   createTask,
	Short: "Create Repositories",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(repoNameFlag)
		format, _ := cmd.Flags().GetString(repoFormatFlag)
		rType, _ := cmd.Flags().GetString(repoTypeFlag)
		members, _ := cmd.Flags().GetString(repoMembersFlag)
		remoteURL, _ := cmd.Flags().GetString(remoteURLFlag)
		proxyUser, _ := cmd.Flags().GetString(proxyUserFlag)
		proxyPass, _ := cmd.Flags().GetString(proxyPassFlag)
		dockerHttpPort, _ := cmd.Flags().GetFloat64(dockerHttpPortFlag)
		dockerHttpsPort, _ := cmd.Flags().GetFloat64(dockerHttpsPortFlag)
		blobStoreName, _ := cmd.Flags().GetString(blobStoreNameFlag)
		releases, _ := cmd.Flags().GetBool(releaseFlag)

		switch rType {
		case "hosted":
			nxrm.CreateHosted(name, blobStoreName, format, dockerHttpPort, dockerHttpsPort, releases)
		case "proxy":
			nxrm.CreateProxy(name, blobStoreName, format, remoteURL, proxyUser, proxyPass, dockerHttpPort, dockerHttpsPort, releases)
		case "group":
			nxrm.CreateGroup(name, blobStoreName, format, members, dockerHttpPort, dockerHttpsPort, releases)
		default:
			log.Printf("Invalid repository type. Available types : %+q", nxrm.RepoType)
		}
	},
}

// addMembersCmd represents the repo addMembersToGroup command
var addMembersCmd = &cobra.Command{
	Use:   "addMembers",
	Short: "Add more members to an existing group repository",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(repoNameFlag)
		members, _ := cmd.Flags().GetString(repoMembersFlag)
		nxrm.AddMembersToGroup(name, members)
	},
}

// removeMembersCmd represents the repo removeMembersFromGroup command
var removeMembersCmd = &cobra.Command{
	Use:   "removeMembers",
	Short: "Remove members from an existing group repository",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(repoNameFlag)
		members, _ := cmd.Flags().GetString(repoMembersFlag)
		nxrm.RemoveMembersFromGroup(name, members)
	},
}

// deleteRepoCmd represents the repo delete command
var deleteRepoCmd = &cobra.Command{
	Use:   deleteTask,
	Short: "Delete Repositories",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(repoNameFlag)
		nxrm.DeleteRepository(name)
	},
}

func init() {
	repoCmd.AddCommand(listRepoCmd)
	repoCmd.AddCommand(createRepoCmd)
	repoCmd.AddCommand(addMembersCmd)
	repoCmd.AddCommand(removeMembersCmd)
	repoCmd.AddCommand(deleteRepoCmd)

	listRepoCmd.Flags().String(repoNameFlag, "", repoNameUsage)
	listRepoCmd.Flags().String(repoFormatFlag, "", fmt.Sprintf(repoFormatUsage, nxrm.RepoFormats))
	listRepoCmd.Flags().SortFlags = false

	createRepoCmd.Flags().String(repoNameFlag, "", repoNameUsage)
	_ = createRepoCmd.MarkFlagRequired(repoNameFlag)
	createRepoCmd.Flags().String(repoTypeFlag, "", fmt.Sprintf(repoTypeUsage, nxrm.RepoType))
	_ = createRepoCmd.MarkFlagRequired(repoTypeFlag)
	createRepoCmd.Flags().String(repoFormatFlag, "", fmt.Sprintf(repoFormatUsage, nxrm.RepoFormats))
	_ = createRepoCmd.MarkFlagRequired(repoFormatFlag)
	createRepoCmd.Flags().String(remoteURLFlag, "", remoteURLUsage)
	createRepoCmd.Flags().String(proxyUserFlag, "", proxyUserUsage)
	createRepoCmd.Flags().String(proxyPassFlag, "", proxyPassUsage)
	createRepoCmd.Flags().String(dockerHttpPortFlag, "", dockerHttpPortUsage)
	createRepoCmd.Flags().Float64(dockerHttpsPortFlag, 0, dockerHttpsPortUsage)
	createRepoCmd.Flags().Float64(blobStoreNameFlag, 0, blobStoreNameUsage)
	createRepoCmd.Flags().Bool(releaseFlag, false, releaseUsage)
	createRepoCmd.Flags().SortFlags = false

	addMembersCmd.Flags().String(repoNameFlag, "", repoNameUsage)
	_ = addMembersCmd.MarkFlagRequired(repoNameFlag)
	addMembersCmd.Flags().String(repoMembersFlag, "", repoMembersUsage)
	_ = addMembersCmd.MarkFlagRequired(repoMembersFlag)
	addMembersCmd.Flags().SortFlags = false

	removeMembersCmd.Flags().String(repoNameFlag, "", repoNameUsage)
	_ = removeMembersCmd.MarkFlagRequired(repoNameFlag)
	removeMembersCmd.Flags().String(repoMembersFlag, "", repoMembersUsage)
	_ = removeMembersCmd.MarkFlagRequired(repoMembersFlag)
	removeMembersCmd.Flags().SortFlags = false

	deleteRepoCmd.Flags().String(repoNameFlag, "", repoNameUsage)
	_ = deleteRepoCmd.MarkFlagRequired(repoNameFlag)
}
