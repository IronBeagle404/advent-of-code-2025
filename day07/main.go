package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileStr := string(fileBytes)

	matrix := [][]string{}
	lines := strings.Split(fileStr, "\n")
	for _, line := range lines {
		chars := strings.Split(line, "")
		matrix = append(matrix, chars)
	}

	//printMatrix(matrix)

	splits := 0
	for x, line := range matrix {
		for y, char := range line {
			if char == "S" {
				matrix[x+1][y] = "|"

			}

			if x != 0 && matrix[x-1][y] == "|" {
				if char == "^" {
					matrix[x][y-1] = "|"
					matrix[x][y+1] = "|"
					splits++
				}
				if char == "." {
					matrix[x][y] = "|"
				}
			}
		}
	}

	fmt.Println()
	printMatrix(matrix)

	fmt.Printf("\nThe tachyon beam is split a total of %v times\n", splits)
}

func printMatrix(matrix [][]string) {
	for _, line := range matrix {
		for _, char := range line {
			fmt.Printf("%v ", char)
		}
		fmt.Println()
	}
}
