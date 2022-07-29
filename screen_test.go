package main

import (
	"reflect"
	"testing"
)

func TestDrawScreen(t *testing.T) {
	snake := actor{s: shape{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}}, r: '#'}
	fruit := actor{s: shape{{x: 0, y: 1}}, r: '*'}
	s := screen{nil, 3, 3}
	got := s.drawScreen(snake, fruit)
	want := [][]rune{{35, 35, 35}, {42, 46, 46}, {46, 46, 46}}

	if !reflect.DeepEqual(got.s, want) {
		t.Errorf("wanted: %v, got: %v", want, got)
	}
}

func TestPrintScreen(t *testing.T) {
	conf = Config{
		Width:  3,
		Height: 3,
	}
	snake := actor{s: shape{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}}, r: '#'}
	fruit := actor{s: shape{{x: 0, y: 1}}, r: '*'}
	s := screen{nil, 3, 3}
	_ = s.drawScreen(snake, fruit)
	got := s.s
	want := [][]rune{{35, 35, 35}, {42, 46, 46}, {46, 46, 46}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v, got: %v", want, got)
	}
}

func TestModuloAdd(t *testing.T) {
	conf = Config{
		Width:  3,
		Height: 3,
	}
	s := shape{{x: 1, y: 0}, {x: 2, y: 0}}
	d := coord{x: 0, y: -1}
	got := s.moduloAdd(d)
	want := shape{{1, 0}, {2, 0}, {2, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v, got: %v", want, got)
	}
}
