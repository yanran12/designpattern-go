package template

import "fmt"

type Game interface {
	Start()
	Playing()
	End()
}

type AbsGame struct {
}

func (ag *AbsGame) Start() {
	fmt.Println("start game")
}
func (ag *AbsGame) Playing() {
	fmt.Println("game playing")
}
func (ag *AbsGame) End() {
	fmt.Println("the end")
}

type FootBall struct {
	*AbsGame
}

func (fb *FootBall) Playing() {
	fmt.Println("playing FootBall")
}

type BasketBall struct {
	*AbsGame
}

func (bb *BasketBall) Playing() {
	fmt.Println("playing basketball ")
}

func Run() {

	games := []Game{&BasketBall{}, &FootBall{}}
	for _, game := range games {
		game.Start()
		game.Playing()
		game.End()
	}
}
