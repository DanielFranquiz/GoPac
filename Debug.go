package main

import (
	"fmt"
)


var debug string

func debugLog(s string){
	debug = s;
}


func renderDebug(maze []string, line int){
	moveCursor(len(maze)+line, 0)
	fmt.Println("DEBUG: ", debug)
}

