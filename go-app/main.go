package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// === Aplikasi ===
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go app!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Go server running on :8080")
	http.ListenAndServe(":8080", nil)
}

// === Test ===
func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler(rr, req)

	if rr.Body.String() != "Hello from Go app!" {
		t.Errorf("expected 'Hello from Go app!' but got '%s'", rr.Body.String())
	}
}
