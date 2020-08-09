package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"math"
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
		
		var x_diff,y_diff int = getVectorDifferenceFromPacManToGhost(g);

		/// B - A
		// A == ghost
		// B == pac

		
		xa := int(math.Abs(float64(x_diff)));
		ya := int(math.Abs(float64(y_diff)));

		debugLog(strconv.Itoa(x_diff) + "," + strconv.Itoa(y_diff) +","+ g.dir);

		if(isAtIntersection(maze,g)){
			

			if(xa > ya){// x - axis
				if(x_diff < 0){//we are left of him
					g.dir = "LEFT";
				}else{//we are right of him
					g.dir = "RIGHT";
				}
				debugLog(strconv.Itoa(x_diff) + "," + strconv.Itoa(y_diff) +","+ g.dir + "," + "x");
			}else{// y - axis
				if(y_diff < 0){//below
					g.dir = "UP";
				}else{//above
					g.dir = "DOWN";
				}
				debugLog(strconv.Itoa(x_diff) + "," + strconv.Itoa(y_diff) +","+ g.dir + "," + "y");
			}
			keepMovement(maze,g);// fixes problem a bit
		}else{
			keepMovement(maze,g);
		}
		
		

		g.position.row, g.position.col = makeMove(g.position.row, g.position.col, g.dir, maze)
		//g.row, g.col = makeMove(g.row, g.col, dir, maze)
    }
}

func getVectorDifferenceFromPacManToGhost(g *ghost)(int,int){
	var pacman sprite = player;

	var x_diff,y_diff int = 0,0;

	/// B - A
	// A == ghost
	// B == pac

	x_diff = pacman.col - g.position.col;
	y_diff = pacman.row - g.position.row;

	return x_diff,y_diff;
}

func isAtIntersection(maze []string,g *ghost)(bool){

		var count int = 0;

		var x, y int = GetDirectionToVector("UP")
		var wall = isWall(maze,g.position.row + y, g.position.col + x);

		if(!wall){
			count++;
		}

		x, y = GetDirectionToVector("DOWN")
		wall = isWall(maze,g.position.row + y, g.position.col + x);

		if(!wall){
			count++;
		}
		x, y = GetDirectionToVector("RIGHT")
		wall = isWall(maze,g.position.row + y, g.position.col + x);

		if(!wall){
			count++;
		}

		x, y = GetDirectionToVector("LEFT")
		wall = isWall(maze,g.position.row + y, g.position.col + x);

		if(!wall){
			count++;
		}

		if(count > 2){
			return true;
		}

		return false;
}

func keepMovement(maze []string,g *ghost){
	var x, y int = GetDirectionToVector(g.dir)

	var wall = isWall(maze,g.position.row + y, g.position.col + x);

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
