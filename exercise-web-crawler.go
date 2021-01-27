// Note: this solution is not completely correct, because we are still crawling 
// twice for webpages that are not found. It would be better to redefine the map 
// to include errors.

package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeMap struct {
	mu sync.Mutex
	v  map[string]fakeResult
}
var webpages = SafeMap {v: make(map[string]fakeResult)}

func (m *SafeMap) Set (s string, res fakeResult) {
	m.mu.Lock()
	m.v[s] = res
	m.mu.Unlock()
}

func (m *SafeMap) Find (s string) bool {
	m.mu.Lock()
	_, exist := m.v[s]
	m.mu.Unlock()
	return exist
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	fmt.Println("Entering Crawl:", url, depth)
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	res := fakeResult {body, urls}
	webpages.Set(url, res)
	
	for _, u := range urls {
		if !webpages.Find(u) {
			fmt.Println("Crawling", u)
			go Crawl(u, depth-1, fetcher)
		}
	}
	return
}


func main() {
	Crawl("https://golang.org/", 4, fetcher)
	
	// sleep some time for all threads to finish
	time.Sleep(30 * time.Second)
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
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

