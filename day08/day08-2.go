package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const LayerWidth = 25
const LayerHeight = 6

var LayerLength int

func getLayers(layers string) [][]string {
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

func getPixel(col int, row int, layers [][]string) string {
	val := ""
	for i := 0; i < len(layers); i++ {
		row := []rune(layers[i][row])
		// fmt.Print(string(row[col]))
		if row[col] != '2' {
			switch row[col] {
			case '1':
				val = "+"
			case '0':
				val = " "
			}
			break
		}
	}

	return val
}

func decodeLayers(layers [][]string) []string {
	result := []string{}

	for i := 0; i < LayerHeight; i++ {
		layerRow := ""
		for j := 0; j < LayerWidth; j++ {
			layerRow += getPixel(j, i, layers)
		}
		fmt.Println(layerRow)
		result = append(result, layerRow)
	}

	return result
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	input := getInputLines(string(bytes))

	LayerLength = LayerWidth * LayerHeight
	layers := getLayers(input[0])
	// fmt.Println(layers)

	decodeLayers(layers)
}
