package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

const PaddleSymbol = 0x2588
const PaddleHeight = 4

type Paddle struct {
	row, col, width, height int
}

var screen tcell.Screen
var player1 *Paddle
var player2 *Paddle
var debugLog string

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {

	InitScreen()
	InitGameState()
	inputChan := InitUserInput()

	for {
		drawState()
		time.Sleep(50 * time.Millisecond)

		key := ReadInput(inputChan)
		HandleUserInput(key)
	}
}

func InitScreen() {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err = screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)
}

func InitUserInput() chan string {
	inputChan := make(chan string)
	go func() {
		for {
			switch ev := screen.PollEvent().(type) {
			case *tcell.EventKey:
				// TODO
				inputChan <- ev.Name()
			}
		}
	}()
	return inputChan
}

func InitGameState() {
	width, height := screen.Size()
	paddleStart := height/2 - PaddleHeight/2

	player1 = &Paddle{
		row: paddleStart, col: 0, width: 1, height: PaddleHeight,
	}
	player2 = &Paddle{
		row: paddleStart, col: width - 1, width: 1, height: PaddleHeight,
	}
}

func drawState() {
	screen.Clear()
	PrintString(0, 0, debugLog)
	Print(player1.row, player1.col, player1.width, player1.height, PaddleSymbol)
	Print(player2.row, player2.col, player2.width, player2.height, PaddleSymbol)
	screen.Show()
}

func ReadInput(inputChan chan string) string {
	var key string
	select {
	case key = <-inputChan:
	default:
		key = ""
	}

	return key
}

func HandleUserInput(key string) {
	_, screenHeight := screen.Size()
	if key == "Rune[q]" {
		screen.Fini()
		os.Exit(0)
	} else if key == "Rune[w]" && player1.row > 0 {
		player1.row--
	} else if key == "Rune[s]" && player1.row+player1.height < screenHeight {
		player1.row++
	} else if key == "Up" && player2.row > 0 {
		player2.row--
	} else if key == "Down" && player2.row+player2.height < screenHeight {
		player2.row++
	}
}

func PrintString(row, col int, str string) {
	for _, c := range str {
		screen.SetContent(col, row, c, nil, tcell.StyleDefault)
		col += 1
	}
}

func Print(row, col, width, height int, ch rune) {
	for r := 0; r < height; r++ {

		for c := 0; c < width; c++ {
			screen.SetContent(col+c, row+r, ch, nil, tcell.StyleDefault)
		}

	}
}
