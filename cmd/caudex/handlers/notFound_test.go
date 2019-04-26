package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/voltable/graph/cmd/caudex/handlers"
)

func TestNotFoundHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := handlers.NotFound(GetTestHandler())

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusNotFound)
	}
}

// GetTestHandler returns a http.HandlerFunc for testing http middleware
func GetTestHandler() http.HandlerFunc {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		panic("test entered test handler, this should not happen")
	}
	return http.HandlerFunc(fn)
}
