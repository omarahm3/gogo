package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/omarahm3/gogo/taskino/db"
	"github.com/spf13/cobra"
)

func do(cmd *cobra.Command, args []string) {
	for _, k := range args {
    id, _ := strconv.Atoi(k)

		err := db.DeleteTask(id)

		if err != nil {
			fmt.Printf("error::: [%s]", err.Error())
			os.Exit(1)
		}
	}

	fmt.Println("task is marked complete")
}
