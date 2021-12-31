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

	fishs := make([]fish, 0)

	for _, s := range input {
		timer, _ := strconv.Atoi(s)
		f := fish{timer: timer}
		fishs = append(fishs, f)
	}

	fmt.Println(fishs)

	for i := 0; i < 80; i++ {
		for i, f := range fishs {
			switch f.timer {
			case 0:
				fishs[i] = fish{timer: 6}
				fishs = append(fishs, fish{timer: 8})
			default:
				fishs[i] = fish{timer: f.timer - 1}
			}
		}
	}

	fmt.Println(len(fishs))

}

type fish struct {
	timer int
}
