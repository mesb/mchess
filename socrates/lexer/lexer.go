/*
   The lexer reads a line of text or mchess command and transforms them into mchess tokens

   Flow
   1. Receive text
   2. Emit single words
   3. Convert emited word to token
   4. Return token list

   A simpler method is to use regular expressions to obtain the tokens
*/

package lexer

import (
	"bytes"
	_ "encoding/json"
	"errors"
	"fmt"
	_ "github.com/mesb/mchess/pieces"
	"regexp"
)

// constants for the regexp slice indices and the various pieces of a command
const (
	_ = iota
	CMD
	CHESSPIECE
	FILE1
	RANK1
	FILE2
	RANK2
)

type TokenType uint

// enumerate the various token types
const (
	TokenError = iota
	EOT        // end of token stream
	PIECE
	COMMAND
	ADDRESS
)

type Token interface {
	Type() string
	Value() string // you can store this an bytes to make it universal
}

// memory rep and conrete type for a token
type token struct {
	typ   TokenType
	value string // encode as json perhaps
}

// encode all the relevant information of a token as a json object
// and return the json object as an io reader
func (t token) Value() string {
	return t.value
}

// string representation of a token
func (t token) String() string {
	return fmt.Sprintf("(Type: %s, Value: %v)", t.Type(), t.Value())
}

// returns the type of a token as a string
func (t token) Type() string {
	buf := []byte{}
	buff := bytes.NewBuffer(buf)

	switch t.typ {
	case TokenError:
		{
			buff.Write([]byte("TOKEN ERROR"))
		}

	case EOT:
		{
			buff.Write([]byte("EOT"))
		}

	case PIECE:
		{
			buff.Write([]byte("MPIECE"))
		}

	/* case KING:
		{
			buff.Write([]byte("KING"))
		}

	case QUEEN:
		{
			buff.Write([]byte("QUEEN"))
		}

	case ROOK:
		{
			buff.Write([]byte("ROOK"))
		}

	case BISHOP:
		{
			buff.Write([]byte("BISHOP"))
		}

	case KNIGHT:
		{
			buff.Write([]byte("KNIGHT"))
		}

	case PAWN:
		{
			buff.Write([]byte("PAWN"))
		}
	*/

	case COMMAND:
		{
			buff.Write([]byte("COMMAND"))
		}

	case ADDRESS:
		{
			buff.Write([]byte("ADDRESS"))
		}
	default:
		{
			buff.Write([]byte("Invalid Token"))
		}

	}

	return fmt.Sprintf("%s", buff)
}

// describe the various tokens
/*
func (t token) Value() io.Reader {

}*/

var MOVEGRAMMAR *regexp.Regexp

func init() {
	MOVEGRAMMAR = regexp.MustCompile(`^([a-z]{3})\s+(kg|qn|kt|bp|rk|pn)-([a-h])([1-8])\s+([a-h])([1-8])\s*$`)
}

// return the token list. Use regular expressions to deconstruct line of text
func Lex(s string) ([]Token, error) {
	if !MOVEGRAMMAR.MatchString(s) {
		return []Token{}, errors.New(fmt.Sprintf("You entered the wrong command: %v", s))
	}

	// get the slice with all the sub matches or groupings
	// buf[1] is the command name or first group and CMD = 1
	// buf[2] is the chess piece in the PIECE section of the command
	// buf[3] is the file of address1, FILE1 = 2
	// buf[4] is the rank of address1, RANK1 = 3
	// buf[5] is the file of address2, FILE2 = 4
	// buf[6] is the rank of address1, RANK2 = 5
	command := MOVEGRAMMAR.FindAllStringSubmatch(s, CMD)[0] // needed to pass in 1, and CMD=1

	cmd := command[CMD]
	piece := command[CHESSPIECE]
	rank1 := command[RANK1]
	file1 := command[FILE1]

	rank2 := command[RANK2]
	file2 := command[FILE2]

	tokens := []Token{
		token{COMMAND, cmd},
		token{PIECE, (piece + "-" + file1 + rank1)},
		token{ADDRESS, (file2 + rank2)},
	}

	return tokens, nil
}
