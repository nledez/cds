package project

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ovh/cds/sdk"
)

// CmdVariable Command to manage variable on project
var CmdVariable = &cobra.Command{
	Use:     "variable",
	Short:   "",
	Long:    ``,
	Aliases: []string{"v"},
}

var force *bool

func init() {
	CmdVariable.AddCommand(cmdProjectUpdateVariable())
	CmdVariable.AddCommand(cmdProjectRemoveVariable())
}

func cmdProjectUpdateVariable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "cds project variable update <projectKey> <oldVariableName> <variableName> <variableValue> <variableType>",
		Long:  ``,
		Run:   updateVarInProject,
	}
	return cmd
}

func updateVarInProject(cmd *cobra.Command, args []string) {
	if len(args) != 5 {
		sdk.Exit("Wrong usage: %s\n", cmd.Short)
	}
	projectKey := args[0]
	oldName := args[1]
	varName := args[2]
	varValue := args[3]
	varType := args[4]

	err := sdk.UpdateVariableInProject(projectKey, oldName, varName, varValue, varType)
	if err != nil {
		sdk.Exit("Error: cannot update variable %s in project %s (%s)\n", varName, projectKey, err)
	}
	fmt.Printf("OK\n")
}

func cmdProjectRemoveVariable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "cds project variable remove <projectKey> <variableName>",
		Long:  ``,
		Run:   removeVarFromProject,
	}
	return cmd
}

func removeVarFromProject(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		sdk.Exit("Wrong usage: %s\n", cmd.Short)
	}
	projectKey := args[0]
	varName := args[1]

	err := sdk.RemoveVariableFromProject(projectKey, varName)
	if err != nil {
		sdk.Exit("Error: cannot remove variable %s from project %s (%s)\n", varName, projectKey, err)
	}
	fmt.Printf("OK\n")
}
