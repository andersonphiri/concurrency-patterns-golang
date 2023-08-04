/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// channelsCmd represents the channels command
var channelsCmd = &cobra.Command{
	Use:   "channels",
	Short: "runs examples about channels",
	Long: `runs examples about channels. must be run with available subcommands`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("use available subcommands")
	},
}

func init() {
	runCmd.AddCommand(channelsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// channelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// channelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
