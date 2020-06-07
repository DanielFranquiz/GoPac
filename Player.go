package main

import (
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