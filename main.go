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

	"flag"
)

var (
	configFile = flag.String("config-file", "config.json", "path to custom configuration file")
	mazeFile   = flag.String("maze-file", "maze01.txt", "path to a custom maze file")
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
			case 'X':
				fmt.Print(cfg.Pill)
            default:
                fmt.Print(cfg.Space)
            }
        }
        fmt.Println()
	}
	
	//GHOST
	renderGhost(maze)

	//PLAYER
	renderPlayer(maze)

	//PRINT SCORE
	renderGUI(maze, 1)

	//DEBUG?
	renderDebug(maze, 2)
}

//entry point of program defined as function with func
//main program takes no parameters and returns nothing
func main() {
	flag.Parse()

	// initialise game
	initialise()
	defer cleanup()

	mazeloader := MazeLoader{}
	// load resources
	err := mazeloader.loadMaze(*mazeFile)
	if err != nil {
		fmt.Println("failed to load maze:", err)
		return
	}

	//LOAD GRAPHICS ASSETS
	err = loadConfig(*configFile)
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
		checkForCollision(ghosts,maze)

		// process score
		eatNumDots(maze,player)



		/////////////GRAPHICS/////////
		// update screen
		renderScreen(maze)

		// check game over
		var Gameover,HasWon = checkGameOver();

		if Gameover {
			if HasWon {

			} else {
				
			}
			break;
		}
		// repeat
		time.Sleep(200 * time.Millisecond)
	}
}
