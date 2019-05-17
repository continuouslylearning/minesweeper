package main

import (
	"fmt"
	"math/rand"
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

	ms.printGameBoard()
	ms.printMineBoard()

}

func main() {
	ms := MineSweeper{
		rows:     5,
		cols:     5,
		numMines: 4,
	}

	ms.Play()
}
