package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "get input",
	Long: "get input" +
		"For example: TodoApp input -i 5",

	Run: input,
}

func init() {
	rootCmd.AddCommand(inputCmd)
	// Flag here and Flag in input have to be named the same
	inputCmd.Flags().StringP("message", "i", "default", "Get a input from the user")
}

func input(cmd *cobra.Command, args []string) {
	message, _ := cmd.Flags().GetString("message")
	fmt.Println(message)
}
