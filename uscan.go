package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	// Define command-line flags
	address := flag.String("address", "localhost", "the hostname or IP address to scan")
	cidr := flag.String("cidr", "", "the CIDR range to scan (e.g. 192.168.1.0/24)")
	scanType := flag.String("type", "tcp", "the type of scan to perform (tcp or udp)")
	delay := flag.Int("delay", 100, "the delay between connections in milliseconds")
	flag.Parse()

	// Validate the scan type
	if *scanType != "tcp" && *scanType != "udp" {
		fmt.Println("Invalid scan type. Must be 'tcp' or 'udp'.")
		return
	}

	// Parse the address or CIDR range
	var addresses []string
	if *cidr != "" {
		// If a CIDR range is provided, expand it into a list of individual addresses
		ip, ipnet, err := net.ParseCIDR(*cidr)
		if err != nil {
			fmt.Println("Invalid CIDR range:", err)
			return
		}
		for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
			addresses = append(addresses, ip.String())
							} else {
					fmt.Printf("\x1b[31m%d: Closed\x1b[0m\n", port)
				}
				continue
			}
			// Otherwise, the port is open
			fmt.Printf("\x1b[32m%d: Open\x1b[0m\n", port)
			conn.Close()
		}
	}
}

// inc increments the last octet of an IP address
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

		}
	} else {
		// Otherwise, just add the single address
		addresses = append(addresses, *address)
	}

	// Loop through the addresses
	for _, address := range addresses {
		// Try to resolve the address to an IP address
		ipAddr, err := net.ResolveIPAddr("ip", address)
		if err != nil {
			fmt.Printf("Unable to resolve '%s': %v\n", address, err)
			continue
		}

		// Scan the ports for the address
		fmt.Printf("Scanning %s (%s)...\n", address, ipAddr.String())
		for port := 1; port <= 65535; port++ {
			// Try to connect to the host on the current port
			conn, err := net.DialTimeout(*scanType, fmt.Sprintf("%s:%d", ipAddr.String(), port), time.Millisecond*time.Duration(*delay))

			// If there is an error, the port is closed or filtered
			if err != nil {
				if strings.Contains(err.Error(), "too many open files") {
					fmt.Printf("\x1b[33m%d: Too many open files\x1b[0m\n", port)
