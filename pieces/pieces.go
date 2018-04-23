/*
   Packages pieces houses the essential mechanisms for chess pieces

   Each piece has a type, a location on the board, a set of rules, a status,..
*/
package pieces

type rank uint
type file uint

// the address of a piece
type addr struct {
	rank
	file
}

type Typ uint

// the size of the board
const BoardSize = 64

const (
	TypeError typ = iota
	King
	Queen
	Knight
	Bishop
	Rook
	Pawn
)

// a piece is any resource that has the following behavior
type Piece interface {
	Type() Typ
	Address() addr
	Camp() string
	Status() bool
}

// a board is represented as an array of chess board as an array
type Board struct {
	board [BoardSize]Piece
}
