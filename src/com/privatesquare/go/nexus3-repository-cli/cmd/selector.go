/*
Copyright © 2019 atselvan
*/
package cmd

import (
	"github.com/atselvan/go-nxrm-lib"
	"github.com/spf13/cobra"
)

// selectorCmd represents the selector commands
var selectorCmd = &cobra.Command{
	Use:   selectorCommandFlag,
	Short: selectorCommandUsage,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		nxrm.SetConnectionDetails()
		nxrm.SkipTLSVerification, _ = cmd.Flags().GetBool(skipTlsFlag)
		nxrm.Debug, _ = cmd.Flags().GetBool(debugFlag)
		nxrm.Verbose, _ = cmd.Flags().GetBool(verboseFlag)
	},
}

// listSelectorsCmd represents the selector list command
var listSelectorsCmd = &cobra.Command{
	Use:   listTask,
	Short: "List content selectors",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(selectorNameFlag)
		nxrm.ListSelectors(name)
	},
}

// createSelectorCmd represents the selector create command
var createSelectorCmd = &cobra.Command{
	Use:   createTask,
	Short: "Create a new content selector",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(selectorNameFlag)
		desc, _ := cmd.Flags().GetString(descFlag)
		exp, _ := cmd.Flags().GetString(selectorExpressionFlag)
		nxrm.CreateSelector(name, desc, exp)
	},
}

// updateSelectorCmd represents the selector update command
var updateSelectorCmd = &cobra.Command{
	Use:   updateTask,
	Short: "Update an existing content selector",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(selectorNameFlag)
		desc, _ := cmd.Flags().GetString(descFlag)
		exp, _ := cmd.Flags().GetString(selectorExpressionFlag)
		nxrm.UpdateSelector(name, desc, exp)
	},
}

// deleteSelectorCmd represents the selector delete command
var deleteSelectorCmd = &cobra.Command{
	Use:   deleteTask,
	Short: "Delete an existing content selector",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(selectorNameFlag)
		nxrm.DeleteSelector(name)
	},
}

func init() {
	selectorCmd.AddCommand(listSelectorsCmd)
	selectorCmd.AddCommand(createSelectorCmd)
	selectorCmd.AddCommand(updateSelectorCmd)
	selectorCmd.AddCommand(deleteSelectorCmd)

	listSelectorsCmd.Flags().String(selectorNameFlag, "", selectorNameUsage)

	createSelectorCmd.Flags().String(selectorNameFlag, "", selectorNameUsage)
	_ = createSelectorCmd.MarkFlagRequired(selectorNameFlag)
	createSelectorCmd.Flags().String(descFlag, "", selectorDescUsage)
	createSelectorCmd.Flags().String(selectorExpressionFlag, "", selectorExpressionUsage)
	_ = createSelectorCmd.MarkFlagRequired(selectorExpressionFlag)
	createSelectorCmd.Flags().SortFlags = false

	updateSelectorCmd.Flags().String(selectorNameFlag, "", selectorNameUsage)
	_ = updateSelectorCmd.MarkFlagRequired(selectorNameFlag)
	updateSelectorCmd.Flags().String(descFlag, "", selectorDescUsage)
	updateSelectorCmd.Flags().String(selectorExpressionFlag, "", selectorExpressionUsage)
	updateSelectorCmd.Flags().SortFlags = false

	deleteSelectorCmd.Flags().String(selectorNameFlag, "", selectorNameUsage)
	_ = deleteSelectorCmd.MarkFlagRequired(selectorNameFlag)
}
