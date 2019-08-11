/*
Copyright © 2019 atselvan
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any sub-commands
var rootCmd = &cobra.Command{
	Use:   "nexus3-repository-cli",
	Short: "Sonatype Nexus Repository Manager 3 CLI",
	Long: `CLI to interacts with Nexus repository Manager 3
via its API to administer the instance and to create nxrm components`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	cobra.EnableCommandSorting = false

	rootCmd.AddCommand(confCmd)
	rootCmd.AddCommand(scriptsCmd)
	rootCmd.AddCommand(repoCmd)
	rootCmd.AddCommand(selectorCmd)
	rootCmd.AddCommand(privilegeCmd)
	rootCmd.AddCommand(roleCmd)

	rootCmd.PersistentFlags().BoolP(skipTlsFlag, "k", false, skipTlsUsage)
	rootCmd.PersistentFlags().BoolP(debugFlag, "d", false, debugUsage)
	rootCmd.PersistentFlags().BoolP(verboseFlag, "v", false, verboseUsage)
}
