/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/andersonphiri/concurrency-patterns-golang/confinement"
	"github.com/spf13/cobra"
)

// confinementCmd represents the confinement command
var confinementCmd = &cobra.Command{
	Use:   "confinement",
	Short: "runs confinement examples",
	Long: `runs confinement examples.
	To Run:
	./goconpa run channels confinement
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("confinement called")
		confinement.RunBasicLexicalConfinement()
		log.Printf("\noperating on a non thread safe type\n")
		confinement.OperateOnNonThreadSafeTypeConcurrently()
	},
}

func init() {
	channelsCmd.AddCommand(confinementCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// confinementCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// confinementCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
