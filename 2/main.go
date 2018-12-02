package main

import (
	"fmt"
	"strings"

	c "github.com/WarpRat/Advent2018/pkg/common"
)

func main() {
	var twoMatch []string
	var threeMatch []string

	boxScan := c.ScannerFile("input")

	for boxScan.Scan() {
		id := boxScan.Text()

		var match2 bool
		var match3 bool

		for _, char := range id {

			if strings.Count(id, string(char)) == 2 {
				match2 = true
			} else if strings.Count(id, string(char)) == 3 {
				match3 = true
			}

		}

		if match2 {
			twoMatch = append(twoMatch, string(id))
		}
		if match3 {
			threeMatch = append(threeMatch, string(id))
		}

	}
	fmt.Printf("Two: %v\n\nThree: %v\n", twoMatch, threeMatch)
	fmt.Printf("\n\nRESULTS:\nTwo length: %v\nThree length: %v\n", len(twoMatch), len(threeMatch))
	fmt.Printf("Checksum: %v\n", (len(twoMatch) * len(threeMatch)))

}
