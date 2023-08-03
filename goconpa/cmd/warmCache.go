/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (
	"github.com/andersonphiri/concurrency-patterns-golang/pools"
	"github.com/spf13/cobra"
)

// warmCacheCmd represents the warmCache command
var warmCacheCmd = &cobra.Command{
	Use:   "warmCache",
	Short: "runs an example of warm cache",
	Long: `runs an example of warm cache. You may also run this using benchmark testing`,
	Run: func(cmd *cobra.Command, args []string) {
		pools.RunWarmCacheExample("localhost", "8089", 20)
	},
}

func init() {
	runCmd.AddCommand(warmCacheCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// warmCacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// warmCacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
