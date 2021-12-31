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
	file, err := os.Open("./Day6/input.txt")

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

	scanner.Scan()

	input := strings.Split(scanner.Text(), ",")

	fishMap := make(map[int]uint64)

	for _, s := range input {
		timer, _ := strconv.Atoi(s)
		fishMap[timer]++
	}

	fmt.Println(fishMap)

	for i := 0; i < 256; i++ {
		tmp := make(map[int]uint64)

		for j := 0; j < 9; j++ {
			if j == 0 {
				tmp[8] += fishMap[0]
				tmp[6] += fishMap[0]
				fishMap[0] = 0
			} else {
				tmp[j-1] += fishMap[j]
				fishMap[j] -= fishMap[j]
			}
		}

		for timer, count := range tmp {
			fishMap[timer] += count
		}
	}

	fmt.Println(fishMap)

	var sum uint64

	for _, count := range fishMap {
		sum += count
	}

	fmt.Println(sum)

}
