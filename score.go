//runnable programs must have a package called main and a main function
package main

//makes code in other packages accessible to this program
import (
	"fmt"
)

var score int
var numDots int
var lives = 1000

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

	removeDot := func(row, col int) {
        maze[row] = maze[row][0:col] + " " + maze[row][col+1:]
    }

	switch maze[player.row][player.col] {
	case '.':
		numDots--
		score++
		//Remove dot from the maze
		removeDot(player.row, player.col)
	case 'X':
        score += 10
        removeDot(player.row, player.col)
	}
}


func renderGUI(maze []string, line int){
	moveCursor(len(maze)+line, 0)
	fmt.Println("Score: ", score, "\tLives:", lives)
}

func checkGameOver () (bool,bool) {
	var hasWon= false;
	var Gameover = false;

	if numDots == 0 || lives <= 0 {
        Gameover = true;
		}
	
	if numDots == 0{
		hasWon = true;
	}

	return Gameover,hasWon;
}