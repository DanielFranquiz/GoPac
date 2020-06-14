//runnable programs must have a package called main and a main function
package main

//makes code in other packages accessible to this program
import (
	"fmt"
)

var score int
var numDots int
var lives = 3

func recordNumDots (maze []string) {
	for _, line := range maze {
		for _, char := range line {
			switch char {
			case '.':
				numDots++
			}
		}
	}	
}

func eatNumDots(maze []string, player sprite) {
	switch maze[player.row][player.col] {
	case '.':
		numDots--
		score++
		//Remove dot from the maze
		maze[player.row] = maze[player.row][0:player.col] + " " + maze[player.row][player.col+1:]
	}
}

func renderGUI(maze []string, line int){
	moveCursor(len(maze)+line, 0)
	fmt.Println("Score: ", score, "\tLives:", lives)
}

func checkGameOver () bool {
	if numDots == 0 || lives <= 0 {
        return true;
		}
	return false;
}