package main

import (
	"github.com/omarahm3/gogo/taskino/cmd"
	"github.com/omarahm3/gogo/taskino/db"
)

func main() {
  db.Init()
  cmd.Init()
}
