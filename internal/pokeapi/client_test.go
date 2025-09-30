package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestFetchData_CacheMissThenHit tests that the client fetches data on cache miss
// and uses the cache on subsequent requests.
func TestFetchData_CacheMissThenHit(t *testing.T) {
	// 1. Setup a fake server
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("testdata"))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// 2. Create client with a long cache interval
	client := NewClient(5 * time.Second)

	// 3. First call → should fetch from server
	body, err := client.FetchData(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(body) != "testdata" {
		t.Errorf("expected %q, got %q", "testdata", string(body))
	}

	// 4. Stop server to ensure cache is used for the second call
	server.Close()

	// 5. Second call → should come from cache
	body2, err := client.FetchData(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(body2) != "testdata" {
		t.Errorf("expected cached %q, got %q", "testdata", string(body2))
	}
}

// TestFetchData_BadStatus ensures non-200 status codes return error.
func TestFetchData_BadStatus(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot) // 418: I'm a teapot
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(5 * time.Second)
	_, err := client.FetchData(server.URL)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}