package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %s\n", url, err)
		return
	}
	defer response.Body.Close()

	fmt.Printf("Fetched %s, Status: %s\n", url, response.Status)
}

func main() {
	urls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.openai.com",
		"https://www.github.com",
		"https://www.wikipedia.org",
	}

	var wg sync.WaitGroup

	startTime := time.Now()

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}

	wg.Wait()

	elapsedTime := time.Since(startTime)
	fmt.Printf("Time taken: %s\n", elapsedTime)
}
