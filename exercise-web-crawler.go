package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	var c2 func(string, int, chan bool) //try commenting this out later
	var printChannel = make(chan bool, 1)
	var mapChannel = make(chan map[string]bool, 1)

	printChannel <- true
	mapChannel <- map[string]bool{url: true}

	c2 = func(url string, depth int, topDone chan bool) {

		if depth <= 0 {
			topDone <- true
			return
		}
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			<-printChannel
			fmt.Println(err)
			printChannel <- true
			topDone <- true
			return
		}
		<-printChannel
		fmt.Printf("found: %s %q\n", url, body)
		printChannel <- true

		subChannelCount := 0
		subChannelDone := make(chan bool)
		m := <-mapChannel
		for _, u := range urls {
			if !m[u] {
				m[u] = true
				subChannelCount++
				go c2(u, depth-1, subChannelDone)
			}
		}
		mapChannel <- m
		for ; subChannelCount > 0; subChannelCount-- {
			<-subChannelDone
		}
		topDone <- true
		return
	}
	done := make(chan bool)
	go c2(url, depth, done)
	<-done
}

func main() {
	Crawl("http://golang.org/", 3, fetcher)
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
