package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func calcRequiredFuel(mass int) int {
	// fuel := int(math.Ceil(float64(mass)/3)) - 2
	fuel := int(math.Floor(float64(mass)/3)) - 2

	fmt.Printf("mass: %d, fuel: %d\n", mass, fuel)

	return fuel
}

func getLines(input string) []string {
	result := []string{}

	for _, s := range strings.Split(input, "\n") {
		result = append(result, s)
	}

	return result
}

func sumOfReqiredFuel(inputs []string) int {
	sum := 0

	for _, i := range inputs {
		m, _ := strconv.Atoi(i)
		sum += calcRequiredFuel(m)
	}

	return sum
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	inputs := getLines(string(bytes))

	sum := sumOfReqiredFuel(inputs)
	fmt.Println(sum)
}
