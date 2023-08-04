/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (
	"github.com/andersonphiri/concurrency-patterns-golang/forselect"
	"github.com/spf13/cobra"
)

// orChannelCmd represents the orChannel command
var orChannelCmd = &cobra.Command{
	Use:   "orChannel",
	Short: "goconpa run forselect orChannel",
	Long: `To run:
	./goconpa run forselect orChannel`,
	Run: func(cmd *cobra.Command, args []string) {
		forselect.RunOrChannel()
	},
}

func init() {
	forselectCmd.AddCommand(orChannelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// orChannelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// orChannelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
