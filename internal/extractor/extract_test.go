package extractor

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

// Mock the html.Parse function if needed
var htmlParse = html.Parse

func TestExtractLinksValidHTML(t *testing.T) {
	htmlData := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Test Page</title>
		</head>
		<body>
			<a href="http://example.com">Example</a>
			<a href="http://example.org">Example Org</a>
		</body>
	</html>
	`
	body := strings.NewReader(htmlData)
	expectedLinks := []string{"http://example.com", "http://example.org"}

	links, err := ExtractLinks(body)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(links) != len(expectedLinks) {
		t.Fatalf("Expected %d links, got %d", len(expectedLinks), len(links))
	}

	for i, link := range links {
		if link != expectedLinks[i] {
			t.Errorf("Expected link %s, got %s", expectedLinks[i], link)
		}
	}
}

func TestExtractLinksEmptyHTML(t *testing.T) {
	htmlData := ``
	body := strings.NewReader(htmlData)
	expectedLinks := []string{}

	links, err := ExtractLinks(body)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(links) != len(expectedLinks) {
		t.Fatalf("Expected %d links, got %d", len(expectedLinks), len(links))
	}
}

func TestExtractLinksNoLinks(t *testing.T) {
	htmlData := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Test Page</title>
		</head>
		<body>
			<p>No links here!</p>
		</body>
	</html>
	`
	body := strings.NewReader(htmlData)
	expectedLinks := []string{}

	links, err := ExtractLinks(body)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(links) != len(expectedLinks) {
		t.Fatalf("Expected %d links, got %d", len(expectedLinks), len(links))
	}
}
