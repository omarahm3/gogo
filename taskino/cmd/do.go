package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func do(cmd *cobra.Command, args []string) {
	fmt.Println("dummy do command")
}
