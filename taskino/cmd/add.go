package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/omarahm3/gogo/taskino/db"
	"github.com/spf13/cobra"
)

func add(cmd *cobra.Command, args []string) {
	c := strings.Join(args, " ")
	err := db.CreateTask(c)

	if err != nil {
		fmt.Printf("error::: [%s]", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Added '%s' to your list\n", c)
}
