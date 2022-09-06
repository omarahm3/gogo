package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var encodingKey string

var rootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Save secrets encrypted",
}

func Init() {
	rootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "encryption key")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func vaultPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".vault")
}

func check(err error, message string) {
	if err != nil {
		fmt.Println(message)
		os.Exit(1)
	}
}
