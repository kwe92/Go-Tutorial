package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	constants "testing_handlers01/constants"
)

func Test_IsPrimeHandler(t *testing.T) {

	// setup HTTP request multiplexer
	handlers := setUpMux()

	// args: holds the created *http.Request as test arguments
	type args struct {
		req *http.Request
	}

	// Table Driven Test collection
	tests := []struct {
		name     string                  // represents the name of the test
		args     func(t *testing.T) args // creates and returns the request
		wantCode int                     // the HTTP status code you expect
		wantBody string                  // the expected content of the response given the request
	}{
		{
			name: "must return http.StatusBadRequest for an invalid number",

			args: func(t *testing.T) args {

				// create a request with the given HTTP method and URL
				req, err := http.NewRequest(
					constants.HTTPMethods.GET,
					constants.Endpoints.IsPrime,
					nil,
				)

				// check request errors
				checkError(err)

				// retrieve URL query parameter hash map
				queryParameterMap := req.URL.Query()

				// add query parameters as a key value pair to the query parameter map
				queryParameterMap.Add("number", "not_number")

				// append query parameters to the request URL
				req.URL.RawQuery = queryParameterMap.Encode()

				// return the created request in the args struct
				return args{
					req: req,
				}
			},

			wantCode: http.StatusBadRequest,

			wantBody: "invalid number\n",
		},

		{
			name: "must return http.StatusOK and true to prime number (7)",

			args: func(*testing.T) args {

				req, err := http.NewRequest(
					constants.HTTPMethods.GET,
					constants.Endpoints.IsPrime,
					nil,
				)

				checkError(err)

				queryParameters := req.URL.Query()

				queryParameters.Add("number", "7")

				req.URL.RawQuery = queryParameters.Encode()

				return args{
					req: req,
				}
			},

			wantCode: http.StatusOK,

			wantBody: "true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tArgs := tt.args(t)

			response := httptest.NewRecorder()

			handlers.ServeHTTP(response, tArgs.req)

			if response.Result().StatusCode != tt.wantCode {
				t.Fatalf("the status code should be [%d] but received [%d]",
					tt.wantCode,
					response.Result().StatusCode,
				)
			}

			if response.Body.String() != tt.wantBody {
				t.Fatalf("the response body should be [%s] but received [%s]",
					tt.wantBody,
					response.Body.String(),
				)
			}
		})
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatalf("failed to create request: %s", err.Error())
	}
}

// Creating Unit Tests for HTTP Handlers

//   - at its core testing HTTP handlers is simple
//     you build a request to mimic a Client request and compare the response recieved

//   - to build the request you can use the http packages NewRequest function

//   - to build an http.ResponseWriter for the handlers ServeHTTP method
//     you can use the httptest.NewRecorder function

// http.NewRequest

//   - returns a pointer to an http.Request structure
//   - takes three arguments:
//       HTTP Method, URL Path, Request Body `which is typically set to nil`
