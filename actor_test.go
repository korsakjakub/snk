package main

import (
	"testing"
)

func TestContains(t *testing.T) {
	snake := actor{s: shape{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}}, r: '#'}
	fruit1 := actor{s: shape{{x: 0, y: 1}}, r: '*'}
	fruit2 := actor{s: shape{{x: 1, y: 0}}, r: '*'}
	fruit3 := actor{s: shape{{x: 0, y: 0}}, r: '*'}

	if snake.s.contains(fruit1.s[0]) {
		t.Errorf("snake %v shouldn't be eating fruit %v", snake, fruit1)
	}
	if !snake.s.contains(fruit2.s[0]) {
		t.Errorf("snake %v should be eating fruit %v", snake, fruit2)
	}
	if !snake.s.contains(fruit3.s[0]) {
		t.Errorf("snake %v should be eating fruit %v", snake, fruit3)
	}
}
