package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func checkNumber(i int) bool {
	isAcending := true
	hasDouble := false

	num := splitNumber(i)

	for i := 0; i < len(num)-1; i++ {
		if num[i] > num[i+1] {
			isAcending = false
		}
		if num[i] == num[i+1] {
			hasDouble = true
		}
	}
	return isAcending && hasDouble
}

func splitNumber(i int) []int {
	digits := []int{}

	pow := 1

	for pow <= i {
		d := (i / pow) % 10
		digits = append([]int{d}, digits...)
		pow *= 10
	}

	return digits
}

func getInputLines(input string) []string {
	result := []string{}

	for _, s := range strings.Split(input, "\n") {
		result = append(result, s)
	}

	return result
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	input := getInputLines(string(bytes))

	parts := strings.Split(input[0], "-")
	min, _ := strconv.Atoi(parts[0])
	max, _ := strconv.Atoi(parts[1])

	count := 0
	for x := min; x <= max; x++ {
		if checkNumber(x) {
			// fmt.Println(x)
			count++
		}
	}

	fmt.Println(count)
}
