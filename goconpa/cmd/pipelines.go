/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pipelinesCmd represents the pipelines command
var pipelinesCmd = &cobra.Command{
	Use:   "pipelines",
	Short: "runs pipelines examples",
	Long: `To run:
	./goconpa run pipelines [subcommand]
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("please specify a subcommand")
	},
}

func init() {
	runCmd.AddCommand(pipelinesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pipelinesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pipelinesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
