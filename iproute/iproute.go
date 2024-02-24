package main

import (
    "bufio"
    "fmt"
    "os"
    "net"
    "strconv"
    "strings"
)

func main() {
    if len(os.Args) > 1 && os.Args[1] == "r" {
        printRoutingTable()
    } else {
        printInterfaceAddresses()
    }
}

func printInterfaceAddresses() {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    for _, addr := range addrs {
        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                ones, _ := ipnet.Mask.Size()
                fmt.Printf("IPv4 Address: %s/%d\n", ipnet.IP.String(), ones)
            }
        }
    }
}

func printRoutingTable() {
    file, err := os.Open("/proc/net/route")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        if len(fields) >= 3 && fields[1] != "Iface" {
            dest := hexToIP(fields[1])
            gateway := hexToIP(fields[2])
            iface := fields[0]
            fmt.Printf("%s via %s dev %s\n", dest, gateway, iface)
        }
    }
}

func hexToIP(hexStr string) string {
    value, _ := strconv.ParseUint(hexStr, 16, 64)
    a := byte(value & 0xff)
    b := byte((value >> 8) & 0xff)
    c := byte((value >> 16) & 0xff)
    d := byte((value >> 24) & 0xff)
    return fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)
}
