package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func root(cmd *cobra.Command, args []string) {
	fmt.Println("dummy root command")
}

var rootCmd = &cobra.Command{
	Use:   "taskino",
	Short: "Yet another CLI task manager",
	Run:   root,
}

var addCommand = &cobra.Command{
	Use:   "add [task]",
	Short: "Add new task",
	Run:   add,
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all existing tasks",
	Run:   list,
}

var doCommand = &cobra.Command{
	Use:   "do",
	Short: "Mark task as complete",
	Run:   do,
}

func Init() {
	rootCmd.AddCommand(addCommand)
	rootCmd.AddCommand(listCommand)
	rootCmd.AddCommand(doCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
