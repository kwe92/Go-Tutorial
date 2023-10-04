package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	// create a new request to pass to the handler.
	req, err := http.NewRequest("GET", "/health", nil)

	checkError(err, t)

	// create an http.ResponseRecorder, which records the response and satisfies http.ResponseWriter
	rr := httptest.NewRecorder()

	// use the http.HandlerFunc to turn our normal callback into an HTTP handler
	handler := http.HandlerFunc(HealthCheckHandler)

	// call ServeHTTP on the handler directly and pass in the ResponseRecorder and Request
	handler.ServeHTTP(rr, req)

	// check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got: % v want: %v",
			status,
			http.StatusOK,
		)
	}

	expectedContent := `{"alive": true}`

	receivedContent := rr.Body.String()

	// check response body
	if receivedContent != expectedContent {
		t.Errorf("handler returned unexpected body: got %v want: %v",
			receivedContent,
			expectedContent,
		)
	}

}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}
