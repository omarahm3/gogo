package cmd

import (
	"fmt"
	"os"

	"github.com/omarahm3/gogo/secret/secret"
	"github.com/spf13/cobra"
)

var getCommand = &cobra.Command{
	Use:   "get",
	Short: "Get a secret value from vault",
	Run:   get,
}

func get(cmd *cobra.Command, args []string) {
	if encodingKey == "" {
		fmt.Println("key cannot be empty")
		os.Exit(1)
	}

	key := args[0]
	if key == "" {
		fmt.Println("secret key was not provided")
		os.Exit(1)
	}

	v := secret.File(encodingKey, vaultPath())

	value, err := v.Get(args[0])
	check(err, "unknown error occurred, probably value doesn't exist")

	fmt.Printf("%s=%s\n", key, value)
}

func init() {
	rootCmd.AddCommand(getCommand)
}
