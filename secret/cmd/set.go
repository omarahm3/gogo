package cmd

import (
	"fmt"
	"os"

	"github.com/omarahm3/gogo/secret/secret"
	"github.com/spf13/cobra"
)

var setCommand = &cobra.Command{
	Use:   "set",
	Short: "Set a secret value into vault",
	Run:   set,
}

func set(cmd *cobra.Command, args []string) {
	if encodingKey == "" {
		fmt.Println("key cannot be empty")
		os.Exit(1)
	}

	if len(args) < 2 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}

	key, value := args[0], args[1]
	if key == "" {
		fmt.Println("secret key was not provided")
		os.Exit(1)
	}
	if value == "" {
		fmt.Println("secret value was not provided")
		os.Exit(1)
	}

	v := secret.File(encodingKey, vaultPath())
	err := v.Set(key, value)
	check(err, "unknown error occurred")

	fmt.Printf("Value [%s] is set\n", value)
}

func init() {
	rootCmd.AddCommand(setCommand)
}
