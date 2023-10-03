package main

import (
	"fmt"
	"net/http"
	"testing"
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
				return args{}
			},
			wantCode: 200,
			wantBody: "",
		},
	}

	fmt.Printf("format", handlers, tests)

}
