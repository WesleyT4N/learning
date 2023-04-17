package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

const PlayerPrompt = "Please enter the number of players"

const BadPlayerInputErrorMsg = "Invalid value received for number of players"

const BadWinnerInputErrorMessage = "Invalid value received for winner"

func (cli *CLI) PlayPoker() {
	fmt.Fprintf(cli.out, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	if err != nil {
		fmt.Fprintf(cli.out, BadPlayerInputErrorMsg)
		return
	}

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner, err := extractWinner(winnerInput)
	if err != nil {
		fmt.Fprintf(cli.out, BadWinnerInputErrorMessage)
		return
	}

	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		bufio.NewScanner(in),
		out,
		game,
	}
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", fmt.Errorf("Error occurred extracting winner")
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}
