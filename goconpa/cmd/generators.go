/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (
	"github.com/andersonphiri/concurrency-patterns-golang/pipelines"
	"github.com/spf13/cobra"
)

// generatorsCmd represents the generators command
var generatorsCmd = &cobra.Command{
	Use:   "generators",
	Short: "runs pipelines examples",
	Long: `To run:
	./goconpa run pipelines generators
	`,
	Run: func(cmd *cobra.Command, args []string) {
		pipelines.RunUseGenerateAndTake()
	},
}

func init() {
	pipelinesCmd.AddCommand(generatorsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generatorsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generatorsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
