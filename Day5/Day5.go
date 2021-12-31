package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), " -> ")
		start := convertToPoint(strings.Split(coords[0], ","))
		end := convertToPoint(strings.Split(coords[1], ","))
		lines = append(lines, Line{
			Start: start,
			End:   end,
		})
	}

	grid := plotGrid(lines)
	count := 0

	for _, i := range grid {
		if i > 1 {
			count++
		}
	}

	fmt.Println(count)

}

func plotGrid(lines []Line) map[Point]int {

	dotGrid := make(map[Point]int)

	for _, line := range lines {
		start := line.Start
		end := line.End
		v, h := line.getLengths()
		vertical := int(math.Abs(float64(v)))
		horizontal := int(math.Abs(float64(h)))

		verticalSign := 1
		horizontalSign := 1

		if math.Signbit(float64(v)) {
			verticalSign = -1
		}
		if math.Signbit(float64(h)) {
			horizontalSign = -1
		}

		// dots
		if start.equal(end) {
			dotGrid[start]++
		}
		// diagonal
		if vertical == horizontal {
			for i := 0; i <= vertical; i++ {
				point := Point{
					X: start.X + (i * verticalSign),
					Y: start.Y + (i * horizontalSign),
				}
				dotGrid[point]++
			}
		}
		// vertical lines
		if horizontal == 0 {
			for i := 0; i <= vertical; i++ {
				point := Point{
					X: start.X + (i * verticalSign),
					Y: start.Y,
				}
				dotGrid[point]++
			}
		}
		// horizontal lines
		if vertical == 0 {
			for i := 0; i <= horizontal; i++ {
				point := Point{
					X: start.X,
					Y: start.Y + (i * horizontalSign),
				}
				dotGrid[point]++
			}
		}

	}

	return dotGrid
}

func convertToPoint(coords []string) Point {
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])

	return Point{
		X: x,
		Y: y,
	}
}

func (p Point) equal(p2 Point) bool {
	if p.X == p2.X && p.Y == p2.Y {
		return true
	}
	return false
}

type Point struct {
	X int
	Y int
}

func (l Line) getLengths() (int, int) {
	start := l.Start
	end := l.End

	vertical := end.X - start.X
	horizontal := end.Y - start.Y

	return vertical, horizontal
}

type Line struct {
	Start Point
	End   Point
}
