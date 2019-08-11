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
	Use:   scriptCommandFlag,
	Short: scriptCommandUsage,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		nxrm.SetConnectionDetails()
		nxrm.SkipTLSVerification, _ = cmd.Flags().GetBool(skipTlsFlag)
		nxrm.Debug = true
		nxrm.Verbose, _ = cmd.Flags().GetBool(verboseFlag)
	},
}

var initScriptsCmd = &cobra.Command{
	Use:   initTask,
	Short: "Init Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		nxrm.ScriptsInit()
	},
}

var listScriptsCmd = &cobra.Command{
	Use:   listTask,
	Short: "List Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		scriptName, _ := cmd.Flags().GetString(scriptNameFlag)
		nxrm.ListScripts(scriptName)
	},
}

var addScriptsCmd = &cobra.Command{
	Use:   addTask,
	Short: "Add Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		scriptName, _ := cmd.Flags().GetString(scriptNameFlag)
		nxrm.AddScript(scriptName)
	},
}

var updateScriptsCmd = &cobra.Command{
	Use:   updateTask,
	Short: "Update Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		scriptName, _ := cmd.Flags().GetString(scriptNameFlag)
		nxrm.UpdateScript(scriptName)
	},
}

var addOrUpdateScriptsCmd = &cobra.Command{
	Use:   addOrUpdateTask,
	Short: "Add or Update Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		scriptName, _ := cmd.Flags().GetString(scriptNameFlag)
		nxrm.AddOrUpdateScript(scriptName)
	},
}

var deleteScriptsCmd = &cobra.Command{
	Use:   deleteTask,
	Short: "Delete Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		scriptName, _ := cmd.Flags().GetString(scriptNameFlag)
		nxrm.DeleteScript(scriptName)
	},
}

var runScriptsCmd = &cobra.Command{
	Use:   runTask,
	Short: "Run Scripts",
	Run: func(cmd *cobra.Command, args []string) {
		scriptName, _ := cmd.Flags().GetString(scriptNameFlag)
		payload, _ := cmd.Flags().GetString(scriptPayloadFlag)
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

	scriptsCmd.PersistentFlags().String(scriptNameFlag, "", scriptNameUsage)
	_ = scriptsCmd.MarkPersistentFlagRequired(scriptNameFlag)
	scriptsCmd.PersistentFlags().String(scriptPayloadFlag, "", scriptPayloadUsage)
}
