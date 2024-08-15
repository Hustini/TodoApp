package cmd

import (
	_ "encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	// "io/ioutil"
	_ "os"
)

func init() {
	// add command to root
	rootCmd.AddCommand(inputCmd)
	// Flag here and Flag in input have to be named the same
	inputCmd.Flags().StringP("message", "i", "default", "Get a input from the user")
}

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "get input",
	Long: "get input" +
		"For example: TodoApp input -i Hello World",

	Run: input,
}

func input(cmd *cobra.Command, args []string) {
	message, _ := cmd.Flags().GetString("message")
	fmt.Println(message)
}
