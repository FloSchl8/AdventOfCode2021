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

	optimalAlign := 0
	fuelCostTotal := math.MaxInt

	for align := 0; align < 1000; align++ {

		fmt.Println("current align:", align)
		fuelCostTmp := 0
		for number, count := range numbers {
			diff := math.Abs(float64(align - number))
			fuelCost := (diff / 2) * (diff + 1)
			fuelCostTmp += count * int(fuelCost)
		}
		fmt.Println("fuel cost", fuelCostTmp, "at", align)
		if fuelCostTmp < fuelCostTotal {
			optimalAlign = align
			fuelCostTotal = fuelCostTmp
		}

	}

	fmt.Println("total fuel:", fuelCostTotal, "at", optimalAlign)

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
