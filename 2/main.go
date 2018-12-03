package main

import (
	"fmt"
	"strings"

	c "github.com/WarpRat/Advent2018/pkg/common"
)

func oneOffScan() (*[]string, *string) {
	var inputArray []string
	var oneOff []string
	var commonLetters string

	boxScan := c.ScannerFile("input")
	for boxScan.Scan() {
		id := boxScan.Text()
		inputArray = append(inputArray, id)
	}

	for _, line := range inputArray {
		for _, compLine := range inputArray {
			var dissimilar int

			for i, char := range line {
				if char != rune(compLine[i]) {
					dissimilar++
				}
			}
			if dissimilar == 1 {
				oneOff = append(oneOff, line)
			}
		}
	}

	for i, letter := range oneOff[0] {
		if letter == rune(oneOff[1][i]) {
			commonLetters = fmt.Sprintf("%v%v", commonLetters, string(letter))
		}
	}

	return &oneOff, &commonLetters
}

func main() {
	var twoMatch []string
	var threeMatch []string
	var lastID string

	boxScan := c.ScannerFile("input")

	for boxScan.Scan() {
		id := boxScan.Text()

		if lastID == "" {
			lastID = id
		} else {

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
		lastID = id

	}

	oneOff, commonLetters := oneOffScan()

	fmt.Printf("Two: %v\n\nThree: %v\n", twoMatch, threeMatch)
	fmt.Printf("\n\nRESULTS:\nTwo length: %v\nThree length: %v\n", len(twoMatch), len(threeMatch))
	fmt.Printf("Checksum: %v\n", (len(twoMatch) * len(threeMatch)))
	fmt.Printf("Labels only one digit off: %v\n", *oneOff)
	fmt.Printf("Common letters: %v\n", *commonLetters)

}
