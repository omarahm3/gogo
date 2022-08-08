package main

import (
	"flag"
	"fmt"
	"os"

	"mrgeek-link/link"

	"golang.org/x/net/html"
)

var htmlFile = flag.String("htmlFile", "", "HTML file to parse")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadHtmlFile(path string) *html.Node {
	data, err := os.Open(path)

	check(err)

  defer data.Close()

	doc, err := html.Parse(data)

	check(err)

	return doc
}

func printLinks(links []link.Link) {
  for _, l := range links {
    fmt.Printf("Link: %s\nData: %s\n--------\n", l.Href, l.Text)
  }
}

func main() {
  flag.Parse()

  if *htmlFile == "" {
    panic("valid HTML file is needed")
  }

	doc := ReadHtmlFile(*htmlFile)
	links := link.Parse(doc)

  printLinks(links)
}
