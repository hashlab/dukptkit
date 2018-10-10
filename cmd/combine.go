package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashlab/dukptkit/lib"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

var showCombinedKey bool

var questions = []*survey.Question{
	{
		Name:   "component-key-1",
		Prompt: &survey.Password{Message: "Enter the component key 1:"},
	},
	{
		Name:   "component-key-2",
		Prompt: &survey.Password{Message: "Enter the component key 2:"},
	},
	{
		Name:   "component-key-3",
		Prompt: &survey.Password{Message: "Enter the component key 3:"},
	},
}

func init() {
	rootCmd.AddCommand(combineCmd)
	combineCmd.Flags().BoolVarP(&showCombinedKey, "show-combined-key", "s", false, "Show combined key")
}

var combineCmd = &cobra.Command{
	Use:   "combine",
	Short: "Generates a combined key along with it's KCV using 3 component keys",
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			ComponentKey1 string `survey:"component-key-1"`
			ComponentKey2 string `survey:"component-key-2"`
			ComponentKey3 string `survey:"component-key-3"`
		}{}

		err := survey.Ask(questions, &answers)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		keyBytes1, err := lib.HexToBytes(strings.ToUpper(answers.ComponentKey1))

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		keyBytes2, err := lib.HexToBytes(strings.ToUpper(answers.ComponentKey2))

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		keyBytes3, err := lib.HexToBytes(strings.ToUpper(answers.ComponentKey3))

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		cKeyBytes, cKey := lib.GenerateCombinedKey(keyBytes1, keyBytes2, keyBytes3)

		_, cKcv, err := lib.CalculateKcv(cKeyBytes)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		_, cK1, err := lib.CalculateKcv(keyBytes1)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		_, cK2, err := lib.CalculateKcv(keyBytes2)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		_, cK3, err := lib.CalculateKcv(keyBytes3)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		formatedString, err := lib.FormatString(cKey)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("")
		fmt.Printf("KCV from Component Key 1: %s\n", cK1)
		fmt.Printf("KCV from Component Key 2: %s\n", cK2)
		fmt.Printf("KCV from Component Key 3: %s\n", cK3)
		fmt.Println("")

		if showCombinedKey {
			fmt.Printf("Combined Key: %s\n", formatedString)
		} else {
			fmt.Printf("Combined Key: %s\n", "Top Secret!")
		}

		fmt.Printf("Final KCV: %s\n", cKcv)
		fmt.Printf("Odd Parity: %t\n", lib.IsOddParityAdjusted(cKeyBytes))
	},
}
