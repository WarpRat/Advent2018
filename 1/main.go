package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/WarpRat/Advent2018/pkg/common"
)

func calculator(frequency int, op []byte, nextNum int) int {

	if bytes.Equal(op, []byte("-")) {
		return (frequency - nextNum)
	} else if bytes.Equal(op, []byte("+")) {
		return (frequency + nextNum)
	} else {
		panic("Failed")

	}

}

func scanLooper() (int, int) {

	var allFrq []int
	var finalFrq int
	var frequency int

	//Presume that we won't need to do more than 10k loops.
	for loops := 0; loops < 10000; loops++ {

		inputScanner := adventreader.ScannerFile("input")

		for inputScanner.Scan() {
			line := inputScanner.Bytes()
			nextNum, err := strconv.Atoi(string(line[1:]))
			if err != nil {
				panic(err)
			}

			frequency = calculator(frequency, line[:1], nextNum)

			for _, x := range allFrq {

				if frequency == x {
					return frequency, finalFrq
				}
			}

			allFrq = append(allFrq, frequency)
		}

		if loops == 0 {
			finalFrq = frequency
		}
	}

	panic("No match found in 10,000 loops. You can try increasing it or check your code.")
}

func main() {

	dbl, frq := scanLooper()

	fmt.Printf("Final frequency is: %v\n", frq)
	fmt.Printf("First double was: %v\n", dbl)

}
