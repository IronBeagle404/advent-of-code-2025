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

	dial := 50
	zeroCount := 0

	rotations := strings.Split(fileStr, "\n")
	for _, rota := range rotations {
		rotaValue, _ := strconv.Atoi(rota[1:])
		if rota[0] == 'R' {
			for x := 0; x < rotaValue; x++ {
				if dial == 99 {
					dial = 0
				} else {
					dial++
				}
				if dial == 0 {
					zeroCount++
				}
			}
		} else {
			for x := 0; x < rotaValue; x++ {
				if dial == 0 {
					dial = 99
				} else {
					dial--
				}
				if dial == 0 {
					zeroCount++
				}
			}
		}
		fmt.Println(dial)
	}

	fmt.Printf("\nRESULT : %v\n", zeroCount)
}
