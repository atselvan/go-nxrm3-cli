/*
Copyright © 2019 atselvan
*/
package cmd

import (
	"github.com/atselvan/go-nxrm-lib"
	"github.com/spf13/cobra"
)

// confCmd represents the configure command
var confCmd = &cobra.Command{
	Use:     ConfCommandFlag,
	Short:   ConfCommandUsage,
	Example: "./nexus3-repository-cli configure --nexus-url http://nexus-domain --username user --password pass",
	Run: func(cmd *cobra.Command, args []string) {
		nexusURL, _ := cmd.Flags().GetString(NexusURLFlag)
		nexusUser, _ := cmd.Flags().GetString(NexusUsernameFlag)
		nexusPass, _ := cmd.Flags().GetString(NexusPasswordFlag)
		nxrm.NexusURL = nexusURL
		nxrm.AuthUser = nxrm.AuthUserStruct{Username: nexusUser, Password: nexusPass}
		nxrm.StoreConnectionDetails()
	},
}

func init() {
	confCmd.Flags().String(NexusURLFlag, "", NexusURLUsage)
	confCmd.MarkFlagRequired(NexusURLFlag)
	confCmd.Flags().String(NexusUsernameFlag, "", NexusUsernameUsage)
	confCmd.MarkFlagRequired(NexusUsernameFlag)
	confCmd.Flags().String(NexusPasswordFlag, "", NexusPasswordUsage)
	confCmd.MarkFlagRequired(NexusPasswordFlag)
	confCmd.Flags().SortFlags = false
}
