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

	claimScanner := c.ScannerFile("test")
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
		height, err := strconv.Atoi(size[0])
		c.Check(err)
		width, err := strconv.Atoi(size[1])
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
	fmt.Printf("%+v\n", claims)

	return nil

}

func main() {
	splitter()
}
