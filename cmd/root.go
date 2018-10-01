package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dukptkit",
	Short: "DUKPTKit is a tool for generating keys",
	Long: `A golang tool to generate keys required for using 3DES-DUKPT.
Complete documentation is available at https://github.com/hashlab/dukptkit`,
}

// Execute is a function for running the command line tool
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
