/*
Copyright © 2019 atselvan
*/
package cmd

import (
	"fmt"
	"github.com/atselvan/go-nxrm-lib"
	"github.com/spf13/cobra"
)

// roleCmd represents the role command
var roleCmd = &cobra.Command{
	Use:   roleCommandFlag,
	Short: roleCommandUsage,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		nxrm.SetConnectionDetails()
		nxrm.SkipTLSVerification, _ = cmd.Flags().GetBool(skipTlsFlag)
		nxrm.Debug, _ = cmd.Flags().GetBool(debugFlag)
		nxrm.Verbose, _ = cmd.Flags().GetBool(verboseFlag)
	},
}

// listRolesCmd represents the role list command
var listRolesCmd = &cobra.Command{
	Use:   listTask,
	Short: "List roles",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(roleIDFlag)
		nxrm.ListRoles(id)
	},
}

// createRoleCmd represents the role create command
var createRoleCmd = &cobra.Command{
	Use:   createTask,
	Short: "Create a new role",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(roleIDFlag)
		desc, _ := cmd.Flags().GetString(descFlag)
		roleMembers, _ := cmd.Flags().GetString(roleMembersFlag)
		rolePrivileges, _ := cmd.Flags().GetString(rolePrivilegesFlag)
		nxrm.CreateRole(id, desc, roleMembers, rolePrivileges)
	},
}

// updateRoleCmd represents the role update command
var updateRoleCmd = &cobra.Command{
	Use:   updateTask,
	Short: "Update an existing role",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(roleIDFlag)
		desc, _ := cmd.Flags().GetString(descFlag)
		roleMembers, _ := cmd.Flags().GetString(roleMembersFlag)
		rolePrivileges, _ := cmd.Flags().GetString(rolePrivilegesFlag)
		updateAction, _ := cmd.Flags().GetString(updateActionFlag)
		nxrm.UpdateRole(id, desc, roleMembers, rolePrivileges, updateAction)
	},
}

// deleteRoleCmd represents the role delete command
var deleteRoleCmd = &cobra.Command{
	Use:   deleteTask,
	Short: "Delete an existing role",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(roleIDFlag)
		nxrm.DeleteRole(id)
	},
}

func init() {
	roleCmd.AddCommand(listRolesCmd)
	roleCmd.AddCommand(createRoleCmd)
	roleCmd.AddCommand(updateRoleCmd)
	roleCmd.AddCommand(deleteRoleCmd)

	listRolesCmd.Flags().String(roleIDFlag, "", roleIDUsage)

	createRoleCmd.Flags().String(roleIDFlag, "", roleIDUsage)
	_ = createRoleCmd.MarkFlagRequired(roleIDFlag)
	createRoleCmd.Flags().String(descFlag, "", roleDescUsage)
	createRoleCmd.Flags().String(roleMembersFlag, "", roleMembersUsage)
	createRoleCmd.Flags().String(rolePrivilegesFlag, "", rolePrivilegesUsage)
	createRoleCmd.Flags().SortFlags = false

	updateRoleCmd.Flags().String(roleIDFlag, "", roleIDUsage)
	_ = updateRoleCmd.MarkFlagRequired(roleIDFlag)
	updateRoleCmd.Flags().String(descFlag, "", roleDescUsage)
	updateRoleCmd.Flags().String(roleMembersFlag, "", roleMembersUsage)
	updateRoleCmd.Flags().String(rolePrivilegesFlag, "", rolePrivilegesUsage)
	updateRoleCmd.Flags().String(updateActionFlag, "", fmt.Sprintf(updateActionUsage, nxrm.UpdateActions))
	_ = updateRoleCmd.MarkFlagRequired(updateActionFlag)
	updateRoleCmd.Flags().SortFlags = false

	deleteRoleCmd.Flags().String(roleIDFlag, "", roleIDUsage)
	_ = deleteRoleCmd.MarkFlagRequired(roleIDFlag)
}
