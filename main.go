package main

import (
	"guess_number/guess_game"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	game := guess_game.New(guess_game.Option{
		Max:   100,
		Limit: 10,
	})
	game.Play()
}
