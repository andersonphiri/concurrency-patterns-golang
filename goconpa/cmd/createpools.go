/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (
	"github.com/andersonphiri/concurrency-patterns-golang/pools"
	"github.com/spf13/cobra"
)

// createpoolsCmd represents the createpools command
var createpoolsCmd = &cobra.Command{
	Use:   "createpools",
	Short: "demonstrates an example of using pools",
	Long: `demonstrates an example of using pools. see how many bufferred items are created`,
	Run: func(cmd *cobra.Command, args []string) {
		pools.CreatePools{}.RunCreatingPools()
	},
}

func init() {
	runCmd.AddCommand(createpoolsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createpoolsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createpoolsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
