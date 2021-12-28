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
	list := make([][]string, 0)

	// map: count 0 and 1 [0: 123] [1: 234]

	counter := 0

	for scanner.Scan() {
		s := scanner.Text()
		counter++
		split := strings.Split(s, "")
		list = append(list, split)
		// split string after every sign
		getByteCount(split, bytes)
	}

	o2filter := "0"
	co2filter := "0"
	o2list := list
	co2list := list

	for i := 0; i < 12; i++ {

		o2filter = getListFilter(o2list, i, true)
		co2filter = getListFilter(co2list, i, false)

		if len(o2list) > 1 {
			o2list = filterStringListAtPosition(o2list, i, o2filter)
		}
		if len(co2list) > 1 {
			co2list = filterStringListAtPosition(co2list, i, co2filter)
		}
	}

	//fmt.Println(o2list, co2list)

	gamma := ""
	epsilon := ""
	for _, zeroones := range bytes {
		zeros := zeroones[0]
		if zeros > 500 {
			gamma = buildBinaryString("0", gamma)
			epsilon = buildBinaryString("1", epsilon)
		} else {
			gamma = buildBinaryString("1", gamma)
			epsilon = buildBinaryString("0", epsilon)
		}
	}

	o2 := ""
	co2 := ""

	for i := 0; i < 12; i++ {
		o2 = buildBinaryString(o2list[0][i], o2)
		co2 = buildBinaryString(co2list[0][i], co2)
	}

	//fmt.Println("gamma", gamma, "epsilon", epsilon)
	//fmt.Println("o2", o2, "co2", co2)

	gammaVal, _ := strconv.ParseInt(gamma, 2, 32)
	epsilonVal, _ := strconv.ParseInt(epsilon, 2, 32)
	o2Val, _ := strconv.ParseInt(o2, 2, 32)
	co2Val, _ := strconv.ParseInt(co2, 2, 32)

	fmt.Println("gamma", gammaVal, "epsilon", epsilonVal, "gamma*epsilon", gammaVal*epsilonVal)
	fmt.Println("o2", o2Val, "co2", co2Val, "o2*co2", o2Val*co2Val)
}

// binary string builder
func buildBinaryString(c interface{}, s string) string {
	switch v := c.(type) {
	case string:
		s += v
	case int:
		s += strconv.Itoa(v)
	}

	return s
}

// filter list for most or least common 0 or 1 at position
func getListFilter(list [][]string, position int, most bool) string {

	filter := ""
	zeros := 0
	ones := 0

	for _, chars := range list {
		if chars[position] == "1" {
			ones++
		} else {
			zeros++
		}
	}

	if zeros > ones && most {
		filter = "0"
	} else if ones >= zeros && most {
		filter = "1"
	} else if zeros <= ones && !most {
		filter = "0"
	} else if ones < zeros && !most {
		filter = "1"
	} else {
		log.Fatal("something wrong")
	}

	return filter
}

func filterStringListAtPosition(list [][]string, position int, char string) [][]string {

	newList := make([][]string, 0)

	for _, s := range list {
		if s[position] == char {
			newList = append(newList, s)
		}
	}

	return newList
}

func getByteCount(s []string, bytes []map[int]int) {
	for i, num := range s {
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
