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

	partTwoRes := TwelveDigits(banks)

	fmt.Printf("PART 1 RESULT : %v\nPART 2 RESULT : %v\n", res, partTwoRes)
}

func TwelveDigits(banks []string) int {
	res := 0

	for _, bank := range banks {
		joltArr := []int{}
		for len(joltArr) < 12 {
			max := 0
			maxIndex := 0

			for x, char := range bank {
				charN, _ := strconv.Atoi(string(char))
				if charN > max {
					if len(bank)-x < 12-len(joltArr) {
						break
					}
					max = charN
					maxIndex = x
				}
			}
			joltArr = append(joltArr, max)

			if maxIndex < len(bank)-1 {
				bank = bank[maxIndex+1:]
			} else {
				break
			}
		}

		totalStr := ""
		for _, n := range joltArr {
			totalStr += strconv.Itoa(n)
		}

		totalN, _ := strconv.Atoi(totalStr)
		res += totalN

	}

	return res
}
