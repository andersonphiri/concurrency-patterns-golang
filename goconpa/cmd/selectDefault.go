/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (
	"github.com/andersonphiri/concurrency-patterns-golang/channels"
	"github.com/spf13/cobra"
)

// selectDefaultCmd represents the selectDefault command
var selectDefaultCmd = &cobra.Command{
	Use:   "selectDefault",
	Short: "runs an example of doing something while waiting for any channel to be ready, inside a select block",
	Long: `runs an example of doing something while waiting 
	for any channel to be ready, inside a select block.
	To Run:
	goconpa run channels selectDefault
	`,
	Run: func(cmd *cobra.Command, args []string) {
		channels.RunSelect()
	},
}

func init() {
	channelsCmd.AddCommand(selectDefaultCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selectDefaultCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selectDefaultCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
