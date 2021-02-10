package main

import (
	"fmt"
	"os"
	"os/exec"
	"unicode/utf8"
)

func main() {
	snake := actor{s: shape{{x: 0, y: 0},{x : 1, y: 0}, {x:2, y:0}}, r: '#'}
	fruit := actor{s: shape{{x: 0, y: 1}}, r: '*'}
	s := screen{}
	_ = s.drawScreen(snake,fruit)
	points := 0
	game := game{snake: snake,fruit: fruit,screen: s,points: points}
	move := zero

	for {
		fmt.Printf("Points: %d\n\n", game.points)
		fmt.Print(game.drawScreen())
		input := getChar()
		switch input {
		case 'w':
			move = up
		case 's':
			move = down
		case 'd':
			move = right
		case 'a':
			move = left
		}
		game = game.move(move)
		clearCommand()
	}
}


func getChar() rune {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	// restore the echoing state when exiting
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	b := make([]byte, 1)
	os.Stdin.Read(b)
	r, _ := utf8.DecodeRune(b)
	return r
}
func clearCommand() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}