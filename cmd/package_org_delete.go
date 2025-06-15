/*
Copyright © 2025 Pentti Laitinen (pentti.laitinen@gmail.com)
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageOrgDeleteCmd = &cobra.Command{
	Use:   "package",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("package called")
	},
}

func init() {
	packageOrgCmd.AddCommand(packageOrgDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// packageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// packageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	packageOrgDeleteCmd.Flags().BoolP("untagged", "u", false, "remove untagged version of package")
	packageOrgDeleteCmd.Flags().BoolP("tagged", "t", false, "remove tagged versions of package")
	packageOrgDeleteCmd.Flags().Int("keep", 10, "how many versions to keep")
	packageOrgDeleteCmd.Flags().String("orgName", "", "organization name who owns the package")
	packageOrgDeleteCmd.Flags().String("packageName", "", "package name")
	packageOrgDeleteCmd.MarkFlagRequired("orgName")
	packageOrgDeleteCmd.MarkFlagRequired("packageName")
	packageOrgDeleteCmd.MarkFlagsOneRequired("tagged", "untagged")
}
