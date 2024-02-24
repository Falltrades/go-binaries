package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Check if the command-line arguments are provided correctly
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <domain>")
		return
	}

	// Extract the domain from the command-line arguments
	domain := os.Args[1]

	// Perform IP lookup
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the results
	for _, ip := range ips {
		fmt.Printf("Name:    %s\n", domain)
		fmt.Printf("Address: %s\n", ip)
	}
}
