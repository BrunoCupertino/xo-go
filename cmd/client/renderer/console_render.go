package renderer

import (
	"fmt"

	"github.com/BrunoCupertino/xo-go/internal/client"
	"github.com/BrunoCupertino/xo-go/internal/state"
)

type ConsoleRenderer struct{}

func NewConsoleRenderer() *ConsoleRenderer {
	return &ConsoleRenderer{}
}

func (r *ConsoleRenderer) Render(s *client.ClientGameState) {
	c := s.CurrentChange()

	switch c {
	case state.TeamSelected:
		fmt.Printf("your team is >> %s\n", s.MyTeam())
		printBoard(s)
		if s.MyTeam() == state.OTeam {
			fmt.Printf("%s is your turn, input square number (0...8) > ", s.MyTeam())
		} else {
			fmt.Println("waiting other player to play...")
		}
	case state.BoardChanged:
		printBoard(s)
		if s.IsMyTurn() {
			fmt.Printf("%s is your turn, input square number (0...8) > ", s.MyTeam())
		} else {
			fmt.Println("waiting other player to play...")
		}
	case state.GameOvered:
		printBoard(s)
		if s.Winner() == state.NoTeam {
			fmt.Print("GAME OVER it was a tie")
		} else {
			fmt.Printf("GAME OVER the winner was %s\n", s.Winner())
		}
	}
}

func printBoard(s *client.ClientGameState) {
	round, board := s.GetBoard()

	fmt.Printf("round %d:\n", round)

	fmt.Print(board[0])
	fmt.Print(" | ")
	fmt.Print(board[1])
	fmt.Print(" | ")
	fmt.Print(board[2])
	fmt.Print("\n")

	printRowSeparator()

	fmt.Print(board[3])
	fmt.Print(" | ")
	fmt.Print(board[4])
	fmt.Print(" | ")
	fmt.Print(board[5])
	fmt.Print("\n")

	printRowSeparator()

	fmt.Print(board[6])
	fmt.Print(" | ")
	fmt.Print(board[7])
	fmt.Print(" | ")
	fmt.Print(board[8])
	fmt.Print("\n\n")
}

func printRowSeparator() {
	fmt.Print("---------\n")
}
