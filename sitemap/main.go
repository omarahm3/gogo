package main

import (
	"flag"
	"fmt"

	"github.com/omarahm3/gogo/sitemap/sitemap"
)

var l = flag.String("link", "", "Website link")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()

	if *l == "" {
		panic("valid URL is needed")
	}

	content := sitemap.Generate(*l)

	fmt.Println(content)
}
