package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileByes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileStr := string(fileByes)

	lines := strings.Split(fileStr, "\n")

	matrix := [][]rune{}
	for _, line := range lines {
		matrix = append(matrix, []rune(line))
	}
	//printMatrix(matrix)

	matrixCopy := make([][]rune, len(matrix))
	for i := range matrix {
		matrixCopy[i] = make([]rune, len(matrix[i]))
		copy(matrixCopy[i], matrix[i])
	}

	accessibleRolls := 0

	for x, line := range matrix {
		for y, char := range line {
			if char == '@' {
				count := 0

				// upper line check
				if x != 0 {
					for i := y - 1; i <= y+1; i++ {
						if (i < 0) || (i > len(matrix[0])-1) {
							continue
						}
						if matrix[x-1][i] == '@' {
							count++
						}
					}
				}

				// middle line check
				if y != 0 {

					if matrix[x][y-1] == '@' {
						count++
					}
				}
				if y != len(line)-1 {
					if matrix[x][y+1] == '@' {
						count++
					}
				}

				// lower line check
				if x < len(matrix)-1 {
					for i := y - 1; i <= y+1; i++ {
						if (i < 0) || (i > len(matrix[0])-1) {
							continue
						}
						if matrix[x+1][i] == '@' {
							count++
						}
					}
				}

				if count < 4 {
					matrixCopy[x][y] = 'X'
					accessibleRolls++
				}
			}
		}
	}

	fmt.Println()
	printMatrix(matrixCopy)
	fmt.Printf("\n\nACCESSIBLE ROLLS : %v\n", accessibleRolls)
}

func printMatrix(matrix [][]rune) {
	for x, line := range matrix {
		for _, char := range line {
			fmt.Printf("%c", char)
		}
		if x != len(matrix)-1 {
			fmt.Println()
		}
	}
}
