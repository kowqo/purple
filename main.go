package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getGoogle()
		}()
	}

	wg.Wait()
	fmt.Println(time.Since(t))

}

func getGoogle() {
	resp, _ := http.Get("https://www.google.com")

	fmt.Println(resp.StatusCode)
}
