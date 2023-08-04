/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (
	"github.com/andersonphiri/concurrency-patterns-golang/channels"
	"github.com/spf13/cobra"
)

// selectWaitForResultCmd represents the selectWaitForResult command
var selectWaitForResultCmd = &cobra.Command{
	Use:   "selectWaitForResult",
	Short: "runs an example of a go routine waiting for result from another, meanwhile performing some other work",
	Long: `
	Runs an example of a go routine waiting for result from another, meanwhile performing some other work
	To Run:
	./goconpa run channels selectWaitForResult
	`,
	Run: func(cmd *cobra.Command, args []string) {
		channels.WaitForResultFromAnother(nil)
	},
}

func init() {
	channelsCmd.AddCommand(selectWaitForResultCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selectWaitForResultCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selectWaitForResultCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
