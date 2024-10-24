package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Assume this is your API handler
func YourAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello, World!"}`))
}

func BenchmarkYourAPI(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		YourAPIHandler(w, req)
	}
}
