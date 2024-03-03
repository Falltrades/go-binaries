package main

import (
	"github.com/pixelbender/go-traceroute/traceroute"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go <IP address or hostname>")
	}

	ipAddr := os.Args[1]
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		addrs, err := net.LookupIP(ipAddr)
		if err != nil {
			log.Fatalf("Failed to resolve %s: %v", ipAddr, err)
		}
		// Filter out IPv6 addresses
		for _, addr := range addrs {
			if addr.To4() != nil {
				ip = addr
				break
			}
		}
		if ip == nil {
			log.Fatalf("No IPv4 address found for %s", ipAddr)
		}
	}

	hops, err := traceroute.Trace(ip)
	if err != nil {
		log.Fatal(err)
	}

	for _, h := range hops {
		for _, n := range h.Nodes {
			log.Printf("%d. %v %v", h.Distance, n.IP, n.RTT)
		}
	}
}
