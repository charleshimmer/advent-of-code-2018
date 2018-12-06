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

// We need to return the position of the ONLY character diff between the two strings
func findOnlyDiffPosition(s1 string, s2 string) *int {
	var onlyDiffPosition *int
	fmt.Println("-----------------------------")
	fmt.Println("012345678901234567890123456")
	fmt.Println(s1)
	fmt.Println(s2)

	for i := range s1 {
		if s1[i] != s2[i] {
			if onlyDiffPosition == nil {
				onlyDiffPosition = &i
			} else {
				return nil
			}
		}
	}
	fmt.Println("-----------------------------")
	return onlyDiffPosition
}

func main() {
	start := time.Now()

	boxIDs := parseFile("../../2/box-ids.txt")

	// Need to loop and check every boxID against all the following
	// boxIDs in the group.
	var diffPosition *int
	var s1 string
	var s2 string
	for i := range boxIDs {
		for j := i + 1; j < len(boxIDs); j++ {
			s1 = boxIDs[i]
			s2 = boxIDs[j]
			diffPosition = findOnlyDiffPosition(s1, s2)
			if diffPosition != nil {
				break
			}
		}
		// Allows us to break out if we find the diff
		if diffPosition != nil {
			break
		}
	}

	fmt.Println("Diff position is", *diffPosition, "(", string(s1[*diffPosition]), ")", "between", s1, "and", s2)
	fmt.Printf("Runtime took %s\n", time.Since(start))
}
