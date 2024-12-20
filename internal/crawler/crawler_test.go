package crawler

import (
	"fmt"
	"os"
	"testing"
)

func TestCrawl(t *testing.T) {
	// Test case 1: Valid baseURL
	t.Run("Valid baseURL", func(t *testing.T) {
		baseURL := "https://example.com"
		id := "test1"
		maxDepth := 3

		err := Crawl(baseURL, id, maxDepth)

		if err != nil {
			t.Errorf("Crawl function returned an error for valid input: %v", err)
		}

		// Check if the file was created
		filePath := fmt.Sprintf("%s.txt", id)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("Expected file '%s' to be created, but it wasn't", filePath)
		} else {
			// Cleanup: Remove the created file
			err := os.Remove(filePath)
			if err != nil {
				t.Errorf("Failed to remove test file: %v", err)
			}
		}
	})

	// Test case 2: Invalid baseURL
	t.Run("Invalid baseURL", func(t *testing.T) {
		baseURL := "invalid-url"
		id := "test2"
		maxDepth := 3

		err := Crawl(baseURL, id, maxDepth)

		if err == nil {
			t.Error("Expected Crawl function to return an error for an invalid baseURL, but it didn't")
		}
	})

}
