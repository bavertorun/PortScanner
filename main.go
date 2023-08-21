/*
* Github: @BaverTorun
*/
package main

import (
	"fmt"

	"github.com/bavertorun/portScanner/port"
)

func main() {
	// Scanning ports on "localhost"
	results := port.InitialScan("localhost")
	fmt.Println(results) // Printing results to the console
}
