// The Bouncing Ball Project
// 18-project-bouncing-ball
// 02-add-a-buffer
//
// Performance results for
//
// 		Non-Buffered Version
// ⇒  time go run main.go > /dev/null
// go run main.go > /dev/null  0.65s user 0.43s system 104% cpu 1.035 total
//
// 		Buffered, but multiple backing arrays
// ⇒  time go run main.go > /dev/null
// go run main.go > /dev/null  5.19s user 0.54s system 115% cpu 4.975 total
//
// 		Buffered
//⇒  time go run main.go > /dev/null
// go run main.go > /dev/null  0.34s user 0.27s system 110% cpu 0.556 total

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
	}

	// create a rune for handling and printing a default character
	var mark rune

	// here, we define a buffer to hold the output of the loops
	// into a rune and then print the entire buffer at once there
	// by avoiding the invocation of expensive print function
	// calls during each loop
	buffer := make([]rune, 0, boardWidth*boardHeight)

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

	// debugging...
	// _ = mark
	fmt.Println("width from board:", len(board))
	fmt.Println("height from board[0]:", len(board[0]))
	fmt.Println()

	// introducing the for loop for measuring the
	// performance difference between, buffered and non
	// buffered variants
	// _ = buffer
	for i := 0; i < 1000; i++ {
		// the current version of buffer provided allocates
		// a new backing array for each append there by
		// causing performace degradation.
		// so, reduce it's length to use the same buffer
		// for each append to start from first as below
		buffer = buffer[:0]

		// printing the board with truth values
		for y := range board[0] {
			for x := range board {
				mark = no
				if board[x][y] {
					mark = yes
				}
				buffer = append(buffer, mark, ' ')
				// fmt.Printf("%2s", string(mark))
			}
			buffer = append(buffer, '\n')
			// fmt.Println()
		}
		// now print the buffer
		fmt.Printf("%s\n", string(buffer))
	}
}
