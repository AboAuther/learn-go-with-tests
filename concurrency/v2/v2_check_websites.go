package concurrency

import (
	"sync"
	"time"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {

	result := make(map[string]bool)
	mutex := sync.RWMutex{}
	for _, url := range urls {
		go func(u string) {
			mutex.Lock()
			result[u] = wc(u)
			mutex.Unlock()
		}(url)
	}
	time.Sleep(2 * time.Second)
	return result
}
