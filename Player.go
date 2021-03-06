package main

import (
	"fmt"
	"time"
)

var player sprite

func capturePlayerPosition(maze []string) {
	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = sprite{row, col, row, col}
			}
		}
	}
}

func movePlayer(dir string,maze []string) {
    player.row, player.col = makeMove(player.row, player.col, dir, maze)
}

func checkForCollision(ghosts []*ghost, maze []string) {
	for _, g := range ghosts {
		if player.row == g.position.row && player.col == g.position.col {

			//TODO Make Ghost die. Player Lives if touchs blue ghost. 
			if g.status == GhostStatusBlue{
				ghostReturnHome(g)

			}else{
				lives--

				if lives != 0 {
					playerHit(maze)
					
					for _, g := range ghosts { // reset Ghosts Position
						ghostReturnHome(g)
					}
           		 }
			}

			break
		}
	}
}

func playerHit(maze []string) {
	playerDeathAnimation(maze)
	time.Sleep(1000 * time.Millisecond) //dramatic pause before resetting player position
	player.row, player.col = player.startRow, player.startCol
}

func playerDeathAnimation(maze []string) {
	moveCursor(player.row, player.col)
    fmt.Print(cfg.Death)
    moveCursor(len(maze)+2, 0)
}

func renderPlayer(maze []string) {
	////PLAYER
	moveCursor(player.row, player.col)
	fmt.Print(cfg.Player)

	// Move cursor outside of maze drawing area
	moveCursor(len(maze)+1, 0)
}