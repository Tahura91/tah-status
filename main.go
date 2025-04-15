package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// checkURL sends a GET request to the specified URL and returns the status code and response time.
func checkURL(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("❌ %s - Error: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	statusText := http.StatusText(statusCode)

	switch {
	case statusCode >= 200 && statusCode < 300:
		fmt.Printf("✅ %s - %d %s (%.2fs)\n", url, statusCode, statusText, duration.Seconds())
	case statusCode >= 300 && statusCode < 400:
		fmt.Printf("🔄 %s - %d %s (%.2fs)\n", url, statusCode, statusText, duration.Seconds())
	default:
		fmt.Printf("❌ %s - %d %s (%.2fs)\n", url, statusCode, statusText, duration.Seconds())
	}
}

// readURLsFromFile reads URLs from a given file and returns them as a slice of strings.
func readURLsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}

	return urls, scanner.Err()
}

func main() {
	urlFlag := flag.String("url", "", "Specify a single URL to check.")
	fileFlag := flag.String("file", "", "Specify a file containing a list of URLs to check.")
	helpFlag := flag.Bool("help", false, "Display usage instructions.")
	flag.Parse()

	if *helpFlag || (*urlFlag == "" && *fileFlag == "") {
		fmt.Println("Usage:")
		fmt.Println("  --url <URL>    Check a single URL.")
		fmt.Println("  --file <FILE>  Check multiple URLs listed in a file.")
		fmt.Println("  --help         Display this help message.")
		return
	}

	if *urlFlag != "" {
		checkURL(*urlFlag)
	}

	if *fileFlag != "" {
		urls, err := readURLsFromFile(*fileFlag)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		for _, url := range urls {
			checkURL(url)
		}
	}
}
