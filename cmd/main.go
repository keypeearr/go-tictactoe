package main

import "github.com/keypeearr/tictactoe/src/server"

func main() {
	if err := server.Run(); err != nil {
		panic(err)
	}
}
