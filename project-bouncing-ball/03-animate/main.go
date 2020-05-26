// The Bouncing Ball Project
// 18-project-bouncing-ball
//   - 03-animate
//

package main

import (
	"fmt"
	"time"

	// screen package
	"github.com/inancgumus/screen"
)

// Constans declaration for controlling
const (
	boardWidth  = 60 // number of columns
	boardHeight = 20 // numer of rows

	// Whether to print a ball or not
	yes = '🥎'
	no  = ' '

	// The frames variable, controls the running of animation for upto 60 secs
	// Frames Per Second: 20 fps
	// Total for 20 fps = 20 * 60 = 1200 times
	frames = 1200

	// Controlling the speed of the loop with interval
	speed = time.Second / 20
)

func main() {

	var (
		// For animating the  ball, it is necessary to track  the position of the
		// ball w.r.t to both x and y axes.  For that, declare the same as px and
		// py respectively
		px, py int

		// Define the velocities (not speed) of the  in the x & y directions. The
		// velocity will have both magnitude  and direction helping us to control
		// the orientation of  the ball, the values of +1  indicate ball moves to
		// right and down at 1 step a time
		vx, vy = 1, 1

		// Create a rune for handling and printing the ball character
		mark rune
	)

	// Create a  matrix with form  and size = width  X height (rows  X cols);
	// first create base multidimensional structure
	//
	// 		 ------ w ------
	// 		|__|__|__|__|__|  |
	// 		|__|__|__|__|__|  |
	// 		|__|__|__|__|__|  |
	// 		|__|__|__|__|__|  h
	// 		|__|__|__|__|__|  |
	// 		|__|__|__|__|__|  |
	// 		|__|__|__|__|__|  |
	//
	board := make([][]bool, boardWidth)


	// Now, populate  each col of  the matrix spanning across  the boardWidth
	// (no. of columns) and boardHeight (no. of rows) with a default value of
	// boolean (false)
	for col := range board {
		board[col] = make([]bool, boardHeight)
	}

	// Define a buffer to  hold the output of the loops into  a rune and then
	// print the entire  buffer at once, there by avoiding  the invocation of
	// expensive print function calls during each loop.
	buffer := make([]rune, 0, boardWidth*boardHeight)


	// clear the screen to start on a fresh canvas during the run
	screen.Clear()

	// Using a loop  to calculate and update the next  ball position and draw
	// to the screen for the animation frames period
	for i := 0; i < frames; i++ {

		// Update the positions of the ball for each loop using the velocities in
		// the x and  y directions. The velocities will be  positive to the right
		// and  to down;  everything else  should be  negative in  the respective
		// directions. We will use the standard physics formula s = v x t formula
		// to calculate the distance travelled in unit time as s = v x 1 = v
		px += vx
		py += vy

		// If the ball hits a border, it  should recoil back from the border. So,
		// the boundaries of  the ball area should be checked  first The velocity
		// will be reversed in such cases, where the ball is bounced.
		//
		// conditions to check if ball has hit a border
		// x-axis: px <= 0 or px >= boardWidth - 1
		if px <= 0 || px >= boardWidth-1 {
			vx *= -1
		}
		// now carry same logic to y direction
		// y-axis: py <= 0 or py >= boardHeight - 1
		if py <= 0 || py >= boardHeight-1 {
			vy *= -1
		}

		// now set the initial position of the ball
		// board[col][row] = board[width][height]
		board[px][py] = true


		// The default  version of buffer  defined allocates a new  backing array
		// for each append operation there by causing performace degradation. So,
		// reduce it's  length to use  the same buffer  for each append  to start
		// from first as below
		buffer = buffer[:0]

		// now print the board with truth values
		for y := range board[0] {
			for x := range board {
				mark = no
				if board[x][y] {
					mark = yes
				}
				buffer = append(buffer, mark, ' ')
				// clear the previous position once it's covered
				board[x][y] = false
			}
			// fmt.Println()
			buffer = append(buffer, '\n')
		}

		// Now print  the buffer making sure  to start drawing the  ball from the
		// same position after every iteration
		screen.MoveTopLeft()
		fmt.Printf("%s\n", string(buffer))

		// sleep for a period
		time.Sleep(speed)
	}
}
