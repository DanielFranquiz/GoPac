package main

import (
	"fmt"
	"math/rand"
	"github.com/danicat/simpleansi"
)

var ghosts []*sprite

func CaptureGhostPosition(maze []string) {
	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'G':
				ghosts = append(ghosts, &sprite{row, col})
			}
		}
	}
}

func moveGhosts(maze []string) {
    for _, g := range ghosts {
		dir := drawDirection()
        g.row, g.col = makeMove(g.row, g.col, dir, maze)
    }
}

func drawDirection() string {//temp AI
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
		simpleansi.MoveCursor(g.row, g.col)
		fmt.Print("G")
	}
}