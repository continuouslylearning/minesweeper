package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type MineSweeper struct {
	mineBoard [][]byte
	gameBoard [][]byte
	rows      int
	cols      int
	numMines  int
	movesLeft int
}

func PrintBoard(board [][]byte) {

	fmt.Printf("  ")

	for i := 0; i < len(board[0]); i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Printf("\n")

	for i := 0; i < len(board); i++ {

		fmt.Printf("%d ", i)

		for j := 0; j < len(board[0]); j++ {
			fmt.Printf("%c ", board[i][j])
		}

		fmt.Printf("\n")
	}
}

func (ms *MineSweeper) printGameBoard() {
	PrintBoard(ms.gameBoard)
}

func (ms *MineSweeper) printMineBoard() {
	PrintBoard(ms.mineBoard)
}

func (ms *MineSweeper) setMines() {
	// use Seed method on the rand struct to make sure the generator doesn't produce the same sequence of integers in each game
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < ms.numMines; {
		x := rand.Intn(ms.rows)
		y := rand.Intn(ms.cols)

		// a mine has already been set at this location, so continue to next iteration of loop
		if ms.mineBoard[x][y] == '*' {
			continue
		}

		// successfully set a mine at this location, so increment `i`
		ms.mineBoard[x][y] = '*'
		i++
	}
}

func (ms *MineSweeper) createBoards() {
	ms.mineBoard = make([][]byte, ms.rows)
	ms.gameBoard = make([][]byte, ms.rows)

	for i := 0; i < ms.rows; i++ {
		ms.mineBoard[i] = make([]byte, ms.cols)
		ms.gameBoard[i] = make([]byte, ms.cols)
	}

	for i := 0; i < ms.rows; i++ {
		for j := 0; j < ms.cols; j++ {
			ms.gameBoard[i][j] = '-'
			ms.mineBoard[i][j] = '-'
		}
	}
}

func (ms *MineSweeper) isValid(x, y int) bool {
	if x < 0 || x >= ms.rows {
		return false
	}

	if y < 0 || y >= ms.cols {
		return false
	}

	return true
}

func (ms *MineSweeper) countMines(x, y int) int {

	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if ms.isValid(x+i, y+j) && ms.mineBoard[x+i][y+j] == '*' {
				count++
			}
		}
	}
	return count
}

func (ms *MineSweeper) visit(x, y int) bool {

	// this location has already been visited
	if ms.gameBoard[x][y] != '-' {
		return false
	}

	// a mine is at this location
	if ms.mineBoard[x][y] == '*' {
		return true
	}

	ms.movesLeft--
	count := ms.countMines(x, y)
	ms.gameBoard[x][y] = byte(count) + '0'
	// if the location has no adjacent mines, recursively visit the adjacent locations
	if count == 0 {

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if ms.isValid(x+i, y+j) {
					ms.visit(x+i, y+j)
				}
			}
		}
	}

	return false
}

func (ms *MineSweeper) makeChoice() (int, int) {

	r := bufio.NewReader(os.Stdin)
	fmt.Println("Make a choice: ")
	input, _, _ := r.ReadLine()
	location := strings.Split(string(input), ",")
	x, _ := strconv.Atoi(location[0])
	y, _ := strconv.Atoi(location[1])

	return x, y
}

/*
	call Play method on instance of MineSweeper to start a game
*/
func (ms *MineSweeper) Play() {

	/*
		first, initialize the board properties on the instance
		then, randomly set mines on the mine board
	*/
	ms.createBoards()
	ms.setMines()

	// set initialize value of `movesLeft` property on Minesweeper instance
	ms.movesLeft = ms.rows*ms.cols - ms.numMines
	ms.printMineBoard()
	gameOver := false

	for !gameOver {

		ms.printGameBoard()
		x, y := ms.makeChoice()

		gameOver = ms.visit(x, y)

		if gameOver {
			fmt.Println("You hit a mine. Game over!")
			ms.printMineBoard()
		}

		if !gameOver && ms.movesLeft == 0 {
			ms.printGameBoard()
			fmt.Printf("\nCongratulations, you won the game!")
			gameOver = true
		}
	}
}

func main() {
	ms := MineSweeper{
		rows:     5,
		cols:     5,
		numMines: 4,
	}

	ms.Play()
}
