package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func executeProgram(program []int) {
	ptr := 0

	for program[ptr] != 99 {
		a := program[ptr+1]
		b := program[ptr+2]
		c := program[ptr+3]

		switch program[ptr] {
		case 1: // sum
			program[c] = program[a] + program[b]
		case 2: // multiple
			program[c] = program[a] * program[b]
		}
		ptr += 4
	}
}

func getLines(input string) []string {
	result := []string{}

	for _, s := range strings.Split(input, "\r\n") {
		result = append(result, s)
	}

	return result
}

func readProgram(input string) []int {
	result := []int{}

	parts := strings.Split(input, ",")

	for _, val := range parts {
		intVal, _ := strconv.Atoi(val)
		result = append(result, intVal)
	}

	return result
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	input := getLines(string(bytes))

	program := readProgram(input[0])
	fmt.Println(program)

	// 1202 adjustment
	program[1] = 12
	program[2] = 2

	executeProgram(program)

	fmt.Println(program[0])
}
