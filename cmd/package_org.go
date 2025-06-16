/*
Copyright © 2025 Pentti Laitinen (pentti.laitinen@gmail.com)
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageOrgCmd = &cobra.Command{
	Use:   "org",
	Short: "org sub comands related to packages",
}

func init() {
	packageCmd.AddCommand(packageOrgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// packageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// packageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
