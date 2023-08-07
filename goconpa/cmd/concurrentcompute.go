/*
Copyright © 2023 anderson phiri phirianderson23@gmail.com
*/
package cmd

import (
	"fmt"
	"runtime"

	"github.com/andersonphiri/concurrency-patterns-golang/fanoutfanin"
	"github.com/spf13/cobra"
)
var start uint64 
var end uint64 
var parallelFactor uint64
// concurrentcomputeCmd represents the concurrentcompute command
var concurrentcomputeCmd = &cobra.Command{
	Use:   "concurrentcompute",
	Short: "runs concurrentcompute",
	Long: `To run:
	./goconpa run fanoutfanin concurrentcompute
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if start == end || start == 0 {
			return fmt.Errorf("please use reasonable inputs. see default values")
		}
		fmt.Printf("Number of logical CPUs: %v\n", runtime.NumCPU())
		fanoutfanin.RunConcurrentComputesPrimeNumbers(start, end, parallelFactor)
		fmt.Printf("Number of existing MPG goroutines: %v\n", runtime.NumGoroutine())
		return nil
	},
}

func init() {
	fanoutfaninCmd.AddCommand(concurrentcomputeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// concurrentcomputeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// concurrentcomputeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	concurrentcomputeCmd.Flags().Uint64VarP(&start, "start", "s", 1, "--start 1 | -s 1")
	concurrentcomputeCmd.Flags().Uint64VarP(&end, "end", "e", 1000_000, "--end 1 | -e 1")
	concurrentcomputeCmd.Flags().Uint64VarP(&parallelFactor, "parallel-factor", "p", 5, "--parallel-factor 5 | -p 1")
}
