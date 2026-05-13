package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	// 1. Array containing 3 different URLs
	urls := []string{
		"https://httpbin.org",
		"https://example.com",
		"https://golang.org",
	}

	var wg sync.WaitGroup

	fmt.Println("Starting to fetch pages concurrently...\n")

	for _, url := range urls {

		wg.Add(1)

		go func(siteURL string) {

			defer wg.Done()

			fetchAndDisplay(siteURL)
		}(url)
	}

	wg.Wait()
	fmt.Println("\nAll pages fetched successfully.")
}

func fetchAndDisplay(url string) {
	// Send  HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[-] Error fetching %s: %v\n", url, err)
		return
	}
	// Close the response body to prevent resource leaks
	defer resp.Body.Close()

	// Read the entire response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[-] Error reading body from %s: %v\n", url, err)
		return
	}

	// Print organized results to the console
	fmt.Printf("\nk")
	fmt.Printf("[+] SITE: %s (Status: %s)\n", url, resp.Status)
	fmt.Printf("\n")

	// Convert byte slice to string
	bodyString := string(bodyBytes)

	// Display only the first 300 characters to prevent console overflow
	if len(bodyString) > 300 {
		fmt.Println(bodyString[:300] + "\n... [Truncated / Remaining content hidden] ...\n")
	} else {
		fmt.Println(bodyString + "\n")
	}
}
