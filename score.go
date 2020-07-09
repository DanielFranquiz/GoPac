//runnable programs must have a package called main and a main function
package main

//makes code in other packages accessible to this program
import (
	"fmt"
	"strconv"
	"bytes"
	"time"
	"sync"
)

var score int
var numDots int
var lives = 3

var pillTimer *time.Timer
var pillMx sync.Mutex
var ghostsStatusMx sync.RWMutex

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
//should move this function into player
func eatNumDots(maze []string, player sprite) {

	//can separate as own function
	removeDot := func(row, col int) {
        maze[row] = maze[row][0:col] + " " + maze[row][col+1:]
    }

	switch maze[player.row][player.col] {
	case '.':
		numDots--
		score++ //<- make function for this
		//Remove dot from the maze
		removeDot(player.row, player.col)
	case 'X':
        score += 10//<- make function for this
		removeDot(player.row, player.col)
		go processPill()
	}
}

func processPill() {
	pillMx.Lock()
	updateGhosts(ghosts, GhostStatusBlue)
	if pillTimer != nil {
		pillTimer.Stop()
	}
	pillTimer = time.NewTimer(time.Second * cfg.PillDurationSecs)
	pillMx.Unlock()
	<-pillTimer.C
	pillMx.Lock()
	pillTimer.Stop()
	updateGhosts(ghosts, GhostStatusNormal)
	pillMx.Unlock()
}

func updateGhosts(ghosts []*ghost, ghostStatus GhostStatus) {
	ghostsStatusMx.Lock()
	defer ghostsStatusMx.Unlock()
	for _, g := range ghosts {
		g.status = ghostStatus
	}
}

func renderGUI(maze []string, line int){
	moveCursor(len(maze)+line, 0)

	livesRemaining := strconv.Itoa(lives) //converts lives int to a string
    if cfg.UseEmoji {
        livesRemaining = getLivesAsEmoji()
	}
	
	fmt.Println("Score: ", score, "\tLives:", livesRemaining)
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


//concatenate the correct number of player emojis based on lives
func getLivesAsEmoji() string{
    buf := bytes.Buffer{}
    for i := lives; i > 0; i-- {
        buf.WriteString(cfg.Player)
    }
    return buf.String()
}