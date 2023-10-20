package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Emplyee struct {
	id     int
	fname  string
	lname  string
	salary float32
}

func spacedPrint(str string) {
	fmt.Println("\n" + str + "\n")
}

type Logger struct {
	Handler http.Handler
	Prefix  string
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if len(l.Prefix) == 0 {
		log.SetPrefix("Logged Event:")
	} else {
		log.SetPrefix(l.Prefix)
	}

	start := time.Now()

	l.Handler.ServeHTTP(w, r)

	log.Printf("%s %v %v", r.Method, r.URL, time.Since(start))

}

func NewLogger(handler http.Handler) *Logger {

	var prefix string

	fmt.Println("prefix first element:", len(prefix))
	return &Logger{}
}

func main() {

	arr0 := []int{1, 2, 3, 4, 5}

	fmt.Println(arr0[len(arr0)-1])

	m := map[string]string{
		"LOG_LEVEL": "DEBUG",
		"API_KEY":   "12345678-1234-1234-1234-1234-123456789abc",
	}
	println(createKeyValuePairs(m))

	// arr0 := []int{2, 3, 5, 7}

	// arr1 := arr0[2:4]

	// fmt.Printf("%v %v %p %p", arr0, arr1, &arr0, &arr1)

	// arr1[0] = 9999999

	// fmt.Printf("%v %v %p %p", arr0, arr1, &arr0, &arr1)

	// Logger := NewLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	// fmt.Println(Logger)

}

func createKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

// fmt
//   - a package part of the standard library
//   - used to format and print strings

// intArr := [3]int{1, 2, 3}

// const variable0 = 4

// fmt.Printf("\n%v\n", variable0)

// emp1 := Emplyee{id: 1101, fname: "Kweayon", lname: "Clark", salary: 130000}

// spacedPrint("Hello Universe.")
// fmt.Printf("%v\n\n", intArr)
// fmt.Printf("%v\n\n", emp1)

// ?----------Working With Strings----------?

// s0 := "welcome"

// s1 := s0[0:4]

// fmt.Println(s0, s1, &s0, &s1)

// s1 = "to the void"

// fmt.Println(s0, s1, &s0, &s1)
