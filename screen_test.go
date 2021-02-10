package main

import (
	"fmt"
	"testing"
)

func TestDrawScreen(t *testing.T) {
	snake := actor{s: shape{{x: 0, y: 0},{x : 1, y: 0}, {x:2, y:0}}, r: '#'}
	fruit := actor{s: shape{{x: 0, y: 1}}, r: '*'}
	s := screen{}
	got := s.drawScreen(snake,fruit)
	want := [][]rune{{35,42,46},{35,46,46},{35,46,46}}

	for i := 0; i < len(want); i += 1 {
		for j := 0; j < len(want[0]); j += 1 {
			if got[i][j] != want[i][j] {
				t.Logf("wrong screen draw. Got %v, expected %v", got, want)
			}
		}
	}
}

func TestPrintScreen(t *testing.T) {
	snake := actor{s: shape{{x: 0, y: 0},{x : 1, y: 0}, {x:2, y:0}}, r: '#'}
	fruit := actor{s: shape{{x: 0, y: 1}}, r: '*'}
	s := screen{}
	got := s.drawScreen(snake,fruit)
	fmt.Print(got)
}

func TestModuloAdd(t *testing.T) {
	s := shape{{x : 1, y: 0}, {x:2, y:0}}
	d := coord{x: 0, y: -1}
	out := s.moduloAdd(d)
	t.Log(out)

	t.Log((-1)%3)
}