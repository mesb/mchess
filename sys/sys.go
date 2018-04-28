/*
Package Sys stands for the mchess system. It controls the display of menus,
the reception of input from the environment and the return of results.

Sys is the the reactive engine for mchess. It is mchess's means of interacting with
its environment...
*/

package sys

import (
	"bufio"
	"fmt"
	"github.com/mesb/mchess/board"
	"github.com/mesb/mchess/socrates"
	"os"
	"strings"
)

// some silly constants like ZERO
const (
	ZERO = 0
)

// system files
// for example, input is read from the Stdio file
var input *bufio.Reader

// a board for this session
var brd *board.Board

// auxiliary method for displaying messages
func display(s string) {
	fmt.Print(s)
}

// function to display welcome message
func displayMessage() {
	display("       Welcome to MCHESS\n")
	display("-----------------------------\n")
	display("Enter 'q' or 'quit' to exit\n")
	display("Enter 'b' or 'board' to see board\n")
}

// initialize package
// create file for reading
func init() {
	input = bufio.NewReader(os.Stdin)
	brd = board.InitBoard()

	displayMessage()
}

func Run() {

FORERUNNER: // a label for the main event loop of mchess
	for {
		go display("0~-> ")
		text, _ := input.ReadString('\n')
		text = strings.TrimSpace(text)

		if len(text) <= 0 {
			continue
		}

		text = strings.ToLower(text)
		cmd := string(text[0])

		switch cmd {
		case "b":
			{
				board.Printer(*brd)
			}

		case "q":
			{
				fmt.Println("Good Bye!")
				break FORERUNNER
			}
		case "m":
			{
				err := socrates.Dialog(text, brd)

				if err != nil {
					fmt.Printf("%s", err)
					break
				}

				board.Printer(*brd)
			}

		default:
			{
				fmt.Println("Enter 'b' or 'q'")
			}

		}

	}
}
