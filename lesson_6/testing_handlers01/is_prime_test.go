package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	constants "testing_handlers01/constants"
)

func Test_IsPrimeHandler(t *testing.T) {

	handlers := setUpMux()

	type args struct {
		req *http.Request
	}

	tests := []struct {
		name     string
		args     func(t *testing.T) args
		wantCode int
		wantBody string
	}{
		{
			name: "must return http.StatusBadRequest for an invalid number",
			args: func(t *testing.T) args {
				req, err := http.NewRequest(
					constants.HTTPMethods.GET,
					constants.Endpoints.IsPrime,
					nil,
				)

				checkError(err)

				queryParameterMap := req.URL.Query()

				queryParameterMap.Add("number", "not_number")

				req.URL.RawQuery = queryParameterMap.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusBadRequest,

			// suffix new line as the response body has a new line
			wantBody: "invalid number\n",
		},

		// TODO: CONTINUE implementing tests

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

// Testing HTTP Handlers

//   - build a request and compare response
