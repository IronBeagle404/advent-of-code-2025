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
	firstRes := 0
	secondRes := 0

	for _, rangeVar := range ranges {
		rangeArr := strings.Split(rangeVar, "-")
		firstVal, _ := strconv.Atoi(rangeArr[0])
		secondVal, _ := strconv.Atoi(rangeArr[1])
		for x := firstVal; x <= secondVal; x++ {
			match := IsDoubleRepeat(strconv.Itoa(x))
			if match {
				firstRes += x
			}
			match = HasRepeatedPattern(strconv.Itoa(x))
			if match {
				secondRes += x
			}
		}
	}

	fmt.Printf("\nPART 1 RESULT : %v\nPART 2 RESULT : %v\n", firstRes, secondRes)
}

func IsDoubleRepeat(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	half := n / 2
	return s[:half] == s[half:]
}

func HasRepeatedPattern(s string) bool {
	n := len(s)
	if n < 2 {
		return false
	}
	for size := 1; size <= n/2; size++ {
		if n%size != 0 {
			continue
		}
		block := s[:size]
		ok := true
		for i := size; i < n; i += size {
			if s[i:i+size] != block {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}
