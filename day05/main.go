package main

import (
	"fmt"
	"log"
	"os"
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

	freshCount := 0
	availableIds := strings.Split(splitted[1], "\n")

	for _, IdStr := range availableIds {
		ID, err := strconv.Atoi(IdStr)
		if err != nil {
			log.Fatal(err)
		}

		fresh := false
		for _, rangeStr := range freshRanges {
			split := strings.Split(rangeStr, "-")
			lowRange, err := strconv.Atoi(split[0])
			if err != nil {
				log.Fatal(err)
			}
			highRange, err := strconv.Atoi(split[1])
			if err != nil {
				log.Fatal(err)
			}
			if ID >= lowRange && ID <= highRange {
				fresh = true
				break
			}
		}
		if fresh {
			freshCount++
		}
	}

	fmt.Printf("FRESH INGREDIENTS AVAILABLE : %v\n", freshCount)
}
