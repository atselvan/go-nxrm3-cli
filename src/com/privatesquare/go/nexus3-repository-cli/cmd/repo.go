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
	Use:   RepoCommandFlag,
	Short: RepoCommandUsage,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		nxrm.SetConnectionDetails()
		nxrm.SkipTLSVerification, _ = cmd.Flags().GetBool(SkipTlsFlag)
		nxrm.Debug, _ = cmd.Flags().GetBool(DebugFlag)
		nxrm.Verbose, _ = cmd.Flags().GetBool(VerboseFlag)
	},
}

// lostRepoCmd represents the repo list command
var listRepoCmd = &cobra.Command{
	Use:   "list",
	Short: "List Repositories",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(RepoNameFlag)
		format, _ := cmd.Flags().GetString(RepoFormatFlag)
		nxrm.ListRepositories(name, format)
	},
}

// createRepoCmd represents the repo create command
var createRepoCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Repositories",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(RepoNameFlag)
		format, _ := cmd.Flags().GetString(RepoFormatFlag)
		rType, _ := cmd.Flags().GetString(RepoTypeFlag)
		members, _ := cmd.Flags().GetString(RepoMembersFlag)
		remoteURL, _ := cmd.Flags().GetString(RemoteURLFlag)
		proxyUser, _ := cmd.Flags().GetString(ProxyUserFlag)
		proxyPass, _ := cmd.Flags().GetString(ProxyPassFlag)
		dockerHttpPort, _ := cmd.Flags().GetFloat64(DockerHttpPortFlag)
		dockerHttpsPort, _ := cmd.Flags().GetFloat64(DockerHttpsPortFlag)
		blobStoreName, _ := cmd.Flags().GetString(BlobStoreNameFlag)
		releases, _ := cmd.Flags().GetBool(ReleaseFlag)

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
		name, _ := cmd.Flags().GetString(RepoNameFlag)
		members, _ := cmd.Flags().GetString(RepoMembersFlag)
		nxrm.AddMembersToGroup(name, members)
	},
}

// removeMembersCmd represents the repo removeMembersFromGroup command
var removeMembersCmd = &cobra.Command{
	Use:   "removeMembers",
	Short: "Remove members from an existing group repository",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(RepoNameFlag)
		members, _ := cmd.Flags().GetString(RepoMembersFlag)
		nxrm.RemoveMembersFromGroup(name, members)
	},
}

// deleteRepoCmd represents the repo delete command
var deleteRepoCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Repositories",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(RepoNameFlag)
		nxrm.DeleteRepository(name)
	},
}

func init() {
	repoCmd.AddCommand(listRepoCmd)
	repoCmd.AddCommand(createRepoCmd)
	repoCmd.AddCommand(addMembersCmd)
	repoCmd.AddCommand(removeMembersCmd)
	repoCmd.AddCommand(deleteRepoCmd)

	listRepoCmd.Flags().String(RepoNameFlag, "", RepoNameUsage)
	listRepoCmd.Flags().String(RepoFormatFlag, "", fmt.Sprintf(RepoFormatUsage, nxrm.RepoFormats))
	listRepoCmd.Flags().SortFlags = false

	createRepoCmd.Flags().String(RepoNameFlag, "", RepoNameUsage)
	createRepoCmd.MarkFlagRequired(RepoNameFlag)
	createRepoCmd.Flags().String(RepoFormatFlag, "", fmt.Sprintf(RepoFormatUsage, nxrm.RepoFormats))
	createRepoCmd.MarkFlagRequired(RepoFormatFlag)
	createRepoCmd.Flags().String(RepoTypeFlag, "", fmt.Sprintf(RepoTypeUsage, nxrm.RepoFormats))
	createRepoCmd.MarkFlagRequired(RepoTypeFlag)
	createRepoCmd.Flags().String(RemoteURLFlag, "", RemoteURLUsage)
	createRepoCmd.Flags().String(ProxyUserFlag, "", ProxyUserUsage)
	createRepoCmd.Flags().String(ProxyPassFlag, "", ProxyPassUsage)
	createRepoCmd.Flags().String(DockerHttpPortFlag, "", DockerHttpPortUsage)
	createRepoCmd.Flags().Float64(DockerHttpsPortFlag, 0, DockerHttpsPortUsage)
	createRepoCmd.Flags().Float64(BlobStoreNameFlag, 0, BlobStoreNameUsage)
	createRepoCmd.Flags().Bool(ReleaseFlag, false, ReleaseUsage)
	createRepoCmd.Flags().SortFlags = false

	addMembersCmd.Flags().String(RepoNameFlag, "", RepoNameUsage)
	addMembersCmd.MarkFlagRequired(RepoNameFlag)
	addMembersCmd.Flags().String(RepoMembersFlag, "", RepoMembersUsage)
	addMembersCmd.MarkFlagRequired(RepoMembersFlag)
	addMembersCmd.Flags().SortFlags = false

	removeMembersCmd.Flags().String(RepoNameFlag, "", RepoNameUsage)
	removeMembersCmd.MarkFlagRequired(RepoNameFlag)
	removeMembersCmd.Flags().String(RepoMembersFlag, "", RepoMembersUsage)
	removeMembersCmd.MarkFlagRequired(RepoMembersFlag)
	removeMembersCmd.Flags().SortFlags = false

	deleteRepoCmd.Flags().String(RepoNameFlag, "", RepoNameUsage)
}
