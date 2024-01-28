package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"crypto/tls"
)

func main() {
	// Define flags
	verbose := flag.Bool("v", false, "Print verbose output")
	insecure := flag.Bool("k", false, "Allow connections to SSL sites without certs")
	failFast := flag.Bool("f", false, "Fail fast with no output on HTTP errors")
	flag.Parse()

	// Check if a URL is provided as a command-line argument
	if flag.NArg() < 1 {
		fmt.Println("Usage: go run main.go [-v] [-k] [-f] <url>")
		return
	}

	// Get the URL from the command-line arguments
	url := flag.Arg(0)

	// Make a GET request
	if *verbose {
		fmt.Println("Making GET request to", url)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: *insecure},
		},
	}

	response, err := client.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		if *failFast {
			os.Exit(1)
		}
		return
	}
	defer response.Body.Close()

	// Check for HTTP errors and fail fast if needed
	if response.StatusCode >= 400 {
		fmt.Printf("HTTP error: %s\n", response.Status)
		if *failFast {
			os.Exit(1)
		}
		return
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		if *failFast {
			os.Exit(1)
		}
		return
	}

	// Print the response body
	fmt.Println(string(body))
}

