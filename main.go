//runnable programs must have a package called main and a main function
package main

//makes code in other packages accessible to this program
import (
	"fmt"
	//"os/exec"
	//"bufio" //buffered IO package will have the scanner object
	"log"
	//"os" //package to open outside files and release the file handler

	"github.com/danicat/simpleansi"
)

var score int
var numDots int
var lives = 3

func renderScreen(maze []string) { //array of whatever size
	simpleansi.ClearScreen()
	for _, line := range maze { //_ is a placeholder for where the compiler would expect a variable name. we are ignoring that value
		for _, chr := range line {
            switch chr {
            case '#':
                fallthrough
            case '.':
                fmt.Printf("%c", chr)
            default:
                fmt.Print(" ")
            }
        }
		fmt.Println()
	}
	//PLAYER
	renderPlayer(maze)

	//GHOST
	renderGhost(maze)

	//PRINT SCORE
	simpleansi.MoveCursor(len(maze)+1, 0)
	fmt.Println("Score: ", score, "\tLives:", lives)
}

//entry point of program defined as function with func
//main program takes no parameters and returns nothing
func main() {
	// initialize game
	// initialise game
	initialise()
	defer cleanup()

	mazeloader := MazeLoader{}
	// load resources
	err := mazeloader.loadMaze("maze01.txt")
	if err != nil {
		fmt.Println("failed to load maze:", err)
		return
	}


	maze := mazeloader.getmaze()


	//Record num of dots
	for _, line := range maze {
		for _, char := range line {
			switch char {
			case '.':
				numDots++
			}
		}
	}

	capturePlayerPosition(maze);
	
	captureGhostPosition(maze);

	// game loop
	for {
		// update screen
		renderScreen(maze)

		// process input
		input, err := readInput()
		if err != nil {
			log.Print("error reading input:", err)
			break
		}

		// process movement
		movePlayer(input,maze)
		moveGhosts(maze)

		// process collisions
		for _, g := range ghosts {
			if player == *g {
				lives--
			}
		}

		// check game over
		if input == "ESC" || numDots == 0 || lives <= 0 {
        break
    	}

		// Temp: break infinite loop
		//fmt package formatted input/output
		//fmt.Println("Hello, Pac Go!")
		//break

		// repeat
	}

	// Rendering the last frame
	renderScreen(maze)
}
