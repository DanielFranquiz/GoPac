//runnable programs must have a package called main and a main function
package main

//makes code in other packages accessible to this program
import (
	"fmt"
	"bufio" //buffered IO package will have the scanner object
	"log"
	"os" //package to open outside files and release the file handler
)

var maze []string

func loadMaze(file string) error {
	//:= short-assignment operator that automatically infers the type of variable(s) based on the value(s) on the right hand side. commonly used in if, for, switch
	f, err := os.Open(file) //Open function from os package which returns multiple values file and error (common Go pattern)
	if err != nil { //nil means no value is assigned to a variable
		return err
	}
	defer f.Close() //defer statement defers the execution of a function until the surrounding function returns. in this case we will close the file we opened. puts f.Close in the call stack
	
	scanner := bufio.NewScanner(f) //scanner object from bufio to read it to memory(global variable called maze). reads file line by line and appends it to maze slice (clean up purposes)
	for scanner.Scan() { //scanner.Scan() will return true while there is something to be read from the file
		line := scanner.Text() //returns the next line of input
		maze = append(maze, line) //built-in function append adds a new element to the maze slice
		//slices: lightweight ds, similar to array but more powerful. variable-length sequence which stores elements of similar type. Not a fixed size 
		//a[low : high]
	}
	
	return nil
	//f.Close is called implicitly
}

func printScreen() {
	for _, line := range maze { //_ is a placeholder for where the compiler would expet a variable name. we are ignoring that value
		//range - first return value is the index of the element starting from 0, second return value is the value itself
		fmt.Println(line)
	}
}

//entry point of program defined as function with func
//main program takes no parameters and returns nothing
func main() {
	// initialize game

	// load resources
	err := loadMaze("maze01.txt")
	if err != nil {
		log.Println("failed to load maze:", err)
		return
	}

	// game loop never ending
	for {
		// update screen
		printScreen()

		// process input

		// process movement

		// process collisions

		// check game over

		// Temp: break infinite loop
		//fmt package formatted input/output
		fmt.Println("Hello, Pac Go!")
		break

		// repeat
	}
}
