package main

import "fmt"

const (
	width = 40
	height = 20
)

type screen [height][width]rune

func (s * screen) clear() *screen {
	for i := 0; i < height; i += 1 {
		for j := 0; j < width; j += 1 {
			s[i][j] = '.'
		}
	}
	return s
}

func (s *screen) getWidth() int {
	return width
}
func (s *screen) getHeight() int {
	return height
}

func (s *screen) drawScreen(snake actor, fruit actor) *screen {
	s.clear()
	for _, e := range snake.s {
		s[e.y][e.x] = snake.r
	}
	for _, e := range fruit.s {
		s[e.y][e.x] = fruit.r
	}
	return s
}

func (s *screen) String() string {
	str := ""
	for _, e := range s {
		for _, f := range e {
			str += fmt.Sprintf("%c ",f)
		}
		str += fmt.Sprintf("\n")
	}
	return str
}

type coord struct {
	x int
	y int
}

func (s shape) moduloAdd(direction coord) shape{
	x := (s[len(s)-1].x + direction.x) % width
	if x < 0 {
		x += width
	}
	y := (s[len(s)-1].y + direction.y) % height
	if y < 0 {
		y += height
	}
	return append(s, coord{x: x, y: y})
}

type shape []coord