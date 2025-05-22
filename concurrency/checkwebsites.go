package main

import (
	"encoding/json"
	"os"
	"time"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	ch := make(chan bool)
	result := make(map[string]bool)

	for _, url := range urls {
		go func() {
			ch <- wc(url)
		}()
	}

	for _, url := range urls {
		result[url] = <-ch
	}

	return result
}

func slowStubWebsiteChecker_(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func main() {
	result := CheckWebsites(slowStubWebsiteChecker_, []string{
		"a",
		"b",
		"c",
		"d",
		"e",
	})
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	encoder.Encode(result)
}
