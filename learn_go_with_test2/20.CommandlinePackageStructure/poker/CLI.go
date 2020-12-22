package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI {
		playerStore:store,
		in: bufio.NewScanner(in),
	}
}

func (c *CLI)PlayPoker() {
	userInput := c.readLine()
	name := c.getWinnerName(userInput)
	c.playerStore.RecordWin(name)
}

func (c *CLI)getWinnerName(input string) string {
	return strings.Split(input, " ")[0]
}

func (c *CLI)readLine() string {
	c.in.Scan()
	return c.in.Text()
}