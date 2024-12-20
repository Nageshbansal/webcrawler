package connector

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchURL(t *testing.T) {
	// Mock server for testing
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Mock response"))
	}))
	defer mockServer.Close()

	wc := NewWebsiteConnector()
	domain := mockServer.URL

	t.Run("Successful fetch", func(t *testing.T) {
		resp, err := wc.FetchURL(domain, mockServer.URL)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK, got %d", resp.StatusCode)
		}
	})

	t.Run("Invalid domain", func(t *testing.T) {
		_, err := wc.FetchURL(domain, "http://invalid-domain.com")
		if err == nil {
			t.Errorf("Expected error for invalid domain, got nil")
		}
	})

	t.Run("Parsing error", func(t *testing.T) {
		_, err := wc.FetchURL(domain, ":")
		if err == nil {
			t.Errorf("Expected error for parsing error, got nil")
		}
	})

	t.Run("Unexpected status code", func(t *testing.T) {
		// Mock server with a non-OK status
		mockServerErr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer mockServerErr.Close()

		_, err := wc.FetchURL(domain, mockServerErr.URL)
		if err == nil {
			t.Errorf("Expected error for unexpected status code, got nil")
		}
	})
}
