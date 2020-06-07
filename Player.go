package main

import (
	"fmt"
	"github.com/danicat/simpleansi"
)

var player sprite

func CapturePlayerPosition(maze []string) {
	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = sprite{row, col}
			}
		}
	}
}

func movePlayer(dir string,maze []string) {
    player.row, player.col = makeMove(player.row, player.col, dir, maze)
}

func renderPlayer(maze []string) {
	////PLAYER
	simpleansi.MoveCursor(player.row, player.col)
	fmt.Print("P")

	// Move cursor outside of maze drawing area
	simpleansi.MoveCursor(len(maze)+1, 0)
}