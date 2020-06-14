package main

import (
	"github.com/danicat/simpleansi"
)

func moveCursor(row, col int) {
    if cfg.UseEmoji {
        simpleansi.MoveCursor(row, col*2)
    } else {
        simpleansi.MoveCursor(row, col)
    }
}