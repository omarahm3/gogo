package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func list(cmd *cobra.Command, args []string) {
	fmt.Println("dummy list command")
}
