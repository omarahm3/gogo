package main

import (
	"fmt"

	"github.com/omarahm3/gogo/secret/secret"
)

func main() {
	v := secret.File("some key", ".vault")
	v.Set("test", "hello")
	value, err := v.Get("test")
	fmt.Println(value, err)
}
