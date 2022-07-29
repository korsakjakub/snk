package main

import (
	"reflect"
	"testing"
)

func TestMove(t *testing.T) {
	conf = Config{
		Width:  4,
		Height: 4,
	}
	snake := actor{s: shape{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}}, r: '#'}
	fruit := actor{s: shape{{x: 0, y: 1}}, r: '*'}
	s := screen{nil, 4, 4}
	_ = s.drawScreen(snake, fruit)
	points := 0
	nextMove := up
	game := game{snake: snake, fruit: fruit, points: points, screen: s}
	game = game.move(nextMove)
	t.Log(game.drawScreen())
	game = game.move(nextMove)
	t.Log(game.drawScreen())
	game = game.move(nextMove)
	t.Log(game.drawScreen())

	want := [][]rune{{46, 46, 46, 46}, {42, 46, 35, 46}, {46, 46, 35, 46}, {46, 46, 35, 46}}
	if !reflect.DeepEqual(game.screen.s, want) {
		t.Errorf("wanted: %v, got: %v", want, game.screen.s)
	}
}
