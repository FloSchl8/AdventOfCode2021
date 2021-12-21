package main

import (
	"bufio"
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

	i := 0
	j := 0
	count := 0

	for scanner.Scan() {
		i, _ = strconv.Atoi(scanner.Text())
		if i > j && j != 0 {
			count++
		}
		j = i
	}

	println("Count: ", count)

}
