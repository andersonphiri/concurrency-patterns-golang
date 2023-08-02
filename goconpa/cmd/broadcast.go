/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (

	"github.com/andersonphiri/concurrency-patterns-golang/waitingforsignal"
	"github.com/spf13/cobra"
)

// broadcastCmd represents the broadcast command
var broadcastCmd = &cobra.Command{
	Use:   "broadcast",
	Short: "runs the cond broadcast example",
	Long: `runs the cond broadcast example`,
	Run: func(cmd *cobra.Command, args []string) {
		waitingforsignal.BroadcastExample{}.RunBroadcast()
	},
}

func init() {
	runCmd.AddCommand(broadcastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// broadcastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// broadcastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
