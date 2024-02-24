package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Ullaakut/nmap"
)

func main() {
	// Command line flags
	flag.Parse()
	args := flag.Args()

	// Check if target is provided
	if len(args) < 1 {
		log.Fatal("Please provide a target")
	}

	// Parse target
	target := args[0]

	// Parse port range or set default
	var ports string
	if len(args) > 1 {
		ports = args[1]
	} else {
		ports = "1-1024"
	}

	// Create a new scanner
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(target),
		nmap.WithPorts(ports),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	// Run the scan
	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	// Print any warnings
	if warnings != nil {
		fmt.Printf("Warnings: \n%s\n", warnings)
	}

	// Process scan results
	for _, host := range result.Hosts {
		fmt.Printf("Host: %s\n", host.Addresses[0])
		for _, port := range host.Ports {
			fmt.Printf("Port: %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}
}
