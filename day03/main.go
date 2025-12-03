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
	banks := strings.Split(fileStr, "\n")

	res := 0

	for _, bank := range banks {
		max := 0
		maxIndex := 0

		for x, char := range bank {
			charN, _ := strconv.Atoi(string(char))
			if (charN > max) && (x != len(bank)-1) {
				max = charN
				maxIndex = x
			}
		}
		tens := strconv.Itoa(max)

		max = 0
		newBank := bank[maxIndex+1:]
		for _, char := range newBank {
			charN, _ := strconv.Atoi(string(char))
			if charN > max {
				max = charN
			}
		}
		units := strconv.Itoa(max)

		totalValue, _ := strconv.Atoi(tens + units)
		res += totalValue
	}

	fmt.Printf("RESULT : %v\n", res)
}
