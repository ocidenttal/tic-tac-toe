package main

import (
	"fmt"
	. "tic-tac-toe/table"
	"time"
)

func main() {
	table := NewTable()

	// game mainloop
	for table.Result == "" {
		table.Print()

		x, y, err := GetCoordinate()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
			ClearScreen()
			continue
		}

		input, err := GetInput()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
			ClearScreen()
			continue
		}

		if table.IsMarked(x, y) {
			fmt.Println("this place is already marked")
			time.Sleep(time.Second)
			ClearScreen()
			continue
		}

		table.Mark(x, y, input)
		table.Validate()
		ClearScreen()
	}
	table.Print()
	fmt.Printf("\n%s!\n", table.Result)
}

// ClearScreen will clear the terminal screen
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
