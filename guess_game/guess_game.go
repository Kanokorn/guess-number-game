package guess_game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type GuessResult struct {
	Number int
	Err    error
}

type GameState int

const (
	CONTINUE GameState = 0
	OVER     GameState = 1
)

type GuessGame struct {
	Target      int
	Limit       int
	Used        int
	GuessResult GuessResult
	State       GameState
	Msg         string
	Max         int
}

type Option struct {
	Max   int
	Limit int
}

func New(opt Option) *GuessGame {
	return &GuessGame{
		Target: rand.Intn(opt.Max) + 1,
		Limit:  opt.Limit,
		Max:    opt.Max,
	}
}

func (gg *GuessGame) Play() {
	for gg.Used = 1; ; gg.Used++ {
		gg.Guess()
		fmt.Println(gg.Msg)

		if gg.IsOver() {
			return
		}

		fmt.Println(gg.Limit-gg.Used, "guess left")
	}
}

func (gg *GuessGame) Guess() {
	gg.getGuessNumber()
	gg.checkGuessResult()
}

func (gg *GuessGame) getGuessNumber() {
	var number int

	fmt.Printf("Guess the number from 1 to %d: ", gg.Max)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		gg.GuessResult = GuessResult{Err: err}
		return
	}

	input = strings.TrimSpace(input)
	number, err = strconv.Atoi(input)
	if err != nil {
		gg.GuessResult = GuessResult{Err: err}
		return
	}

	gg.GuessResult = GuessResult{Number: number}
	return
}

func (gg *GuessGame) checkGuessResult() {
	switch {
	case gg.GuessResult.Number == gg.Target:
		gg.State, gg.Msg = OVER, "Good job! You guessed it!"

	case gg.Used == gg.Limit:
		gg.State, gg.Msg = OVER, fmt.Sprint("Sorry. You didn't guess my number. It was: ", gg.Target)

	case gg.GuessResult.Err != nil:
		gg.State, gg.Msg = CONTINUE, fmt.Sprintf("Your input is invalid! Please put the number during 1 to %d.", gg.Max)

	case gg.GuessResult.Number < gg.Target:
		gg.State, gg.Msg = CONTINUE, "Oops. Your guess was LOW."

	case gg.GuessResult.Number > gg.Target:
		gg.State, gg.Msg = CONTINUE, "Oops. Your guess was HIGH."
	}
}

func (gg *GuessGame) IsOver() bool {
	return gg.State == OVER
}
