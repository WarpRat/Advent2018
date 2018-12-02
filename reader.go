package adventreader

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//ScannerFile returns a bufio scanner for the desired file
func ScannerFile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	return scanner
}
