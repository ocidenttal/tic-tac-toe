// functions implements the functions
// (that are not methods) used in the
// package table
package table

import (
    "fmt"
    "strings"
)


// IsAll checks if a slice of strings contains just a
// certain string
func IsAll(a []string, c string) bool {
    for i := range a {
        if a[i] != c {
            return false
        }
    }
    return true
}

// GerCoordinate is the function that will be ran
// to get user's coordinates to the next option
func GetCoordinate() (int, int, error) {
    var x, y int
    fmt.Printf("line: ")
    fmt.Scanf("%d", &x)

    fmt.Printf("column: ")
    fmt.Scanf("%d", &y)

    if !IsValid(x) {
        return 0, 0, InvalidCoordinate
    }
    if !IsValid(y) {
        return 0, 0, InvalidCoordinate
    }
    return x, y, nil
}

// GetInput is a function that will take the input of the
// user
func GetInput() (string, error) {
    var input string
    fmt.Printf("input (X | O): ")
    fmt.Scanf("%s", &input)

    input = strings.ToUpper(input)
    if input == "X" || input == "O" {
        return input, nil
    }
    return "", InvalidInput
}

// IsValid will check if the coordinate inserted is
// or not valid
func IsValid(n int) bool {
    if n>=1 && n<=3 {
        return true
    }
    return false
}
