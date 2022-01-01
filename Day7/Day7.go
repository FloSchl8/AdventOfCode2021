package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./Day7/input.txt")

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	input := strings.Split(scanner.Text(), ",")

	numbers := make(map[int]int)
	m := new(median)

	for i := 0; i < len(input); i++ {
		number, _ := strconv.Atoi(input[i])
		m.numbers = append(m.numbers, number)
		numbers[number]++
	}

	optimalAlgin := m.calcMedian()
	fmt.Println("optimal align:", optimalAlgin)

	fuelCostTotal := 0

	for number, count := range numbers {
		fuelCost := math.Abs(float64(optimalAlgin - number))
		fuelCostTotal += count * int(fuelCost)
	}

	fmt.Println("total fuel:", fuelCostTotal)
}

func (m median) calcMedian() int {

	sort.Ints(m.numbers)

	middle := len(m.numbers) / 2

	if len(m.numbers)%2 == 0 {
		return m.numbers[middle]
	}

	return (m.numbers[middle-1] + m.numbers[middle]) / 2
}

type median struct {
	numbers []int
}
