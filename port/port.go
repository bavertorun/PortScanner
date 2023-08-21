/*
* Github: @BaverTorun
*/
package port

import (
	"net"
	"strconv"
	"time"
)

// ScanResult represents the data structure used to store port scanning results.
type ScanResult struct {
	Port  string // Name and number of the scanned port (e.g., "tcp/80")
	State string // State of the port ("Open" or "Closed")
}

// ScanPort scans the status of a port based on the specified protocol, hostname, and port number.
func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}

	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	// Checking the result
	if err != nil {
		result.State = "Closed\n"
		return result
	}

	defer conn.Close()
	result.State = "Open\n"
	return result
}

// InitialScan scans all TCP and UDP ports between 1 and 1024 on the specified hostname.
func InitialScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 1; i <= 1024; i++ {
		// Scanning TCP port
		results = append(results, ScanPort("tcp", hostname, i))
		// Scanning UDP port
		results = append(results, ScanPort("udp", hostname, i))
	}
	return results
}
