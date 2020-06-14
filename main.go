//runnable programs must have a package called main and a main function
package main

//makes code in other packages accessible to this program
import (
	"fmt"
	//"os/exec"
	//"bufio" //buffered IO package will have the scanner object
	"log"
	//"os" //package to open outside files and release the file handler
	"time"

	"github.com/danicat/simpleansi"
)

func renderScreen(maze []string) { //array of whatever size
	simpleansi.ClearScreen()
	for _, line := range maze {
        for _, chr := range line {
            switch chr {
            case '#':
                fmt.Print(simpleansi.WithBlueBackground(cfg.Wall))
            case '.':
                fmt.Print(cfg.Dot)
            default:
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
	//PLAYER
	//renderPlayer(maze)

	//GHOST
	//renderGhost(maze)

	//PRINT SCORE
	renderGUI(maze, 1)

	//DEBUG?
	renderDebug(maze, 2)
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

	//LOAD GRAPHICS ASSETS
	err = loadConfig("config.json")
	if err != nil {
		log.Println("failed to load configuration:", err)
		return
	}

	maze := mazeloader.getmaze()

	//Record num of dots
	recordNumDots(maze);
	capturePlayerPosition(maze);
	captureGhostPosition(maze);

	/////////CHANNELS////////////////
		// process input
		inputChannel := make(chan string)
		go func(ch chan<- string) {
			for {
				input, err := readInput()
				if err != nil {
					log.Println("error reading input:", err)
					ch <- "ESC"
				}
				ch <- input
			}
		}(inputChannel)
	
	// game loop
	for {
		
		//////////////////////////LOGIC
		// process movement
		select {
		case input:= <-inputChannel:
			if input == "ESC" {
				lives = 0
			}
			debugLog(input)
			movePlayer(input,maze)
		default:
		}
		
		moveGhosts(maze)

		// process collisions
		checkForLiveLost(ghosts)

		// process score
		eatNumDots(maze,player)

		// check game over
		if checkGameOver() {
			break;
		}
		/////////////GRAPHICS/////////
		// update screen
		renderScreen(maze)

		// repeat
		time.Sleep(200 * time.Millisecond)
	}
}
