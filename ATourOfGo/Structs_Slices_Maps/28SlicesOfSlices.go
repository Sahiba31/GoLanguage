package main
import (
	"fmt"
	"strings"
)


//tic-tac-toe board game
func main() {
	board := [][]string {
		[]string {"-", "-", "-"},
		[]string {"-", "-", "-"},
		[]string {"-", "-", "-"},
	}

	board[0][0] = "X"
	board[0][2] = "O"
	board[1][1] = "X"
	board[2][2] = "O"
	board[2][0] = "O"
	for i := 0 ; i < len(board) ; i++ {
		fmt.Printf("%s\n", strings.Join(board[i]," "))
	}
}