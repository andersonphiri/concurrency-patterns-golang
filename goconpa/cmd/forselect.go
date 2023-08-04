/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// forselectCmd represents the forselect command
var forselectCmd = &cobra.Command{
	Use:   "forselect",
	Short: "forselect parent command",
	Long: `forselect parent command:
	For example:
	goconpa run forselect doneChannel
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("use subcommands available")
	},
}

func init() {
	runCmd.AddCommand(forselectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forselectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forselectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
