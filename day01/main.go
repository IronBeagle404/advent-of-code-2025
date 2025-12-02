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
			dial += rotaValue
		} else {
			dial -= rotaValue
		}
		if dial < 0 || dial > 99 {
			dial = (dial%100 + 100) % 100
		}
		if dial == 0 {
			zeroCount++
		}
		fmt.Println(dial)
	}

	fmt.Printf("\nRESULT : %v\n", zeroCount)
}
