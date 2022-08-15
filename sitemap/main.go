package main

import (
	"flag"
	"fmt"

	"github.com/omarahm3/gogo/sitemap/sitemap"
)

var (
	url   = flag.String("link", "", "Website link")
	depth = flag.Int("depth", 3, "Maximum depth")
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()

	if *url == "" {
		panic("valid URL is needed")
	}

	content := sitemap.Generate(*url, *depth)

	fmt.Println(content)
}
