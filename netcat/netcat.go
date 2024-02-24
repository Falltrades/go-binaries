package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	var isUDP bool
	var timeout int
	flag.BoolVar(&isUDP, "u", false, "Specify UDP protocol")
	flag.IntVar(&timeout, "w", 10, "Specify timeout duration in seconds")
	flag.Parse()

	if len(flag.Args()) != 2 && !(len(flag.Args()) == 3 && flag.Arg(1) == "-u") {
		fmt.Println("Usage: go run netcat.go [-u] [-w <timeout>] <host> <port>")
		fmt.Println("Timeout defaulted to 10s. Also allow to specify port range.")
		return
	}

	host := flag.Arg(0)
	portSpec := flag.Arg(len(flag.Args()) - 1)

	var protocol string
	if isUDP {
		protocol = "udp"
	} else {
		protocol = "tcp"
	}

	ports := parsePorts(portSpec)

	for _, port := range ports {
		// Attempt connection
		conn, err := net.DialTimeout(protocol, fmt.Sprintf("%s:%d", host, port), time.Duration(timeout)*time.Second)
		if err == nil {
			conn.Close()
			fmt.Printf("Connection to %s %d port [%s] succeeded!\n", host, port, protocol)
		} else {
			fmt.Printf("Connection to %s %d port [%s] failed: %s\n", host, port, protocol, err)
		}
	}
}

func parsePorts(portSpec string) []int {
	var ports []int
	if strings.Contains(portSpec, "-") {
		parts := strings.Split(portSpec, "-")
		if len(parts) != 2 {
			fmt.Println("Invalid port range format. Please use startPort-endPort format.")
			return ports
		}

		startPort, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid start port:", err)
			return ports
		}

		endPort, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Invalid end port:", err)
			return ports
		}

		if startPort > endPort {
			fmt.Println("Start port must be less than or equal to end port.")
			return ports
		}

		for port := startPort; port <= endPort; port++ {
			ports = append(ports, port)
		}
	} else {
		port, err := strconv.Atoi(portSpec)
		if err != nil {
			fmt.Println("Invalid port:", err)
			return ports
		}
		ports = append(ports, port)
	}
	return ports
}
