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

var crawledUrls = make(map[string]struct{})
var mutex = new(sync.Mutex)

func AddIfNotExist(url string) bool {
	mutex.Lock()
	_, exists := crawledUrls[url]
	if !exists {
		crawledUrls[url] = struct{}{}
	}
	mutex.Unlock()
	return !exists
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	if depth > 0 && AddIfNotExist(url) {
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("found: %s %q\n", url, body)
			for _, u := range urls {
				wg.Add(1)
				go Crawl(u, depth-1, fetcher, wg)
			}
		}
	}
	wg.Done()
	return
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	Crawl("http://golang.org/", 4, fetcher, wg)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
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
