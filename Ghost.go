package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var ghosts []*ghost
//var ghosts []*sprite

type GhostStatus string

const (
	GhostStatusNormal GhostStatus = "Normal"
	GhostStatusBlue   GhostStatus = "Blue"
)

type ghost struct {
	position sprite
	status   GhostStatus
	dir string
}

func captureGhostPosition(maze []string) {
	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'G':
				ghosts = append(ghosts, &ghost{sprite{row, col, row, col}, GhostStatusNormal,"UP"})
				//ghosts = append(ghosts, &sprite{row, col, row, col})
			}
		}
	}
}

func moveGhosts(maze []string) {
    for _, g := range ghosts {
		//check inter
			//if inter
				//set direction

		//g.dir = randomDirection()
		var x, y int = GetDirectionToVector(g.dir)

		var wall = isWall(maze,g.position.row + y, g.position.col + x);

		debugLog(strconv.FormatBool(wall))

		if wall { // if we are at a wall
			if g.dir == "UP" || g.dir == "DOWN"  { // if we are looking up or down
				var xa, ya int = GetDirectionToVector("RIGHT")
				if  !isWall(maze,g.position.row + ya, g.position.col + xa){ // decide left or right
					g.dir = "RIGHT"
				} else {
					g.dir = "LEFT"
				}
			} else if g.dir == "RIGHT" || g.dir == "LEFT" { // if we are looking left or right
				var xa, ya int = GetDirectionToVector("UP")
				if  !isWall(maze,g.position.row + ya, g.position.col + xa){ // decide up or down
					g.dir = "UP"
				} else {
					g.dir = "DOWN"
				}
			}
		}
			

		g.position.row, g.position.col = makeMove(g.position.row, g.position.col, g.dir, maze)
		//g.row, g.col = makeMove(g.row, g.col, dir, maze)
    }
}

func ghostReturnHome(g *ghost){
	g.position.row, g.position.col = g.position.startRow, g.position.startCol
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
		moveCursor(g.position.row, g.position.col)
		if g.status == GhostStatusNormal {
			fmt.Printf(cfg.Ghost)
		} else if g.status == GhostStatusBlue {
			fmt.Printf(cfg.GhostBlue)
		}
	}
}
