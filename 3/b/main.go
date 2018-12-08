package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func parseFile(path string) []string {
	filename, _ := filepath.Abs(path)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

const fabricInches = 1000

type fabricManager struct {
	fabric       [][]int
	overlapCount int
}

func newFabricManager() *fabricManager {
	fm := &fabricManager{}
	fm.fabric = make([][]int, fabricInches)
	return fm
}

// Used to claim a point on the fabric
func (fm *fabricManager) claimPoint(x int, y int) {
	if fm.fabric[x] == nil {
		fm.fabric[x] = make([]int, fabricInches)
	}
	// Every time a claim wants this coordinate of fabric we increase
	// the number of times it's been claimed
	fm.fabric[x][y]++

	// If this is the second time we are claiming the same point increase overlap counter
	if fm.fabric[x][y] == 2 {
		fm.overlapCount++
	}
}

type claim struct {
	ID     int
	left   int
	top    int
	width  int
	height int
}

func main() {
	start := time.Now()
	claimData := parseFile("../../3/claims.txt")
	// claimData := parseFile("../../3/a/example.txt")
	var claims []claim

	// Split the strings so we can organize the data into a slice of structs
	for _, data := range claimData {
		data = strings.Replace(data, " ", "", -1) // Remove all spaces
		parts := strings.Split(data, "@")
		parts[0] = strings.Trim(parts[0], "#")
		detailParts := strings.Split(parts[1], ":")
		tlParts := strings.Split(detailParts[0], ",")
		whParts := strings.Split(detailParts[1], "x")
		ID, _ := strconv.Atoi(parts[0])
		left, _ := strconv.Atoi(tlParts[0])
		top, _ := strconv.Atoi(tlParts[1])
		width, _ := strconv.Atoi(whParts[0])
		height, _ := strconv.Atoi(whParts[1])
		c := claim{
			ID:     ID,
			left:   left,
			top:    top,
			width:  width,
			height: height,
		}
		claims = append(claims, c)
	}

	fm := newFabricManager()

	// Layout claims
	for _, c := range claims {
		// We need to know the max x and y and we loop to those values
		maxY := c.top + c.height
		maxX := c.left + c.width
		for y := c.top; y < maxY; y++ {
			for x := c.left; x < maxX; x++ {
				fm.claimPoint(x, y)
			}
		}
	}

	fmt.Println("overlap count", fm.overlapCount)

	// For part 2 check if any of the claims didn't conflict. This had to be
	// done after all claims had been made.
	for _, c := range claims {
		// We need to know the max x and y and we loop to those values
		maxY := c.top + c.height
		maxX := c.left + c.width
		overlap := false
		for y := c.top; y < maxY; y++ {
			for x := c.left; x < maxX; x++ {
				if fm.fabric[x][y] > 1 {
					overlap = true
					break
				}
			}
			if overlap == true {
				break
			}
		}

		if !overlap {
			fmt.Println("claim with no overlap", c.ID)
		}
	}

	fmt.Printf("Runtime took %s\n", time.Since(start))
}
