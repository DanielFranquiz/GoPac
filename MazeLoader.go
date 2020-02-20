package main;

import "bufio";
import "os";
import "fmt"

/*		USAGE:
		mazeLoader := MazeLoader{};
		// load resources
		err := mazeLoader.loadMaze("maze01.txt")
		if err != nil {
			fmt.Println("failed to load maze:", err)
			return
		}
		mazeLoader.printScreen();
*/

type MazeLoader struct{
	maze []string;
}

func (i *MazeLoader) loadMaze(file string) error {
	fmt.Println("LOADING MAZE")
    f, err := os.Open(file)// opening file
    if err != nil {// if error freak out
        return err
    }
    defer f.Close()// ?????????

    scanner := bufio.NewScanner(f)/// scanner class
    for scanner.Scan() {
        line := scanner.Text()
        i.maze = append(i.maze, line)
    }
	fmt.Println("Finished loading MAZE")
    return nil/// returning null
}

func (i *MazeLoader) printScreen() {
	fmt.Println("PRINTING SCREEN")
	
    for _, line := range i.maze {// looping through each line
        fmt.Println(line)
    }
}
