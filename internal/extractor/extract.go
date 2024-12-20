package extractor

import (
	"golang.org/x/net/html"
	"io"
)

// ExtractLinks parses the HTML body from an io.Reader and returns all found links as a slice of strings.
func ExtractLinks(body io.Reader) ([]string, error) {
	var links []string
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err // Properly return the error to the caller
	}

	// Recursive function to traverse the HTML nodes
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val) // Collect the link
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c) // Recurse on child nodes
		}
	}

	f(doc) // Start the traversal with the document node
	return links, nil
}
