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

	lines := strings.Split(fileStr, "\n")
	matrix := [][]string{}
	for _, line := range lines {
		lineParts := []string{}
		lineElement := ""

		for _, char := range line {
			if char == ' ' {
				if lineElement != "" {
					lineParts = append(lineParts, lineElement)
					lineElement = ""
				}
			} else {
				lineElement += string(char)
			}
		}
		if lineElement != "" {
			lineParts = append(lineParts, lineElement)
		}

		matrix = append(matrix, lineParts)
	}

	//printMatrix(matrix)

	totalRes := 0
	y := 0
	for range matrix[0] {
		operator := matrix[len(matrix)-1][y]
		res := 0
		if operator == "+" {
			for x := 0; x < len(matrix)-1; x++ {
				n, _ := strconv.Atoi(matrix[x][y])
				res += n
			}
		}
		if operator == "*" {
			res = 1
			for x := 0; x < len(matrix)-1; x++ {
				n, _ := strconv.Atoi(matrix[x][y])
				res *= n
			}
		}
		totalRes += res
		fmt.Printf("%v ", res)
		if y != len(matrix)-1 {
			fmt.Print("+ ")
		}
		y++
	}
	fmt.Printf("= %v\n", totalRes)
}

func printMatrix(matrix [][]string) {
	for _, line := range matrix {
		for _, char := range line {
			fmt.Printf("'%v' ", char)
		}
		fmt.Println()
	}
}
