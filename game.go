package main

import (
	"math/rand"
	"time"
)

type move rune

const (
	zero move = iota
	up
	down
	left
	right
)

type game struct {
	snake  actor
	fruit  actor
	screen screen
	points int
}

func (g game) move(move move) game {
	snake := g.snake.s
	snake = append(snake[:0], snake[1:]...)
	direction := coord{x: 0, y: 0}
	switch move {
	case up:
		direction.y = -1
	case down:
		direction.y = 1
	case left:
		direction.x = -1
	case right:
		direction.x = 1
	}
	g.snake.s = snake.moduloAdd(direction)
	snake = g.snake.s
	if g.snake.s.contains(g.fruit.s[0]) {
		rand.Seed(time.Now().UnixNano())
		g.fruit.s[0] = g.generateNewCoords()
		g.snake.s = snake.moduloAdd(direction)
		g.points += 1
	}
	g.screen = *g.screen.drawScreen(g.snake, g.fruit)
	return g
}

func (g game) drawScreen() *screen {
	return g.screen.drawScreen(g.snake, g.fruit)
}

func (g game) generateNewCoords() coord {
	i := 1
	c := coord{}
	for {
		i += 1
		c = coord{x: rand.Int() % g.screen.width, y: rand.Int() % g.screen.height}
		if !g.snake.s.contains(c) || i == g.screen.width*g.screen.height {
			break
		}
	}
	return c
}
