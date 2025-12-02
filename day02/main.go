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

	ranges := strings.Split(fileStr, ",")
	res := 0

	for _, rangeVar := range ranges {
		rangeArr := strings.Split(rangeVar, "-")
		firstVal, _ := strconv.Atoi(rangeArr[0])
		secondVal, _ := strconv.Atoi(rangeArr[1])
		for x := firstVal; x <= secondVal; x++ {
			match := IsDoubleRepeat(strconv.Itoa(x))
			if match {
				fmt.Println(x)
				res += x
			}
		}
	}

	fmt.Printf("\nRESULT : %v\n", res)
}

func IsDoubleRepeat(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	half := n / 2
	return s[:half] == s[half:]
}
