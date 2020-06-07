package main

import (
	"bufio"
	"fmt"
	"os"
)

type MazeLoader struct {
	maze []string
}

func (mazeloader *MazeLoader) loadMaze(file string) error {
	fmt.Println("LOADING MAZE")
	//:= short-assignment operator that automatically infers the type of variable(s) based on the value(s) on the right hand side. commonly used in if, for, switch
	f, err := os.Open(file) //Open function from os package which returns multiple values file and error (common Go pattern)
	if err != nil {         //nil means no value is assigned to a variable
		return err
	}
	defer f.Close() //defer statement defers the execution of a function until the surrounding function returns. in this case we will close the file we opened. puts f.Close in the call stack

	scanner := bufio.NewScanner(f) //scanner object from bufio to read it to memory(global variable called maze). reads file line by line and appends it to maze slice (clean up purposes)
	for scanner.Scan() {           //scanner.Scan() will return true while there is something to be read from the file
		line := scanner.Text()                          //returns the next line of input
		mazeloader.maze = append(mazeloader.maze, line) //built-in function append adds a new element to the maze slice
		//slices: lightweight ds, similar to array but more powerful. variable-length sequence which stores elements of similar type. Not a fixed size
		//a[low : high]
	}

	fmt.Println("Finished loading MAZE")
	return nil
	//f.Close is called implicitly
}

func (mazeloader *MazeLoader) getmaze() []string {
	maze := mazeloader.maze
	return maze
}
