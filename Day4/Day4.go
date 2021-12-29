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
	file, err := os.Open("./Day4/input.txt")

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

	// first line -> numbers
	scanner.Scan()
	drawnNumbers := getDrawnNumbers(scanner.Text())

	//fmt.Println("numbers drawn", drawnNumbers)

	// slice (boards) of slice (rows) of slices (columns)
	boards := make([][][]int, 0)

	// cols and rows -> [][]string
	collected := make([][]string, 0)

	for scanner.Scan() {
		if scanner.Text() != "" {
			row := strings.Split(scanner.Text(), " ")
			collected = append(collected, row)
		} else if len(collected) > 0 {
			boards = append(boards, convertBoardStringToInt(collected))
			collected = make([][]string, 0)
		}
	}
	boards = append(boards, convertBoardStringToInt(collected))

	winner := false
	winningBoards := make([]int, 0)

	for i := 4; i < len(drawnNumbers); i++ {
		currentDraw := drawnNumbers[:i+1]
		for j, board := range boards {
			winner = checkBoard(board, currentDraw)

			if winner && !wasAlreadyWinner(j, winningBoards) {
				winningBoards = append(winningBoards, j)
				fmt.Println("Winner board number", j)
				fmt.Println("drawn numbers", currentDraw)
				fmt.Println("Winning sum", getWinningBoardSum(boards[j], currentDraw))
			}
		}
	}

	fmt.Println(winningBoards)
	//fmt.Println("Winning sum", getWinningBoardSum(boards[winnerBoard], winningDraw))
}

func wasAlreadyWinner(boardnumber int, boards []int) bool {
	for i := 0; i < len(boards); i++ {
		if boards[i] == boardnumber {
			return true
		}
	}
	return false
}

func getWinningBoardSum(board [][]int, draw []int) int {
	result := 0
	for _, row := range board {
		for _, number := range row {
			shouldNotAdd := false
			for _, i := range draw {
				if i == number {
					shouldNotAdd = true
				}
			}
			if !shouldNotAdd {
				result += number
			}
		}
	}
	fmt.Println(result, draw[len(draw)-1])

	return result * draw[len(draw)-1]
}

func checkBoard(board [][]int, drawnNumbers []int) bool {

	for i := 0; i < len(board[0]); i++ {
		column := make([]int, 0)

		for _, row := range board {
			column = append(column, row[i])

			if checkBingo(row, drawnNumbers) {
				//fmt.Println("WINNER!! Board", board)
				return true
			}
		}

		if checkBingo(column, drawnNumbers) {
			//fmt.Println("WINNER!! Board", board)
			return true
		}
	}
	return false
}

func checkBingo(boardNumbers []int, drawnNumbers []int) bool {

	winner := false
	neededHits := len(boardNumbers)
	hits := 0

	for _, number := range boardNumbers {

		for _, drawnNumber := range drawnNumbers {
			if number == drawnNumber {
				hits++
			}
		}

		if hits == neededHits {
			//fmt.Println("WINNER ROW", boardNumbers)
			winner = true
			break
		}
	}

	return winner
}

func getDrawnNumbers(s string) []int {
	split := strings.Split(s, ",")
	numbers := make([]int, 0)

	for _, s := range split {
		number, _ := strconv.Atoi(s)
		numbers = append(numbers, number)
	}

	return numbers
}

func convertBoardStringToInt(stringBoard [][]string) [][]int {

	intBoard := make([][]int, 0)

	for _, row := range stringBoard {
		column := make([]int, 0)
		for _, col := range row {
			if col != "" {
				number, _ := strconv.Atoi(col)
				column = append(column, number)
			}
		}
		intBoard = append(intBoard, column)
	}

	return intBoard
}
