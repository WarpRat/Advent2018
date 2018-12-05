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
	Squares     [][2]int
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

		m := make(map[int][]int)
		for t := 0; t < height; t++ {
			row := top + t

			for x := 0; x < width; x++ {
				var newCoords [2]int
				le := left + x
				m[row] = append(m[row], le)
				newCoords = [2]int{row, le}
				newClaim.Squares = append(newClaim.Squares, newCoords)

			}

		}

		claims = append(claims, newClaim)
	}
	return &claims
}

func dupeFinder(claims *[]claim) {
	allSquares := make(map[[2]int]int)

	for _, i := range *claims {
		for y := range i.Squares {
			allSquares[i.Squares[y]]++
		}
	}
	fmt.Printf("%+v", allSquares)

	var totalCovered int
	for i := range allSquares {
		if allSquares[i] > 1 {
			totalCovered++
			delete(allSquares, i)
		}
	}

	fmt.Printf("Total covered squares: %v\n", totalCovered)

	/* TODO #### COME  BACK HERE
	for _, i := range *claims {
		for y := range i.Squares {

		}
	}

	*/

	/*
		var totalCovered int
		fullMap := make(map[int]map[int]int)

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
			fullMap[k] = rowCount
		}

		fmt.Printf("\nTotal inches of fabric covered: %v\n", totalCovered)
	*/
}

func main() {

	claims := splitter()
	dupeFinder(claims)
	//fmt.Printf("%+v", claims)

}
