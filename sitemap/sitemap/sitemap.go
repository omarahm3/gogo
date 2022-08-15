package sitemap

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/omarahm3/gogo/link/link"
	"golang.org/x/net/html"
)

type empty struct{}

var (
	baseLink       string
	traversedLinks map[string]bool
)

func links(r io.Reader, base string) []string {
	var result []string

	n, err := html.Parse(r)

	if err != nil {
		panic(err)
	}

	links := link.Parse(n)

	for _, link := range links {
		if _, ok := traversedLinks[link.Href]; ok {
			continue
		}

		switch {
		case strings.HasPrefix(link.Href, "/"):
			result = append(result, fmt.Sprintf("%s%s", base, link.Href))
		case strings.HasPrefix(link.Href, "http"):
			// TODO Check if host is also equal baseUrl host
			result = append(result, link.Href)
		}

		traversedLinks[link.Href] = true
	}

	return result
}

func withPrefix(pfx string) func(string) bool {
	return func(s string) bool {
		return strings.HasPrefix(s, pfx)
	}
}

func filter(links []string, cb func(string) bool) []string {
	var ret []string

	for _, link := range links {
		if cb(link) {
			ret = append(ret, link)
		}
	}

	return ret
}

func get(u string) []string {
	response, err := http.Get(u)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	baseUrl := &url.URL{
		Host:   response.Request.URL.Host,
		Scheme: response.Request.URL.Scheme,
	}

	baseLink = baseUrl.String()

	return filter(links(response.Body, baseLink), withPrefix(baseLink))
}

func traverse(u string, depth int) []string {
	seen := make(map[string]empty)

	var q map[string]empty

	nq := map[string]empty{
		u: empty{},
	}

	for i := 0; i <= depth; i++ {
		q, nq = nq, make(map[string]empty)

    if len(q) == 0 {
      break
    }

		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}

			seen[url] = empty{}

			for _, link := range get(u) {
				nq[link] = empty{}
			}
		}
	}

	ret := make([]string, 0, len(seen))

	for url := range seen {
		ret = append(ret, url)
	}

	return ret
}

func Generate(startLink string, depth int) string {
	traversedLinks = make(map[string]bool)

	pages := traverse(startLink, depth)

	for _, page := range pages {
		fmt.Println(page)
	}

	return ""
}
