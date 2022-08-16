package cmd

import (
	"fmt"
	"os"

	"github.com/omarahm3/gogo/taskino/db"
	"github.com/spf13/cobra"
)

func list(cmd *cobra.Command, args []string) {
  tasks, err := db.GetTasks()

  if err != nil {
    fmt.Printf("error::: [%s]", err.Error())
    os.Exit(1)
  }

  for _, t := range tasks {
    fmt.Printf("[%d] %s\n", t.ID, t.Content)
  }
}
