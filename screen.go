package main

import "fmt"

type screen struct {
	s      [][]rune
	width  int
	height int
}

func (s *screen) clear() *screen {
	grid := make([][]rune, s.width)
	for i := 0; i < s.width; i++ {
		grid[i] = make([]rune, s.height)
	}
	s.s = grid
	for i := 0; i < s.height; i += 1 {
		for j := 0; j < s.width; j += 1 {
			s.s[i][j] = '.'
		}
	}
	return s
}

func (s *screen) drawScreen(snake actor, fruit actor) *screen {
	s.clear()
	for _, e := range snake.s {
		s.s[e.y][e.x] = snake.r
	}
	for _, e := range fruit.s {
		s.s[e.y][e.x] = fruit.r
	}
	return s
}

func (s *screen) String() string {
	str := ""
	for _, e := range s.s {
		for _, f := range e {
			str += fmt.Sprintf("%c ", f)
		}
		str += fmt.Sprintf("\n")
	}
	return str
}

type coord struct {
	x int
	y int
}

func (s shape) moduloAdd(direction coord) shape {
	x := (s[len(s)-1].x + direction.x) % conf.Width
	if x < 0 {
		x += conf.Width
	}
	y := (s[len(s)-1].y + direction.y) % conf.Height
	if y < 0 {
		y += conf.Height
	}
	return append(s, coord{x: x, y: y})
}

type shape []coord
