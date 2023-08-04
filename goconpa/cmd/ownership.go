/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (
	"github.com/andersonphiri/concurrency-patterns-golang/channels"
	"github.com/spf13/cobra"
)

// ownershipCmd represents the ownership command
var ownershipCmd = &cobra.Command{
	Use:   "ownership",
	Short: "runs ownership example as a subcommand for channels command",
	Long: `runs ownership example as a subcommand for channels command.
	For example:
	./goconpa run channels ownership
	`,
	Run: func(cmd *cobra.Command, args []string) {
		channels.RunExampleClearlyDefiningChannelOwnership()
	},
}

func init() {
	channelsCmd.AddCommand(ownershipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ownershipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ownershipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
