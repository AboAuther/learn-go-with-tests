package concurrency

import (
	"sync"
	"time"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	//var c = struct {
	//	sync.RWMutex
	//	result map[string]bool
	//}{result: make(map[string]bool)}
	result := make(map[string]bool)
	Mutex := sync.RWMutex{}
	for _, url := range urls {
		go func(u string) {
			Mutex.Lock()
			result[u] = wc(u)
			Mutex.Unlock()
		}(url)
	}
	time.Sleep(2 * time.Second)
	return result
}
