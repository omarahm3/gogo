package link

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func extractText(n *html.Node) (ret string) {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += extractText(c) + " "
	}

	return strings.Join(strings.Fields(ret), " ")
}

func Parse(n *html.Node) []Link {
	var links []Link

	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			var href string
			for _, a := range n.Attr {
				if a.Key == "href" {
					href = a.Val
				}
			}

			text := extractText(n)

			links = append(links, Link{
				Href: href,
				Text: text,
			})
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)

	return links
}
