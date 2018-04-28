/*
  Package validate contains routines and procedures for checking the
  rules of chess

  Validate can validate the consequences of a command

  Validate has a knowledge of chess rules

  Socrates uses the validate package to check if a movement or command is allowed.
*/
package validate

import (
	_ "fmt"
	"github.com/mesb/mchess/address"
	_ "github.com/mesb/mchess/pieces"
	"github.com/mesb/mchess/socrates/lexer"
)

func ValidateAddress(a address.Addr) error {
	return nil
}

func Validate(commandName, piece, destination lexer.Token) error {
	_ = commandName.Value()
	_ = piece.Value()
	_ = piece.Value()

	// validate operands
	// the addresses should be within bounds
	//

	return nil
}
