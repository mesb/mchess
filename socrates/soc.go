/*
Socrates is an interpreter for chess rules

Socrates receives commands that represent moves and a board
and it moves the pieces as directed by the command.

Socrates knows the rules of the game and reports and error if
there was an illegal move
*/
package socrates

import (
	"errors"
	_ "fmt"
	"github.com/mesb/mchess/address"
	"github.com/mesb/mchess/board"
	"github.com/mesb/mchess/pieces"
	"github.com/mesb/mchess/socrates/lexer"
	"github.com/mesb/mchess/socrates/validator"
	_ "math/rand"
)

type TOKENTYPE uint

// location of command pieces in token stream
const (
	MCOMMAND TOKENTYPE = iota
	MPIECE
	MDEST
)

// moves piece to new location by swapping
func move(piece pieces.Piece, dest address.Addr, b *board.Board) error {
	d := dest.Index()

	t := b.Content[d]
	tmp := pieces.NewPiece(pieces.Empty, piece.Address(), t.Camp(), t.Status())

	b.Content[d] = pieces.NewPiece(piece.Type(), dest, piece.Camp(), piece.Status())
	b.Content[piece.Address().Index()] = tmp

	return nil

}

// Dialog with socrates
/*
   1. Socrates receives command stream
   2. Socrates calls his lexer to lex the command stream
   3. Soc returns immediately with error message is there are errors
   4. Deconstruct the command to get the pieces: command, piece, destination
   5. Evaluate the operands and validate them
   6. Evaluate the command to obtain its object code
   7. Return with errors if there are errors
   8. Apply command to arguments
   9. Reflect new state on board or update board
   10. Exit with no errors at this point
*/
func Dialog(cmd string, b *board.Board) error {

	toks, err := lexer.Lex(cmd)

	if err != nil {
		return errors.New(err.Error())
	}

	// now obtain the tokens
	comm := toks[MCOMMAND]
	piece := toks[MPIECE]
	dest := toks[MDEST]

	if err := validate.Validate(comm, piece, dest); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
