package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("./Day1/input.txt")

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

	lines := make([]int, 0)
	sums := make([]int, 0)

	i := 0
	count := 0

	for scanner.Scan() {
		i, _ = strconv.Atoi(scanner.Text())
		lines = append(lines, i)
	}

	for j := 0; j < len(lines)-2; j++ {
		sums = append(sums, lines[j]+lines[j+1]+lines[j+2])
	}

	for j := 1; j < len(sums); j++ {
		if sums[j] > sums[j-1] {
			count++
		}
	}

	fmt.Println("Count: ", count)

}
