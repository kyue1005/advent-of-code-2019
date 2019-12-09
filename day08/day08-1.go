package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const LayerWidth = 25
const LayerHeight = 6

func countLayerDigit(d int, layer []string) int {
	count := 0
	for i := 0; i < len(layer); i++ {
		count += strings.Count(layer[i], strconv.Itoa(d))
	}
	return count
}

func getLayers(layers string) [][]string {
	LayerLength := LayerWidth * LayerHeight
	result := [][]string{}

	for i := 0; i < len(layers); i += LayerLength {
		layerLine := layers[i : i+LayerLength]

		layer := []string{}
		for j := 0; j < LayerLength; j += LayerWidth {
			layer = append(layer, layerLine[j:j+LayerWidth])
		}

		result = append(result, layer)
	}

	return result
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

	layers := getLayers(input[0])
	fmt.Println(layers)

	// Fewest zero
	targetIndex := 0
	minCount := countLayerDigit(0, layers[0])
	for i := 1; i < len(layers); i++ {
		c := countLayerDigit(0, layers[i])

		if c < minCount {
			targetIndex = i
			minCount = c
		}
	}

	answer := countLayerDigit(1, layers[targetIndex]) * countLayerDigit(2, layers[targetIndex])
	fmt.Println(answer)
}
