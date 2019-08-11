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
	Use:     confCommandFlag,
	Short:   confCommandUsage,
	Example: "./nexus3-repository-cli configure --nexus-url http://nexus-domain --username user --password pass",
	Run: func(cmd *cobra.Command, args []string) {
		nexusURL, _ := cmd.Flags().GetString(nexusURLFlag)
		nexusUser, _ := cmd.Flags().GetString(nexusUsernameFlag)
		nexusPass, _ := cmd.Flags().GetString(nexusPasswordFlag)
		nxrm.NexusURL = nexusURL
		nxrm.AuthUser = nxrm.AuthUserStruct{Username: nexusUser, Password: nexusPass}
		nxrm.StoreConnectionDetails()
	},
}

func init() {
	confCmd.Flags().String(nexusURLFlag, "", nexusURLUsage)
	_ = confCmd.MarkFlagRequired(nexusURLFlag)
	confCmd.Flags().String(nexusUsernameFlag, "", nexusUsernameUsage)
	_ = confCmd.MarkFlagRequired(nexusUsernameFlag)
	confCmd.Flags().String(nexusPasswordFlag, "", nexusPasswordUsage)
	_ = confCmd.MarkFlagRequired(nexusPasswordFlag)
	confCmd.Flags().SortFlags = false
}
