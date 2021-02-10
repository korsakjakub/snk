package main

import (
	"fmt"
	"testing"
)

func TestMove(t *testing.T) {
	snake := actor{s: shape{{x: 0, y: 0},{x : 1, y: 0}, {x:2, y:0}}, r: '#'}
	fruit := actor{s: shape{{x: 0, y: 1}}, r: '*'}
	s := screen{}
	_ = s.drawScreen(snake,fruit)
	points := 0
	nextMove := up
	game := game{snake: snake, fruit: fruit, points: points, screen: s}
	game = game.move(nextMove)
	fmt.Print(game.drawScreen())
	game = game.move(nextMove)
	fmt.Print(game.drawScreen())
	game = game.move(nextMove)
	fmt.Print(game.drawScreen())
}