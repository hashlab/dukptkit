package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashlab/hsmkit/lib"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

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

		fmt.Println("")
		fmt.Printf("Combined Key: %s\n", cKey)
		fmt.Printf("Final KCV: %s\n", cKcv)
	},
}
