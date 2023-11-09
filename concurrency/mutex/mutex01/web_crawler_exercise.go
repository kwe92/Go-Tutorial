package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type UrlCache struct {
	mu    sync.Mutex
	cache map[string]bool
}

var urlCache = UrlCache{
	cache: make(map[string]bool),
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, exit chan bool) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.

	if depth <= 0 {
		// write to exit channel and return
		exit <- true

		return
	}

	// locks access to single thread / process
	urlCache.mu.Lock()

	// check cache for url with comma ok idiom
	_, ok := urlCache.cache[url]

	if ok == false {
		// add url to cache
		urlCache.cache[url] = true

		// unlock access to for other threads and processes
		urlCache.mu.Unlock()

	} else {
		// if url exists
		// write to exit channel, unlock and return

		exit <- true

		urlCache.mu.Unlock()

		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {

		// print error, write to exit channel and return

		fmt.Println("error in Crawl:", err)

		exit <- true

		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	// create new exit channel
	e := make(chan bool)

	for _, url := range urls {
		// invoke recursive go routine
		go Crawl(url, depth-1, fetcher, e)
	}

	// wait for all child goroutines to exit
	for i := 0; i < len(urls); i++ {

		// read from the channel for each goroutine
		<-e
	}

	// write to exit channel
	exit <- true
}

func main() {

	exit := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, exit)

	// read from exit channel
	<-exit
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		body: "The Go Programming Language",
		urls: []string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		body: "Packages",
		urls: []string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		body: "Package fmt",
		urls: []string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		body: "Package os",
		urls: []string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
