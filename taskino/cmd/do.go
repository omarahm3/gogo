package cmd

import (
	"fmt"
	"os"

	"github.com/omarahm3/gogo/taskino/db"
	"github.com/spf13/cobra"
)

func do(cmd *cobra.Command, args []string) {
  if len(args) != 1 {
    fmt.Println("only 1 argument is needed")
    os.Exit(1)
  }

  k := args[len(args)-1]
  err := db.DeleteTask(k)

  if err != nil {
    fmt.Printf("error::: [%s]", err.Error())
    os.Exit(1)
  }

  fmt.Println("task is marked complete")
}
