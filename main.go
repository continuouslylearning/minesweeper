package main

type MineSweeper struct {
	mineBoard [][]byte
	gameBoard [][]byte
	rows      int
	cols      int
	numMines  int
	movesLeft int
}

func main() {
	ms := MineSweeper{
		rows:     5,
		cols:     5,
		numMines: 4,
	}
}
