package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func toInt(b byte) int {
	i, _ := strconv.Atoi(string(b))
	return i
}

/* find the sum of all digits that match the next digit in the list.
The list is circular, so the digit after the last digit is the
first digit in the list. */
func captchaP1(input string) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		next := input[(i+1)%len(input)]
		if input[i] == next {
			sum += toInt(input[i])
		}
	}
	return sum
}

func captchaP2(input string) int {
	offset := len(input) / 2
	sum := 0

	for i := 0; i < len(input); i++ {
		next := input[(i+offset)%len(input)]
		if input[i] == next {
			sum += toInt(input[i])
		}
	}

	return sum
}

func main() {
	fmt.Println("AOC - Day 1")
	raw, _ := ioutil.ReadFile("./data.txt")
	inputStr := strings.TrimSpace(string(raw))

	fmt.Printf("Result part1: %d\n", captchaP1(inputStr))
	fmt.Printf("Result part2: %d\n", captchaP2(inputStr))
}
