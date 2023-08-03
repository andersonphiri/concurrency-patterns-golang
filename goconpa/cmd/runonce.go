/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/andersonphiri/concurrency-patterns-golang/executeonce"
	"github.com/spf13/cobra"
)

// runonceCmd represents the runonce command
var runonceCmd = &cobra.Command{
	Use:   "runonce",
	Short: "runs a runonce example which is using sync.Once",
	Long: `runs a runonce example which is using sync.Once`,
	Run: func(cmd *cobra.Command, args []string) {
		executeonce.ExecuteOnce{}.RunExecuteOnce()
		fmt.Println("testing example two")
		// executeonce.ExecuteOnce{}.RunExecuteOnce_Two()
	},
}

func init() {
	runCmd.AddCommand(runonceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runonceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runonceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
