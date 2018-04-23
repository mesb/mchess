package main

import (
	"fmt"
)

type rank uint
type file uint

// the address of a piece
type Addr struct {
	rank
	file
}

func (a Addr) String() string {
	f := (int(a.file) + 97)
	return fmt.Sprintf("%c%v", f, a.rank)
}

func MakeAddr(r rank, f file) Addr {
	a := Addr{r, f}
	return a
}

type Typ uint

// the size of the board
const BOARDSIZE = 64
const BLOCKSIZE = 8 // blocks represent rows and colums
const FIRSTRANK = 0
const LASTRANK = 7

// default first rank of every player
var RANK1 = [BLOCKSIZE]Typ{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}

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
	Address() Addr
	Camp() string
	Status() bool
	String() string
}

type piece struct {
	typ    Typ
	addr   Addr
	camp   string
	status bool
}

func (p piece) String() string {
	return fmt.Sprintf("(%s-%v) ", p.Type(), p.Address())
}

func NewPiece(typ Typ, a Addr, c string, s bool) Piece {
	p := piece{typ, a, c, s}

	return p
}

// get the type of a piece
func (p piece) Type() Typ {
	return p.typ
}

// get the address of a piece
func (p piece) Address() Addr {
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
func (p *piece) WriteAddr(a Addr) {
	p.addr = a
}

func (p *piece) WriteStatus(s bool) {
	p.status = s
}

// a board is represented as an array of chess board as an array
type Board struct {
	board [BOARDSIZE]Piece
}

func InitBoard() *Board {
	// create board object
	empty := NewPiece(Empty, MakeAddr(FIRSTRANK, file(0)), "Nothing", true)
	b := Board{[BOARDSIZE]Piece{empty}}

	for k := 0; k < BOARDSIZE; k++ {
		b.board[k] = empty
	}

	rank8 := (7 * 8)
	rank7 := (7 * 7) - 1

	// fill in white's player
	for i := 0; i < BLOCKSIZE; i++ {

		// add the first rank of player 1
		addr := MakeAddr(rank(FIRSTRANK), file(i))
		p := NewPiece(RANK1[i], addr, "Player 1", true)
		b.board[i] = p

		// add the pawns of player 1
		addr = MakeAddr(rank(FIRSTRANK+1), file(i))
		p = NewPiece(Pawn, addr, "Player 1", true)
		b.board[i+BLOCKSIZE] = p

		// add the ranks of player 2
		addr = MakeAddr(rank(LASTRANK), file(i))
		p = NewPiece(RANK1[i], addr, "Player 2", true)
		b.board[rank8+i] = p //

		// add the pawns of player 2
		addr = MakeAddr(rank(LASTRANK-1), file(i))
		p = NewPiece(Pawn, addr, "Player 2", true)
		b.board[rank7+i] = p
	}

	// return board
	return &b
}

func printer(b Board) {
	for i := 1; i < BOARDSIZE+1; i++ {
		j := i - 1
		fmt.Printf("%v", b.board[j])

		if (i > 0) && (i%8 == 0) {
			fmt.Println()
		}
	}
}

func main() {
	b := InitBoard()
	printer(*b)
	fmt.Println(len(b.board))
}
