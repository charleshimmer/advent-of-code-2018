package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type FrequencyStatus struct {
	current int
	history map[int]int
}

func checkList(changes []int, fs *FrequencyStatus) *int {
	for _, num := range changes {
		fs.current += num
		if _, ok := fs.history[fs.current]; ok {
			fmt.Println("Duplicate sum found", fs.current)
			return &fs.current
		}
		fs.history[fs.current] = 1
	}
	fmt.Println("Finished with sum of", fs.current)
	return nil
}

func parseFrequencyChanges() []int {
	var changes []int
	absPath, _ := filepath.Abs("../1/advent-01a.txt")
	f, err := os.OpenFile(absPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return nil
		}

		line = strings.TrimSuffix(line, "\n")
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("failed to convert number: %v", err)
			return nil
		}
		changes = append(changes, num)
	}
	return changes
}

func main() {
	start := time.Now()

	changes := parseFrequencyChanges()
	fs := FrequencyStatus{}
	fs.history = make(map[int]int)

	dupFound := false
	iteration := 0
	for !dupFound {
		iteration++
		fmt.Println("Iteration:", iteration)
		fmt.Println("Current:", fs.current)
		fmt.Println("History size:", len(fs.history))
		dup := checkList(changes, &fs)
		if dup != nil {
			dupFound = true
			break
		}
		if iteration >= 500 {
			fmt.Println("past max, bailing")
			break
		}
	}
	// spew.Dump(fs.history)
	// if _, present := fs.history[78]; present {
	// 	fmt.Printf("found 78")
	// }
	elapsed := time.Since(start)
	fmt.Printf("Runtime took %s", elapsed)
}
