/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (
	"github.com/andersonphiri/concurrency-patterns-golang/waitingforsignal"
	"github.com/spf13/cobra"
)

// condCmd represents the cond command
var condCmd = &cobra.Command{
	Use:   "cond",
	Short: "runs an example using go cond struct",
	Long: `runs an example using go cond struct`,
	Run: func(cmd *cobra.Command, args []string) {
		waitingforsignal.WaitForSignal{}.RunWaitForSignal()
	},
}

func init() {
	runCmd.AddCommand(condCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// condCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// condCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
