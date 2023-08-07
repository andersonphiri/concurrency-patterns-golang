/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fanoutfaninCmd represents the fanoutfanin command
var fanoutfaninCmd = &cobra.Command{
	Use:   "fanoutfanin",
	Short: "runs fanoutfanin exaples",
	Long: `To run:
	./goconpa run fanoutfanin concurrentcompute
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("fanoutfanin called. Please use subcommands")
	},
}

func init() {
	runCmd.AddCommand(fanoutfaninCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fanoutfaninCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fanoutfaninCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
