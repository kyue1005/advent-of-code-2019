package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getLines(input string) []string {
	result := []string{}

	for _, s := range strings.Split(input, "\r\n") {
		result = append(result, s)
	}

	return result
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	input := getLines(string(bytes))

	fmt.Println(input)
}
