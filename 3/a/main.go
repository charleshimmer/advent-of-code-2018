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

type claim struct {
	ID     int
	left   int
	top    int
	width  int
	height int
}

func main() {
	start := time.Now()

	claims := parseFile("../../3/claims.txt")

	for _, data := range claims {
		fmt.Println(data)
	}

	fmt.Printf("Runtime took %s\n", time.Since(start))
}
