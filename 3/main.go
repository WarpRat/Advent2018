package main

import (
	"fmt"
	"strconv"
	"strings"

	c "github.com/WarpRat/Advent2018/pkg/common"
)

type claim struct {
	ClaimNumber int
	LeftEdge    int
	Top         int
	Height      int
	Width       int
}

func splitter() *[]claim {
	var claims []claim

	claimScanner := c.ScannerFile("input")
	for claimScanner.Scan() {
		var newClaim claim

		line := claimScanner.Text()
		lineElements := strings.Split(line, " ")
		coords := strings.Split(lineElements[2], ",")
		size := strings.Split(lineElements[3], "x")

		claimNum, err := strconv.Atoi(lineElements[0][1:])
		c.Check(err)
		left, err := strconv.Atoi(coords[0])
		c.Check(err)
		top, err := strconv.Atoi(coords[1][:len(coords[1])-1])
		c.Check(err)
		height, err := strconv.Atoi(size[1])
		c.Check(err)
		width, err := strconv.Atoi(size[0])
		c.Check(err)

		newClaim = claim{
			ClaimNumber: claimNum,
			LeftEdge:    left,
			Top:         top,
			Height:      height,
			Width:       width,
		}

		claims = append(claims, newClaim)

	}
	return &claims
}

func addressMapper(claims *[]claim) *map[int][]int {

	m := make(map[int][]int)
	for _, i := range *claims {
		for t := 0; t < i.Height; t++ {
			row := i.Top + t

			for x := 0; x < i.Width; x++ {
				le := i.LeftEdge + x
				m[row] = append(m[row], le)
			}

		}
	}
	return &m
}

func dupeFinder(adMap map[int][]int) {
	var totalCovered int

	for k := range adMap {

		rowCount := make(map[int]int, len(adMap[k]))

		for _, i := range adMap[k] {
			rowCount[i]++
		}
		for i := range rowCount {
			if rowCount[i] > 1 {
				totalCovered++
			}
		}
	}

	fmt.Printf("\nTotal inches of fabric covered: %v\n", totalCovered)
}

func main() {

	claims := splitter()
	adMap := addressMapper(claims)
	dupeFinder(*adMap)

}
