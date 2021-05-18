package concurrency

import (
	"sync"
	"time"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	var c = struct {
		sync.RWMutex
		result map[string]bool
	}{result: make(map[string]bool)}
	//results := make(map[string]bool)
	for _, url := range urls {
		go func(u string) {
			c.Lock()
			c.result[u] = wc(u)
			c.Unlock()
		}(url)
	}
	time.Sleep(2 * time.Second)
	return c.result
}
