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
	gameFinishedHandler(w, r)
}

func boardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, tictactoe.Board(ticTacToeBoard))
}

func gameFinishedHandler(w http.ResponseWriter, r *http.Request) {
	if tictactoe.CheckWin(ticTacToeBoard) != 0 {
		fmt.Fprintf(w, "Player %d wins", tictactoe.CheckWin(ticTacToeBoard))
		return
	}
	if tictactoe.CheckDraw(ticTacToeBoard) {
		fmt.Fprintf(w, "It's a draw")
		return
	}
	fmt.Fprintf(w, "Game is not finished")
}

func resetGameGameHandler(w http.ResponseWriter, r *http.Request) {
	ticTacToeBoard = [3][3]int{}
	last = 0
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test")
	fmt.Fprintf(w, "Test")
}

func main() {
	http.HandleFunc("/what/", whatHandler)
	http.HandleFunc("/tic-tac-toe/", ticTacToeHandler)
	http.HandleFunc("/board/", boardHandler)
	http.HandleFunc("/game-finished/", gameFinishedHandler)
	http.HandleFunc("/reset-game/", resetGameGameHandler)
	http.HandleFunc("/test/", testHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
