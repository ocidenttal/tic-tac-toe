package main

import (
	"fmt"
	"time"

	"github.com/ocidenttal/tic-tac-toe/table"
)

func main() {
	tab := table.NewTable()

	// game mainloop
	for tab.Result == "" {
		tab.Print()

		x, y, err := table.GetCoordinate()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
			ClearScreen()
			continue
		}

		input, err := table.GetInput()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
			ClearScreen()
			continue
		}

		if tab.IsMarked(x, y) {
			fmt.Println("this place is already marked")
			time.Sleep(time.Second)
			ClearScreen()
			continue
		}

		tab.Mark(x, y, input)
		tab.Validate()
		ClearScreen()
	}
	tab.Print()
	fmt.Printf("\n%s!\n", tab.Result)
}

// ClearScreen will clear the terminal screen
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
