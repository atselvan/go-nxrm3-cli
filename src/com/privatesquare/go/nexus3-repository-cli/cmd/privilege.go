/*
Copyright © 2019 atselvan
*/
package cmd

import (
	"fmt"
	"github.com/atselvan/go-nxrm-lib"
	"github.com/spf13/cobra"
)

// privilegeCmd represents the privilege command
var privilegeCmd = &cobra.Command{
	Use:   privilegeCommandFlag,
	Short: privilegeCommandUsage,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		nxrm.SetConnectionDetails()
		nxrm.SkipTLSVerification, _ = cmd.Flags().GetBool(skipTlsFlag)
		nxrm.Debug, _ = cmd.Flags().GetBool(debugFlag)
		nxrm.Verbose, _ = cmd.Flags().GetBool(verboseFlag)
	},
}

// listPrivilegesCmd represents the privilege list command
var listPrivilegesCmd = &cobra.Command{
	Use:   listTask,
	Short: "List repository privileges",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(privilegeNameFlag)
		nxrm.ListPrivileges(name)
	},
}

// createPrivilegeCmd represents the privilege create command
var createPrivilegeCmd = &cobra.Command{
	Use:   createTask,
	Short: "Create a new repository privilege",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(privilegeNameFlag)
		desc, _ := cmd.Flags().GetString(descFlag)
		selector, _ := cmd.Flags().GetString(pSelectorNameFlag)
		repo, _ := cmd.Flags().GetString(pRepoNameFlag)
		action, _ := cmd.Flags().GetString(actionFlag)
		nxrm.CreatePrivilege(name, desc, selector, repo, action)
	},
}

// updatePrivilegeCmd represents the privilege update command
var updatePrivilegeCmd = &cobra.Command{
	Use:   updateTask,
	Short: "Update an existing repository privilege",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(privilegeNameFlag)
		desc, _ := cmd.Flags().GetString(descFlag)
		selector, _ := cmd.Flags().GetString(pSelectorNameFlag)
		repo, _ := cmd.Flags().GetString(pRepoNameFlag)
		action, _ := cmd.Flags().GetString(actionFlag)
		nxrm.UpdatePrivilege(name, desc, selector, repo, action)
	},
}

// deletePrivilegeCmd represents the privilege delete command
var deletePrivilegeCmd = &cobra.Command{
	Use:   deleteTask,
	Short: "Delete an existing repository privilege",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(privilegeNameFlag)
		nxrm.DeletePrivilege(name)
	},
}

func init() {
	privilegeCmd.AddCommand(listPrivilegesCmd)
	privilegeCmd.AddCommand(createPrivilegeCmd)
	privilegeCmd.AddCommand(updatePrivilegeCmd)
	privilegeCmd.AddCommand(deletePrivilegeCmd)

	listPrivilegesCmd.Flags().String(privilegeNameFlag, "", privilegeNameUsage)

	createPrivilegeCmd.Flags().String(privilegeNameFlag, "", privilegeNameUsage)
	_ = createPrivilegeCmd.MarkFlagRequired(privilegeNameFlag)
	createPrivilegeCmd.Flags().String(descFlag, "", privilegeDescUsage)
	createPrivilegeCmd.Flags().String(pSelectorNameFlag, "", selectorNameUsage)
	_ = createPrivilegeCmd.MarkFlagRequired(pSelectorNameFlag)
	createPrivilegeCmd.Flags().String(pRepoNameFlag, "", repoNameUsage)
	_ = createPrivilegeCmd.MarkFlagRequired(pRepoNameFlag)
	createPrivilegeCmd.Flags().String(actionFlag, "", fmt.Sprintf(actionUsage, nxrm.PrivilegeActions))
	createPrivilegeCmd.Flags().SortFlags = false

	updatePrivilegeCmd.Flags().String(privilegeNameFlag, "", privilegeNameUsage)
	_ = updatePrivilegeCmd.MarkFlagRequired(privilegeNameFlag)
	updatePrivilegeCmd.Flags().String(descFlag, "", privilegeDescUsage)
	updatePrivilegeCmd.Flags().String(pSelectorNameFlag, "", selectorNameUsage)
	updatePrivilegeCmd.Flags().String(pRepoNameFlag, "", repoNameUsage)
	updatePrivilegeCmd.Flags().String(actionFlag, "", fmt.Sprintf(actionUsage, nxrm.PrivilegeActions))
	updatePrivilegeCmd.Flags().SortFlags = false

	deletePrivilegeCmd.Flags().String(privilegeNameFlag, "", privilegeNameUsage)
	_ = deletePrivilegeCmd.MarkFlagRequired(privilegeNameFlag)
}
