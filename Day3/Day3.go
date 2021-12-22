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
	file, err := os.Open("./Day3/input.txt")

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

	// slice mit maps
	bytes := make([]map[int]int, 12)

	// map: anzahl 0er und 1er [0: 123] [1: 234]

	counter := 0

	for scanner.Scan() {
		s := scanner.Text()
		counter++
		// string nach jedem zeichen splitten
		for i, num := range strings.Split(s, "") {
			bytemap := bytes[i]
			if bytemap == nil {
				bytemap = make(map[int]int)
			}
			zeroone, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			// 0 oder 1 hoch zählen
			bytemap[zeroone]++
			bytes[i] = bytemap
		}
	}

	gamma := ""
	epsilon := ""
	for _, zeroones := range bytes {
		zeros := zeroones[0]
		if zeros > 500 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	fmt.Println(bytes, counter)
	fmt.Println(gamma, epsilon)

	gammaVal, _ := strconv.ParseInt(gamma, 2, 32)
	epsilonVal, _ := strconv.ParseInt(epsilon, 2, 32)

	fmt.Println(gammaVal, epsilonVal, gammaVal*epsilonVal)
}
