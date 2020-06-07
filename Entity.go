package main

//"os"

type sprite struct {
	row int
	col int
}

var player sprite

func CaptureEntitiesPosition(maze []string) {
	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = sprite{row, col}
			}
		}
	}
}
