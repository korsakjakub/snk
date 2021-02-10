package main

type actor struct {
	s shape
	r rune
}

func (s shape) contains(c coord) bool {
	for _, e := range s {
		if e.x == c.x && e.y == c.y {
			return true
		}
	}
	return false
}