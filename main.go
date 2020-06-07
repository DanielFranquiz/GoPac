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

func printScreen(maze []string) { //array of whatever size
	simpleansi.ClearScreen()
	for _, line := range maze { //_ is a placeholder for where the compiler would expect a variable name. we are ignoring that value
		for _, chr := range line {
            switch chr {
            case '#':
                fmt.Printf("%c", chr)
            default:
                fmt.Print(" ")
            }
        }
		fmt.Println()
	}

	simpleansi.MoveCursor(player.row, player.col)
    fmt.Print("P")

    // Move cursor outside of maze drawing area
    simpleansi.MoveCursor(len(maze)+1, 0)
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

	CapturePlayerPosition(mazeloader.getmaze());

	// game loop
	for {
		// update screen
		maze := mazeloader.getmaze()
		printScreen(maze)

		// process input
		input, err := readInput()
		if err != nil {
			log.Print("error reading input:", err)
			break
		}

		// process movement
		movePlayer(input,maze)
		// process collisions

		// check game over
		if input == "ESC" {
			break
		}

		// Temp: break infinite loop
		//fmt package formatted input/output
		//fmt.Println("Hello, Pac Go!")
		//break

		// repeat
	}
}
