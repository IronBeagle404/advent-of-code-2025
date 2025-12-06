package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileStr := string(fileBytes)
	splitted := strings.Split(fileStr, "\n\n")

	freshRanges := strings.Split(splitted[0], "\n")
	freshIds := []int{}

	for _, line := range freshRanges {
		rangeParts := strings.Split(line, "-")
		start, err := strconv.Atoi(rangeParts[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(rangeParts[1])
		if err != nil {
			log.Fatal(err)
		}

		for x := start; x <= end; x++ {
			freshIds = append(freshIds, x)
		}
	}

	freshCount := 0
	availableIds := strings.Split(splitted[1], "\n")
	for _, IdStr := range availableIds {
		ID, err := strconv.Atoi(IdStr)
		if err != nil {
			log.Fatal(err)
		}
		if slices.Contains(freshIds, ID) {
			freshCount++
		}
	}

	fmt.Printf("FRESH INGREDIENTS AVAILABLE : %v\n", freshCount)
}
