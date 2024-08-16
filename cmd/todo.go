package cmd

import (
	_ "encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

func test(cmd *cobra.Command, args []string) {
	fmt.Println("Approved")
}

func init() {
	// add command to root
	rootCmd.AddCommand(inputCmd)
	rootCmd.AddCommand(showAllCmd)
	// Flag here and Flag in input have to be named the same
	inputCmd.Flags().StringP("message", "i", "default", "Get a input from the user")
	showAllCmd.Flags().StringP("", "", "", "")
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

var showAllCmd = &cobra.Command{
	Use:   "showAll",
	Short: "Show all",
	Long: "Shows all stored tasks" +
		"For example: TodoApp showAll",

	Run: showAll,
}

func showAll(cmd *cobra.Command, args []string) {
	data, _ := os.Open("data.json")
	bytes, _ := ioutil.ReadAll(data)
	fmt.Println(string(bytes))
	defer data.Close()
}
