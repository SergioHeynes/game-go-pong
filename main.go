package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

func PrintStr(s tcell.Screen, row, col int, str string) {
	for _, c := range str {
		s.SetContent(col, row, c, nil, tcell.StyleDefault)
		col += 1
	}
}

func Print(s tcell.Screen, row, col, width, height int, ch rune) {
	for r := 0; r < height; r++ {

		for c := 0; c < width; c++ {
			s.SetContent(col+c, row+r, ch, nil, tcell.StyleDefault)
		}

	}
}

func displayHelloWorld(screen tcell.Screen) {
	screen.Clear()
	// PrintStr(screen, 2, 1, "Hello, World!")
	Print(screen, 1, 4, 5, 5, '*')
	screen.Show()
}

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {

	screen := InitScreen()
	displayHelloWorld(screen)

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEnter {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}

func InitScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)
	return screen
}
