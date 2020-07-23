package main

import (
)

type sprite struct {
	row int
    col int
    startRow int
    startCol int
}

func makeMove(oldRow, oldCol int, dir string, maze []string) (newRow, newCol int) {
    newRow, newCol = oldRow, oldCol

    switch dir {
    case "UP":
        newRow = newRow - 1
        if newRow < 0 {
            newRow = len(maze) - 1
        }
    case "DOWN":
        newRow = newRow + 1
        if newRow == len(maze) {
            newRow = 0
        }
    case "RIGHT":
        newCol = newCol + 1
        if newCol == len(maze[0]) {
            newCol = 0
        }
    case "LEFT":
        newCol = newCol - 1
        if newCol < 0 {
            newCol = len(maze[0]) - 1
        }
    }

    
    if isWall(maze,newRow,newCol) {
        newRow = oldRow
        newCol = oldCol
    }

    return
}

func GetDirectionToVector(dir string) (int, int) {

    var x, y int = 0,0

    switch dir {
    case "UP":
        y = -1
    case "DOWN":
        y =  1
    case "RIGHT":
        x =   1
    case "LEFT":
        x = -1
    }


    return x,y
}

///////////UTIL///////////////////////////////////////////
func isWall(maze []string, row int , col int) (bool){
    var flag = false;

	if maze[row][col] == '#' {
        flag = true
	}
	return flag
}