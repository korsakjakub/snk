package main

import (
	"fmt"
	"os"
	"os/exec"
	"unicode/utf8"
)

var conf Config

func main() {

	conf = loadConfig("..")
	snake := actor{s: shape{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}}, r: '#'}
	fruit := actor{s: shape{{x: 0, y: 1}}, r: '*'}
	if conf.Width*conf.Height == 0 {
		return
	}
	s := screen{
		s:      nil,
		width:  conf.Width,
		height: conf.Height,
	}
	_ = s.drawScreen(snake, fruit)
	points := 0
	game := game{snake: snake, fruit: fruit, screen: s, points: points}
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
	err := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	if err != nil {
		return 0
	}
	// do not display entered characters on the screen
	err = exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	if err != nil {
		return 0
	}
	// restore the echoing state when exiting
	defer func(command *exec.Cmd) {
		err := command.Run()
		if err != nil {

		}
	}(exec.Command("stty", "-F", "/dev/tty", "echo"))

	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		return 0
	}
	r, _ := utf8.DecodeRune(b)
	return r
}
func clearCommand() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}
}
