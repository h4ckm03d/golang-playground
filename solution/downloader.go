package solution

import (
	"fmt"
	"sync"
	"time"
)

// FIX: this implementation of Downloader contains a data race and
// can access the same URL more than once.

type Downloader struct {
	cache map[string][]byte
	// add necessary fields to protect cache
	mu sync.Mutex
}

func NewDownloader() *Downloader {
	return &Downloader{
		cache: make(map[string][]byte),
	}
}

// Fetch may be called concurrently by multiple goroutines.
// You have to fix data races and improve the logic to access the same URL only once.
func (d *Downloader) Fetch(u string) []byte {

	fmt.Println(u)
	if content, ok := d.cache[u]; ok {
		return content
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	fmt.Println("starting", u)
	content := FakeHTTPGet(u)
	d.cache[u] = content
	return content
}

// Do not touch the code below.

func FakeHTTPGet(u string) []byte {
	time.Sleep(10 * time.Millisecond)
	return []byte(u)
}

var downloader = NewDownloader()

func Solution(A []string) bool {
	wg := new(sync.WaitGroup)
	results := make(map[string][]byte)
	var mu sync.Mutex

	for _, u := range A {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()

			content := downloader.Fetch(u)
			mu.Lock()
			results[u] = content
			mu.Unlock()
		}(u)
	}
	wg.Wait()
	return len(A) > 2
}
