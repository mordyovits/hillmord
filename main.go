package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hillmord/game"
)

func main() {
	fmt.Println(game.TitleScreen())

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What shall we call you, brave fool? > ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Mudrick the Uncertain"
	}

	g := game.New(name)
	g.Run(reader)
}
