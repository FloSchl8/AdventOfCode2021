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

	file, err := os.Open("./Day2/input.txt")

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

	x := 0
	y := 0
	z := 0

	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Fields(text)
		i, _ := strconv.Atoi(fields[1])
		switch fields[0] {
		case "up":
			z -= i
		case "down":
			z += i
		case "forward":
			y += i
			x += i * z
		default:
			log.Fatal("unknown direction")
		}
		fmt.Println("X", x, "Y", y)
	}

	fmt.Println("Position ", x*y)

}
