package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

var defaultTimeout = 10 * time.Second

func Racer(url1, url2 string) (string, error) {
	return ConfigurableRacer(url1, url2, defaultTimeout)
}

// ConfigurableRacer takes two urls and timeout, return the first one to return before timeout,
// otherwise returns error
func ConfigurableRacer(url1, url2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func run(ctx context.Context) {
	fmt.Println("running")
	// time.Sleep(2 * time.Second)
	fmt.Println("done running")
}

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i, ch := range chans {
		go func() {
			for {
				time.Sleep(time.Duration(i+1) * time.Second)
				ch <- i
			}
		}()
	}

	// for {
	// 	select {
	// 	case chan0 := <-chans[0]:
	// 		fmt.Println("received from first channel", chan0)
	// 	case chan1 := <-chans[1]:
	// 		fmt.Println("received from second channel", chan1)
	// 	}
	// 	case
	// }

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()

	}()
	<-ctx.Done()

	fmt.Println("after done")
	err := ctx.Err()
	fmt.Println("ctx err:", err)

	fmt.Println("main done")
}
