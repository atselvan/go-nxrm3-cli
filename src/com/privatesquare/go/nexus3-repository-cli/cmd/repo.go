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
	Use:     RepoCommandFlag,
	Short:   RepoCommandUsage,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		nxrm.SetConnectionDetails()
		nxrm.SkipTLSVerification, _ = cmd.Flags().GetBool(SkipTlsFlag)
		nxrm.Debug, _ = cmd.Flags().GetBool(DebugFlag)
		nxrm.Verbose, _ = cmd.Flags().GetBool(VerboseFlag)
	},
}

var listRepoCmd = &cobra.Command{
	Use:     "list",
	Short:   "List Repositories",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(RepoNameFlag)
		format, _ := cmd.Flags().GetString(RepoFormatFlag)
		nxrm.ListRepositories(name, format)
	},
}

var createRepoCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create Repositories",
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

		if rType == "" {
			log.Printf("Repository type is a required parameter. Available types : %+q", nxrm.RepoType)
		}

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

var deleteRepoCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete Repositories",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(RepoNameFlag)
		nxrm.DeleteRepository(name)
	},
}

func init(){
	repoCmd.AddCommand(listRepoCmd)
	repoCmd.AddCommand(createRepoCmd)
	repoCmd.AddCommand(deleteRepoCmd)

	listRepoCmd.Flags().String(RepoNameFlag, "", RepoNameUsage)
	listRepoCmd.Flags().String(RepoFormatFlag, "", fmt.Sprintf(RepoFormatUsage, nxrm.RepoFormats))

	createRepoCmd.Flags().String(RepoNameFlag, "", RepoNameUsage)
	createRepoCmd.Flags().String(RepoFormatFlag, "", fmt.Sprintf(RepoFormatUsage, nxrm.RepoFormats))
	createRepoCmd.Flags().String(RepoTypeFlag, "", fmt.Sprintf(RepoTypeUsage, nxrm.RepoFormats))
	createRepoCmd.Flags().String(RepoMembersFlag, "", RepoMembersUsage)
	createRepoCmd.Flags().String(RemoteURLFlag, "", RemoteURLUsage)
	createRepoCmd.Flags().String(ProxyUserFlag, "", ProxyUserUsage)
	createRepoCmd.Flags().String(ProxyPassFlag, "", ProxyPassUsage)
	createRepoCmd.Flags().String(DockerHttpPortFlag, "", DockerHttpPortUsage)
	createRepoCmd.Flags().Float64(DockerHttpsPortFlag, 0, DockerHttpsPortUsage)
	createRepoCmd.Flags().Float64(BlobStoreNameFlag, 0, BlobStoreNameUsage)
	createRepoCmd.Flags().Bool(ReleaseFlag, false, ReleaseUsage)

	deleteRepoCmd.Flags().String(RepoNameFlag, "", RepoNameUsage)
}
