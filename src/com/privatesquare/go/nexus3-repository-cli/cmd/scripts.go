/*
Copyright © 2019 atselvan
*/
package cmd

import (
	"github.com/atselvan/go-nxrm-lib"
	"github.com/spf13/cobra"
)

// scriptCmd represents the script command
var scriptsCmd = &cobra.Command{
	Use:   ScriptCommandFlag,
	Short: ScriptCommandUsage,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		nxrm.SetConnectionDetails()
		nxrm.SkipTLSVerification, _ = cmd.Flags().GetBool(SkipTlsFlag)
		nxrm.Debug = true
		nxrm.Verbose, _ = cmd.Flags().GetBool(VerboseFlag)
	},
}

var initScriptsCmd = &cobra.Command{
	Use:   "init",
	Short: "Init Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		nxrm.ScriptsInit()
	},
}

var listScriptsCmd = &cobra.Command{
	Use:   "list",
	Short: "List Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		scriptName, _ := cmd.Flags().GetString(ScriptNameFlag)
		nxrm.ListScripts(scriptName)
	},
}

var addScriptsCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		scriptName, _ := cmd.Flags().GetString(ScriptNameFlag)
		nxrm.AddScript(scriptName)
	},
}

var updateScriptsCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		scriptName, _ := cmd.Flags().GetString(ScriptNameFlag)
		nxrm.UpdateScript(scriptName)
	},
}

var addOrUpdateScriptsCmd = &cobra.Command{
	Use:   "add-or-update",
	Short: "Add or Update Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		scriptName, _ := cmd.Flags().GetString(ScriptNameFlag)
		nxrm.AddOrUpdateScript(scriptName)
	},
}

var deleteScriptsCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		scriptName, _ := cmd.Flags().GetString(ScriptNameFlag)
		nxrm.DeleteScript(scriptName)
	},
}

var runScriptsCmd = &cobra.Command{
	Use:   "run",
	Short: "Run Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		scriptName, _ := cmd.Flags().GetString(ScriptNameFlag)
		payload, _ := cmd.Flags().GetString(ScriptPayloadFlag)
		nxrm.RunScript(scriptName, payload)
	},
}

func init() {
	scriptsCmd.AddCommand(initScriptsCmd)
	scriptsCmd.AddCommand(listScriptsCmd)
	scriptsCmd.AddCommand(addScriptsCmd)
	scriptsCmd.AddCommand(updateScriptsCmd)
	scriptsCmd.AddCommand(addOrUpdateScriptsCmd)
	scriptsCmd.AddCommand(deleteScriptsCmd)
	scriptsCmd.AddCommand(runScriptsCmd)

	scriptsCmd.PersistentFlags().String(ScriptNameFlag, "", ScriptNameUsage)
	scriptsCmd.PersistentFlags().String(ScriptPayloadFlag, "", ScriptPayloadUsage)
}
