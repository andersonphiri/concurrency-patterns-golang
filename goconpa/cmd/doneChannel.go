/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (

	"github.com/andersonphiri/concurrency-patterns-golang/forselect"
	"github.com/spf13/cobra"
)

// doneChannelCmd represents the doneChannel command
var doneChannelCmd = &cobra.Command{
	Use:   "doneChannel",
	Short: "A brief description of your command",
	Long: `To run:
	goconpa run forselect doneChannel
	`,
	Run: func(cmd *cobra.Command, args []string) {
		forselect.RunDoneChannel()
		forselect.RunTellProducerToStop()
	},
}

func init() {
	forselectCmd.AddCommand(doneChannelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneChannelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneChannelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
