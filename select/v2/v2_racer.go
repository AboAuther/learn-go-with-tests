package racer

import (
	"fmt"
	"net/http"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("error occurred while fetching page, error: %s", err.Error())
		}
		defer resp.Body.Close()
		close(ch)
	}()
	return ch
}
