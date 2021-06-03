// package table implements
// the table abstraction for
// the game
package table

import (
    "fmt"
)

// Table is the abstraction to run the game and organize
// game's data
type Table struct {
    Result  string
    Grid    [][]string
    Rounds  int
}

// NewTable generates a new "empty" table to play a
// tic-tac-toe game
func NewTable() Table {
    return Table{
        Result: "",
        Grid:   [][]string{
            {"-", "-", "-"},
            {"-", "-", "-"},
            {"-", "-", "-"},
        },
    }
}

// Print displays the table on screen
func (t *Table) Print() {
    for i := 1; i <= 3; i++ {
        fmt.Printf("  %d  ", i)
    }
    fmt.Println()
    for ki, i := range t.Grid {
        for _, j := range i {
            fmt.Printf("| %s |", j)
        }
        fmt.Printf(" < %d", ki+1)
        fmt.Println()
    }
}

// Mark will insert the input in given coordinates
func (t *Table) Mark(x int, y int, input string) {
    t.Grid[x-1][y-1] = input
}

// IsMarked check if the place selected was already marked
func (t *Table) IsMarked(x, y int) bool {
    if t.Grid[x-1][y-1] != "-" {
        return true
    }
    return false
}

// Validate will check if there's already a game winner,
// if there were 9 or more rounds the game stops, because
// it's "toe"
func (t *Table) Validate() {
    for _, i := range t.Grid {
        if IsAll(i, "X") {
            t.Result = "tic"
        }
        if IsAll(i, "O") {
            t.Result = "tac"
        }
    }

    columns := make([][]string, len(t.Grid))
    for i := range columns {
        columns[i] = make([]string, len(t.Grid)) 
    }

    for i := range t.Grid {
        for j := range t.Grid[i] {
            columns[j][i] = t.Grid[i][j]
        }
    }

    for _, i := range columns {
        if IsAll(i, "X") {
            t.Result = "tic"
        }
        if IsAll(i, "O") {
            t.Result = "tac"
        }
    }

    diagonals := make([][]string, len(t.Grid))
    for i := range diagonals {
        diagonals[i] = make([]string, len(t.Grid))
    }

    for i := 0; i < len(t.Grid); i++ {
        diagonals[0][i] = t.Grid[i][i]
    }

    for i := len(t.Grid)-1; i >= 0; i-- {
        diagonals[1][i] = t.Grid[i][i]
    }

    for _, i := range diagonals {
        if IsAll(i, "X") {
            t.Result = "tic"
        }
        if IsAll(i, "O") {
            t.Result = "tac"
        }
    }

    t.Rounds++
    if t.Rounds >= 9 {
        t.Result = "toe"
    }
}

