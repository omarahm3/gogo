package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func add(cmd *cobra.Command, args []string) {
	fmt.Println("dummy add command")
}
