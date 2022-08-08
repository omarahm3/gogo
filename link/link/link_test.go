package link

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func createNode(data string) *html.Node {
  n, err := html.Parse(strings.NewReader(data))

  if err != nil {
    panic(err)
  }

  return n
}

func TestParseSingleLink(t *testing.T) {
  expectedText := "test text"
  expectedLink := "/test"
  node := createNode(fmt.Sprintf(`<a href="%s">%s</a>`, expectedLink, expectedText))

  links := Parse(node)

  if len(links) != 1 {
    t.Error("Links is not equal 1")
  }

  for _, link := range links {
    if link.Href != expectedLink {
      t.Errorf("Expected [%s], got [%s]", expectedLink, link.Href)
    }

    if link.Text != expectedText {
      t.Errorf("Expected [%s], got [%s]", expectedText, link.Text)
    }
  }
}

func TestParseMultipleLinks(t *testing.T) {
  node := createNode(`
  <main>
    <a href="/test1">Test1</a>
    <a href="/test2">Test2</a>
    <a href="/test3">Test3</a>
  </main>
  `)

  links := Parse(node)

  if len(links) != 3 {
    t.Error("Links is not equal 3")
  }

  for i, link := range links {
    expectedLink := fmt.Sprintf("/test%d", i+1)

    if link.Href != expectedLink {
      t.Errorf("Expected [%s], got [%s]", expectedLink, link.Href)
    }

    expectedText := fmt.Sprintf("Test%d", i+1)
    
    if link.Text != expectedText {
      t.Errorf("Expected [%s], got [%s]", expectedText, link.Text)
    }
  }
}

func TestParseLinkWithoutComment(t *testing.T) {
  expectedText := "Test text"
  expectedLink := "/test"
  node := createNode(fmt.Sprintf(`
  <main>
    <a href="%s">%s <!-- commented text SHOULD NOT be included! --></a>
  </main>
  `, expectedLink, expectedText))

  links := Parse(node)

  if len(links) != 1 {
    t.Error("Links is not equal 1")
  }

  for _, link := range links {
    if link.Href != expectedLink {
      t.Errorf("Expected [%s], got [%s]", expectedLink, link.Href)
    }

    if link.Text != expectedText {
      t.Errorf("Expected [%s], got [%s]", expectedText, link.Text)
    }
  }
}

func TestParseAllLinkText(t *testing.T) {
  expectedText := "Test this on here"
  expectedLink := "/test"
  node := createNode(fmt.Sprintf(`
  <main>
    <a href="%s">
      Test this on <strong>here</strong>
    </a>
  </main>
  `, expectedLink))

  links := Parse(node)

  if len(links) != 1 {
    t.Error("Links is not equal 1")
  }

  for _, link := range links {
    if link.Href != expectedLink {
      t.Errorf("Expected [%s], got [%s]", expectedLink, link.Href)
    }

    if link.Text != expectedText {
      t.Errorf("Expected [%s], got [%s]", expectedText, link.Text)
    }
  }
}
