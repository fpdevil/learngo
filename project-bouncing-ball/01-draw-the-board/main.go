// The Bouncing Ball Project
// 18-project-bouncing-ball
// 01-draw-the-board
package main

import (
	"fmt"
)

// This is the first part of the Bouncing Ball Project
// dealing with the ground work preparation, specifically
// the design and layout of the Board

const (
	boardWidth  = 50 // number of columns
	boardHeight = 10 // numer of rows

	yes = '🥎'
	no  = ' '
)

func main() {
	// create matrix form and size = width X height (rows X cols)
	// first create base multidimensional structure
	board := make([][]bool, boardWidth)
	// fmt.Printf("%#v\n", board)

	// 		 ------ w ------
	// 		|__|__|__|__|__|  |
	// 		|__|__|__|__|__|  |
	// 		|__|__|__|__|__|  |
	// 		|__|__|__|__|__|  h
	// 		|__|__|__|__|__|  |
	// 		|__|__|__|__|__|  |
	// 		|__|__|__|__|__|  |
	//
	// now, we populate each col of the matrix spanning across the
	// boardWidth (no. of columns) and boardHeight (no. of rows)
	// with a default value of boolean (false)
	for col := range board {
		board[col] = make([]bool, boardHeight)
		// fmt.Printf("%#v\n", board[col])
	}

	// create a rune for handling and printing a default character
	var mark rune

	// now populate some default true values to be marked across
	// the board by assiging a character to mark run
	// board[col][row] = board[width][height]

	board[0][0] = true
	board[1][0] = true
	board[2][0] = true
	board[3][0] = true

	board[6][1] = true
	board[7][2] = true
	board[6][3] = true

	board[0][4] = true
	board[1][4] = true
	board[2][4] = true
	board[3][4] = true

	board[15][5] = true
	board[12][3] = true
	board[15][1] = true
	board[12][7] = true
	board[15][9] = true
	board[11][5] = true
	board[19][3] = true
	board[19][5] = true
	board[19][7] = true

	// board[0][9] = true
	// board[1][9] = true
	// board[1][8] = true
	// board[2][7] = true
	// board[3][7] = true
	// board[4][6] = true
	// board[5][6] = true
	// board[6][5] = true
	// board[7][5] = true
	// board[8][4] = true
	// board[9][4] = true
	// board[10][3] = true
	// board[11][3] = true
	// board[12][2] = true
	// board[13][1] = true
	// board[14][0] = true
	// board[15][0] = true

	// fmt.Printf("board: %#v\n", board)
	// debugging...
	// _ = mark
	fmt.Println("width from board:", len(board))
	fmt.Println("height from board[0]:", len(board[0]))
	fmt.Println()
	// for range board[0] {
	// 	for range board {
	// 		fmt.Printf("%2s", "x")
	// 	}
	// 	fmt.Println()
	// }

	// clearing the screen and moving cursor to top left
	// screen.Clear()
	// screen.MoveTopLeft()

	// printing the board with truth values
	for y := range board[0] {
		for x := range board {
			mark = no
			if board[x][y] {
				mark = yes
			}
			fmt.Printf("%2s", string(mark))
			// fmt.Print(string(mark), " ")
		}
		fmt.Println()
	}

}
