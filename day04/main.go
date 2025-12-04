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

	displayMatrix := make([][]rune, len(matrix))
	for i := range matrix {
		displayMatrix[i] = make([]rune, len(matrix[i]))
		copy(displayMatrix[i], matrix[i])
	}

	totalRolls := 0

	// weird init, can't think of another way rn
	accessibleRolls := 1
	turns := 0
	for accessibleRolls != 0 {
		accessibleRolls = 0

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
						displayMatrix[x][y] = 'A'
						accessibleRolls++
					}
				}
			}
		}

		if turns == 0 {
			fmt.Printf("Initial state:\n")
		}
		printMatrix(displayMatrix)
		fmt.Printf("\n")
		if accessibleRolls != 0 {
			fmt.Printf("\nRemove %v roll", accessibleRolls)
			if accessibleRolls > 1 {
				fmt.Print("s")
			}
			fmt.Printf(" of paper\n")
		}

		for x, line := range displayMatrix {
			for y := range line {
				if displayMatrix[x][y] == 'X' {
					displayMatrix[x][y] = '.'
				}
			}
		}

		for x, line := range displayMatrix {
			for y := range line {
				if displayMatrix[x][y] == 'A' {
					displayMatrix[x][y] = 'X'
				}
			}
		}

		for x, line := range matrix {
			for y := range line {
				if displayMatrix[x][y] == 'X' {
					matrix[x][y] = '.'
				}
			}
		}

		totalRolls += accessibleRolls
		turns++
	}

	fmt.Printf("\n\n%v rolls can be removed\n", totalRolls)
}

func printMatrix(matrix [][]rune) {
	for x, line := range matrix {
		for _, char := range line {
			if char == 'A' {
				fmt.Printf("\033[31m@\033[0m")
			} else if char == 'X' {
				fmt.Printf("%c", 'x')
			} else {
				fmt.Printf("%c", char)
			}
		}
		if x != len(matrix)-1 {
			fmt.Println()
		}
	}
}
