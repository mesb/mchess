/*
   Packages pieces houses the essential mechanisms for chess pieces

   Each piece has a type, a location on the board, a set of rules, a status,..
*/
package pieces

import (
	"fmt"
	"github.com/mesb/mchess/address"
)

// There are 6 types of pieces. In our chess interpreter, tokens will
//  be the pieces and are typified by their natural types
// store type of a piece or token as a type Typ
type Typ uint

// now initialize the various types of pieces
const (
	TypeError Typ = iota
	King
	Queen
	Knight
	Bishop
	Rook
	Pawn
	Empty
)

// A string representation for the types of pieces
func (t Typ) String() string {
	switch t {
	case King:
		return fmt.Sprintf("%s", "K")
	case Queen:
		return fmt.Sprintf("%s", "Q")
	case Knight:
		return fmt.Sprintf("%s", "KT")
	case Bishop:
		return fmt.Sprintf("%s", "B")
	case Rook:
		return fmt.Sprintf("%s", "R")
	case Pawn:
		return fmt.Sprintf("%s", "P")
	case Empty:
		return fmt.Sprintf("%s", "E")
	default:
		return fmt.Sprintf("Unrecognized type")
	}
}

// a piece is any resource that has the following behavior
type Piece interface {
	Type() Typ
	Address() address.Addr
	Camp() string
	Status() bool
	String() string
}

// concrete piece struct
type piece struct {
	typ    Typ
	addr   address.Addr
	camp   string
	status bool
}

// string representation of a piece
func (p piece) String() string {
	return fmt.Sprintf("(%s-%v) ", p.Type(), p.Address())
}

// construct a new piece
func NewPiece(typ Typ, a address.Addr, c string, s bool) Piece {
	p := piece{typ, a, c, s}

	return p
}

// get the type of a piece
func (p piece) Type() Typ {
	return p.typ
}

// get the address of a piece
func (p piece) Address() address.Addr {
	return p.addr
}

// get the camp
func (p piece) Camp() string {
	return p.camp
}

// get the status
func (p piece) Status() bool {
	return p.status
}

// the status and the address change
// the type may also change when a piece is promoted
func (p *piece) WriteAddr(a address.Addr) {
	p.addr = a
}

// update status
func (p *piece) WriteStatus(s bool) {
	p.status = s
}
