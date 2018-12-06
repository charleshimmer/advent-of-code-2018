package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
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

// Takes in a string and returns a map of the letter and the number of times
// it occured in the string
func letterCountMap(s string) map[string]int {
	lcMap := make(map[string]int)
	for _, r := range s {
		str := string(r)
		if _, present := lcMap[str]; present {
			lcMap[str]++
		} else {
			lcMap[str] = 1
		}
	}
	return lcMap
}

func main() {
	start := time.Now()

	boxIDs := parseFile("../2/box-ids.txt")
	twoTimesCount := 0
	threeTimesCount := 0

	for _, boxID := range boxIDs {
		lcMap := letterCountMap(boxID)
		countedForTwo := false
		countedForThree := false
		for _, num := range lcMap {
			// A boxID can count towards a 2 and / or 3 but not multiple times
			if num == 2 && !countedForTwo {
				twoTimesCount++
				countedForTwo = true
			}
			if num == 3 && !countedForThree {
				threeTimesCount++
				countedForThree = true
			}
		}
	}

	fmt.Printf("Checksum is %d\n", twoTimesCount*threeTimesCount)
	fmt.Printf("Runtime took %s\n", time.Since(start))
}
