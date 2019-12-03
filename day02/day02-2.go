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
	for noun := 0; noun <= 99; noun++ {
		match := false

		for verb := 0; verb <= 99; verb++ {
			tempProgram := make([]int, len(program))
			copy(tempProgram, program)
			tempProgram[1] = noun
			tempProgram[2] = verb

			executeProgram(tempProgram)

			fmt.Println(noun, verb, tempProgram[0])
			
			if tempProgram[0] == 19690720 {
				fmt.Println("Answer", noun, verb)
        match = true
				break
			}
		}

		if match {
			break
		}
	}

	
}
