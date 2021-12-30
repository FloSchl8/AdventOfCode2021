package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day5/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	lines := make([]Line, 0)

	max := 0

	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), " -> ")
		start := convertToPoint(strings.Split(coords[0], ","))
		end := convertToPoint(strings.Split(coords[1], ","))
		lines = append(lines, Line{
			Start: start,
			End:   end,
		})
		if start.getMaxValue() > max {
			max = start.getMaxValue()
		}
		if end.getMaxValue() > max {
			max = end.getMaxValue()
		}
	}

	lines = filterLines(lines)

	grid := plotGrid(lines, max+1)

	fmt.Println(getNumberOfIntersections(grid))
}

func getNumberOfIntersections(grid [][]int) int {
	count := 0

	for _, ints := range grid {
		for _, i := range ints {
			if i > 1 {
				count++
			}
		}
	}

	return count
}

func plotGrid(lines []Line, gridSize int) [][]int {

	grid := make([][]int, gridSize)

	for i, ints := range grid {
		if ints == nil {
			ints = make([]int, gridSize)
			grid[i] = ints
		}
	}

	for _, line := range lines {
		if line.Start.X == line.End.X {
			for i := line.Start.Y; i <= line.End.Y; i++ {
				grid[line.Start.X][i]++
			}
		}

		if line.Start.Y == line.End.Y {
			for i := line.Start.X; i <= line.End.X; i++ {
				grid[i][line.Start.Y]++
			}
		}
	}

	return grid
}

func filterLines(lines []Line) []Line {

	newLines := make([]Line, 0)

	for _, line := range lines {
		if line.Start.hasEqualXorY(line.End) {
			newLine := Line{
				Start: line.Start,
				End:   line.End,
			}
			newLine = newLine.switchStartAndEnd()
			newLines = append(newLines, newLine)
		}
	}
	return newLines
}

func convertToPoint(coords []string) Point {
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])

	return Point{
		X: x,
		Y: y,
	}
}

func (p Point) hasEqualXorY(point Point) bool {
	if p.X == point.X || p.Y == point.Y {
		return true
	}
	return false
}

func (p Point) getMaxValue() int {

	if p.X > p.Y {
		return p.X
	} else {
		return p.Y
	}
}

type Point struct {
	X int
	Y int
}

func (l Line) switchStartAndEnd() Line {
	if l.Start.X > l.End.X || l.Start.Y > l.End.Y {
		l.Start, l.End = l.End, l.Start
	}
	return l
}

type Line struct {
	Start Point
	End   Point
}
