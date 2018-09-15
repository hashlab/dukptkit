package cmd

import (
	"fmt"
	"os"

	"github.com/hashlab/hsmkit/lib"
	"github.com/spf13/cobra"
)

var forceOdd bool

func init() {
	rootCmd.AddCommand(componentKeyCmd)
	componentKeyCmd.Flags().BoolVarP(&forceOdd, "force-odd", "f", false, "Force odd parity")
}

var componentKeyCmd = &cobra.Command{
	Use:   "component-key",
	Short: "Generates a component key along with it's KCV",
	Run: func(cmd *cobra.Command, args []string) {
		keyBytes, key, err := lib.GenerateComponentKey(16, forceOdd)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		_, kcv, err := lib.CalculateKcv(keyBytes)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("Component Key: %s\n", key)
		fmt.Printf("KCV: %s\n", kcv)
	},
}
