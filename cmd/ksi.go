package cmd

import (
	"fmt"
	"os"

	"github.com/hashlab/dukptkit/lib"
	"github.com/spf13/cobra"
)

var ksiLength int

func init() {
	rootCmd.AddCommand(generateKSICmd)
	generateKSICmd.Flags().IntVarP(&ksiLength, "length", "l", 5, "Change the length of the key")
}

var generateKSICmd = &cobra.Command{
	Use:   "generate-ksi",
	Short: "Generates a BDK identifier (KSI)",
	Run: func(cmd *cobra.Command, args []string) {
		ksi, err := lib.GenerateKSI(ksiLength)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		formatedString, err := lib.FormatString(ksi)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("BDK identifier (KSI): %s\n", formatedString)
	},
}
