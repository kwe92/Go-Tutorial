package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type ResourceStatusCode int

// Enum of resource status code
const (
	DatabaseErrorCode ResourceStatusCode = iota + 300
	NetworkErrorCode
)

const (
	DatabaseResource = "Database"
	NetworkResource  = "Network"
)

type ResourceErr struct {
	Resource string
	ResourceStatusCode
	message string
}

func (r ResourceErr) Error() string {
	return fmt.Sprintf("StatusCode: %d | Resource %s | Error Message: %s",
		r.ResourceStatusCode,
		r.Resource,
		r.message)
}

func (r ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreCode := other.ResourceStatusCode == 0
		ignoreResource := other.Resource == ""
		matchCode := r.ResourceStatusCode == other.ResourceStatusCode
		matchResource := r.Resource == other.Resource
		return matchResource && ignoreCode ||
			matchCode && ignoreResource ||
			matchCode &&
				matchResource
	}
	return false
}

// generateErrors: takes as an argument a slice of errors
// returns a closure that generates a random error within the slice range.
func generateErrors(errs []error) func() error {
	numErrors := len(errs)

	return func() error {
		idx := rand.Intn(numErrors)
		return errs[idx]
	}
}

func main() {
	databaseError00 := ResourceErr{
		Resource:           DatabaseResource,
		ResourceStatusCode: DatabaseErrorCode,
		message:            "there was an issue connecting with the database.",
	}

	databaseError01 := ResourceErr{
		Resource:           DatabaseResource,
		ResourceStatusCode: DatabaseErrorCode,
		message:            "failed to insert value.",
	}

	networkError := ResourceErr{
		Resource:           NetworkResource,
		ResourceStatusCode: NetworkErrorCode,
		message:            "failed to connect to the network.",
	}

	errs := []error{databaseError00, databaseError01, networkError}

	errorGenerator := generateErrors(errs)

	err := errorGenerator()

	if databaseError := errors.Is(err, ResourceErr{Resource: DatabaseResource}); databaseError {
		fmt.Println("database error, matching Resource")
		fmt.Println(err)

	}

	err = errorGenerator()

	if networkError := errors.Is(err, ResourceErr{Resource: NetworkResource}); networkError {
		fmt.Println("network error, matching Resource")
		fmt.Println(err)

	}

	err = errorGenerator()

	if databaseError := errors.Is(err, ResourceErr{ResourceStatusCode: DatabaseErrorCode}); databaseError {
		fmt.Println("database error, matching ResourceStatusCode")
		fmt.Println(err)

	}

	err = errorGenerator()

	if networkError := errors.Is(err, ResourceErr{ResourceStatusCode: NetworkErrorCode}); networkError {
		fmt.Println("network error, matching ResourceStatusCode")
		fmt.Println(err)

	}

}

// Overriding errors.Is | Comparing Error Instances and Values

//   - defined errors can implement their own errors.Is method `overriding the default behavior of errors.Is`
//   - overriding errors.Is allows the comparison of different error instances and pattern matching of errors
//   - you specify the equality of the defined error and the target it will be compared against
//   - similar to overriding the equality operator of an object to compare object instances
//   - you do not have to use all of the fields of your custom error in the equality comparison
//   - a user defined error that implements errors.Is can be passed as the first argument to errors.Is
//   - The second argument is an instance to compare against

// - in the above example we catch all database or network errors regardless of what the error message is
// - we only compare the name of the resource and the resource code to determine the type of error being thrown

// When to Use errors.Is

//   - when you want to look for a specific instance or value an error may contain for comparision
