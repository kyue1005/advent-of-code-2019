package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// Line ...
type Line struct {
	xMin   int
	xMax   int
	yMin   int
	yMax   int
	xDelta int
	yDelta int
}

// Point ...
type Point struct {
	x int
	y int
}

func getWireLines(input string) []Line {
	pt := Point{x: 0, y: 0}
	lines := []Line{}
	parts := strings.Split(input, ",")

	for _, line := range parts {
		lineChar := []rune(line)

		dir := lineChar[0]
		len, _ := strconv.Atoi(string(lineChar[1:]))

		switch dir {
		case 'U':
			lines = append(lines, Line{
				yMin:   pt.y,
				yMax:   pt.y + len,
				xMin:   pt.x,
				xMax:   pt.x,
				xDelta: 0,
				yDelta: len,
			})

			pt.y += len
		case 'R':
			lines = append(lines, Line{
				yMin:   pt.y,
				yMax:   pt.y,
				xMin:   pt.x,
				xMax:   pt.x + len,
				xDelta: len,
				yDelta: 0,
			})

			pt.x += len
		case 'D':
			lines = append(lines, Line{
				yMin:   pt.y - len,
				yMax:   pt.y,
				xMin:   pt.x,
				xMax:   pt.x,
				xDelta: 0,
				yDelta: -len,
			})

			pt.y -= len
		case 'L':
			lines = append(lines, Line{
				yMin:   pt.y,
				yMax:   pt.y,
				xMin:   pt.x - len,
				xMax:   pt.x,
				xDelta: -len,
				yDelta: 0,
			})

			pt.x -= len
		}
	}

	return lines
}

func getLineIntersection(l1 Line, l2 Line) []Point {
	pts := []Point{}

	for x := l1.xMin; x <= l1.xMax; x++ {
		for y := l1.yMin; y <= l1.yMax; y++ {
			// skip central point
			if x != 0 && y != 0 {
				if x >= l2.xMin && x <= l2.xMax && y >= l2.yMin && y <= l2.yMax {
					pts = append(pts, Point{x: x, y: y})
				}
			}
		}
	}

	return pts
}

func getWireIntersections(wires [][]Line) []Point {
	intersections := []Point{}

	for i := 0; i < len(wires); i++ {
		for j := i + 1; j < len(wires); j++ {
			for _, l1 := range wires[i] {
				for _, l2 := range wires[j] {
					pts := getLineIntersection(l1, l2)
					if len(pts) > 0 {
						intersections = append(intersections, pts...)
					}
				}
			}
		}
	}

	return intersections
}

func getWireStep(wire []Line, pt Point) int {
	ptr := Point{x: 0, y: 0}
	step := 0

	// fmt.Println(pt)

	for _, line := range wire {
		if line.xDelta != 0 {
			if line.xMin <= pt.x && line.xMax >= pt.x && line.yMin == pt.y {
				step += int(math.Abs(float64(ptr.x - pt.x)))
				break
			} else {
				step += int(math.Abs(float64(line.xDelta)))
				ptr.x += line.xDelta
			}
		} else if line.yDelta != 0 {
			if line.yMin <= pt.y && line.yMax >= pt.y && line.xMin == pt.x {
				step += int(math.Abs(float64(ptr.y - pt.y)))
				break
			} else {
				step += int(math.Abs(float64(line.yDelta)))
				ptr.y += line.yDelta
			}
		}

		// fmt.Println(line, ptr, step)
	}

	return step
}

func getLeastStepsIntersection(wires [][]Line, pts []Point) int {
	minVal := 0
	for _, pt := range pts {
		step := 0
		for _, w := range wires {
			step += getWireStep(w, pt)
		}

		fmt.Println(pt, step)

		if minVal == 0 || step < minVal {
			minVal = step
		}
	}

	return minVal
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

	wires := [][]Line{}

	for _, line := range input {
		l := getWireLines(line)
		wires = append(wires, l)
	}
	fmt.Println(wires)

	interPts := getWireIntersections(wires)
	fmt.Println(interPts)
	answer := getLeastStepsIntersection(wires, interPts)

	fmt.Println(answer)
}
