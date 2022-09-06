package cmd

import (
	"fmt"
	"os"

	"github.com/omarahm3/gogo/secret/secret"
	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:   "del",
	Short: "Delete a secret from vaul",
	Run:   del,
}

func del(cmd *cobra.Command, args []string) {
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

	err := v.Delete(key)
	check(err, "unknown error occurred, probably value doesn't exist")

	fmt.Printf("Key [%s] was deleted\n", key)
}

func init() {
	rootCmd.AddCommand(deleteCommand)
}
