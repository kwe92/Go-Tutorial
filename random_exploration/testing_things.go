package main

import (
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

	Logger := NewLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	fmt.Println(Logger)

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
