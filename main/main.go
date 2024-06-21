package main

import (
	"AFootGolf9/GoNonsence/tictactoe"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func whatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's how they say %s", r.URL.Path[6:])
}

var ticTacToeBoard [3][3]int
var last int = 0

func ticTacToeHandler(w http.ResponseWriter, r *http.Request) {
	var num1, num2 string
	num1 = r.URL.Query().Get("num1")
	num2 = r.URL.Query().Get("num2")
	if num1 == "" || num2 == "" {
		fmt.Fprintf(w, "Please provide num1 and num2")
		return
	}

	n1, err := strconv.Atoi(num1)
	if err != nil {
		fmt.Fprintf(w, "num1 is not a number")
		return
	}
	n2, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Fprintf(w, "num2 is not a number")
		return
	}

	if ticTacToeBoard[n1][n2] != 0 {
		fmt.Fprintf(w, "This cell is already taken")
		return
	}

	tictactoe.PlayTurn(n1, n2, last+1, &ticTacToeBoard)

	if last == 0 {
		last = 1
	} else {
		last = 0
	}
}

func boardHandler(w http.ResponseWriter, r *http.Request) {
	// for i := 0; i < 5; i++ {
	// 	if i%2 == 0 {
	// 		for j := 0; j < 5; j++ {
	// 			if j%2 == 0 {
	// 				if ticTacToeBoard[i/2][j/2] == 0 {
	// 					fmt.Fprintf(w, " ")
	// 				}
	// 				if ticTacToeBoard[i/2][j/2] == 1 {
	// 					fmt.Fprintf(w, "X")
	// 				}
	// 				if ticTacToeBoard[i/2][j/2] == 2 {
	// 					fmt.Fprintf(w, "O")
	// 				}
	// 			} else {
	// 				fmt.Fprintf(w, "|")
	// 			}
	// 		}
	// 	} else {
	// 		fmt.Fprintf(w, "-+-+-")
	// 	}
	// }

	fmt.Fprint(w, tictactoe.Board(ticTacToeBoard))
}

func main() {
	http.HandleFunc("/what/", whatHandler)
	http.HandleFunc("/tic-tac-toe/", ticTacToeHandler)
	http.HandleFunc("/board/", boardHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
