package main

import (
	"fmt"
	"math/rand"
)

var ghosts []*sprite

func captureGhostPosition(maze []string) {
	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'G':
				ghosts = append(ghosts, &sprite{row, col, row, col})
			}
		}
	}
}

func moveGhosts(maze []string) {
    for _, g := range ghosts {
		dir := randomDirection()
        g.row, g.col = makeMove(g.row, g.col, dir, maze)
    }
}

func randomDirection() string {//temp AI
    dir := rand.Intn(4)//pick random dir
    move := map[int]string{
        0: "UP",
        1: "DOWN",
        2: "RIGHT",
        3: "LEFT",
    }
    return move[dir]
}

func renderGhost(maze []string) {
	for _, g := range ghosts {
		moveCursor(g.row, g.col)
		fmt.Print(cfg.Ghost)
	}
}