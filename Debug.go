package main

import (
	"fmt"
	"github.com/danicat/simpleansi"
)


var debug string

func debugLog(s string){
	debug = s;
}


func renderDebug(maze []string, line int){
	simpleansi.MoveCursor(len(maze)+line, 0)
	fmt.Println("DEBUG: ", debug)
}

