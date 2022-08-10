package sitemap

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/omarahm3/gogo/link/link"
	"golang.org/x/net/html"
)

var (
	baseLink       string
	traversedLinks map[string]bool
)

func getContent(link string) string {
	response, err := http.Get(link)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}

func parseLink(baseLink string) map[string]bool {
  result := make(map[string]bool)

	c := getContent(baseLink)

	n, err := html.Parse(strings.NewReader(c))

	if err != nil {
		panic(err)
	}

	links := link.Parse(n)

	for _, link := range links {
		if _, ok := result[link.Href]; ok {
			continue
		}

		result[link.Href] = true
	}

  return result
}

func traverse(baseLink, link string) map[string]bool {
  outputLinks := parseLink(link)

	for outputLink := range outputLinks {
		log.Printf("traverse:: Parsed link: [%s]\n", outputLink)

		validLink, err := GenerateValidUrl(baseLink, outputLink)

		if err != nil {
			continue
		}

		if _, ok := traversedLinks[*validLink]; ok {
			continue
		}

		traversedLinks[*validLink] = true
	}

	return traversedLinks
}

func merge(m1, m2 map[string]bool) map[string]bool {
	for k, v := range m1 {
		m2[k] = v
	}

	return m2
}

func Generate(firstLink string) string {
	baseLink = firstLink
	traversedLinks = map[string]bool{}

  // Initial link
  outputLinks := parseLink(firstLink)

  // Loop over initial link links'
	for outputLink := range outputLinks {
		log.Printf("Generate::  Output link: [%s]\n", outputLink)

		validLink, err := GenerateValidUrl(firstLink, outputLink)

		if err != nil {
      log.Printf("Generate:: ﲅ Ignoring link because [%s]\n", err.Error())
			continue
		}

		if _, ok := traversedLinks[*validLink]; ok {
      log.Println("Generate:: ﲅ Ignoring link because it is duplicated")
			continue
		}

		traversedLinks[*validLink] = true
		log.Printf("Generate::  Valid link: [%s]\n", *validLink)
		traversedLinks = merge(traverse(firstLink, *validLink), traversedLinks)
	}

	jsonString, _ := json.Marshal(traversedLinks)

	return string(jsonString)
}
