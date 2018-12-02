package adventcommon

import (
	"bufio"
	"os"
)

//Check makes error checking easy
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

//ScannerFile returns a bufio scanner for the desired file
func ScannerFile(filename string) *bufio.Scanner {

	file, err := os.Open(filename)
	Check(err)
	scanner := bufio.NewScanner(file)
	return scanner
}
