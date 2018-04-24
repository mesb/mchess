/*
   Package board implements the rudiments of a chess board
*/
package board

import (
	"fmt"
	"github.com/mesb/mchess/address"
	"github.com/mesb/mchess/pieces"
)

// the size of the board
const BOARDSIZE = 64
const BLOCKSIZE = 8 // blocks represent rows and colums

// the uses of ranks and files appear often, so we need constants
const (
	FIRSTRANK = iota
	SECONDRANK
	THIRDRANK
	FOURTHRANK
	FIFTHRANK
	SIXTHRANK
	SEVENTHRANK
	EIGTHRANK
)

// default first rank of every player
// we can simply use this to populate the first rank of all players
var DEFAULTFIRSTRANK = [BLOCKSIZE]pieces.Typ{pieces.Rook, pieces.Knight, pieces.Bishop, pieces.Queen, pieces.King, pieces.Bishop, pieces.Knight, pieces.Rook}

// a board is represented as an array of chess board as an array
type Board struct {
	board [BOARDSIZE]pieces.Piece
}

/*
func (b Board) Content() [BOARDSIZE]pieces.Piece {
	return b.board
}
*/

// Initialize board
func InitBoard() *Board {
	// create board object
	empty := pieces.NewPiece(pieces.Empty, address.MakeAddr(FIRSTRANK, address.File(0)), "Nothing", true)
	b := Board{[BOARDSIZE]pieces.Piece{empty}}

	for k := 0; k < BOARDSIZE; k++ {
		b.board[k] = empty
	}

	// fill in white's player
	for i := 0; i < BLOCKSIZE; i++ {

		// add the first rank of player 1
		addr := address.MakeAddr(address.Rank(FIRSTRANK), address.File(i))
		p := pieces.NewPiece(DEFAULTFIRSTRANK[i], addr, "Player 1", true)
		b.board[addr.Index()] = p

		// add the pawns of player 1
		addr = address.MakeAddr(address.Rank(SECONDRANK), address.File(i))
		p = pieces.NewPiece(pieces.Pawn, addr, "Player 1", true)
		b.board[addr.Index()] = p

		// add the ranks of player 2
		addr = address.MakeAddr(address.Rank(EIGTHRANK), address.File(i))
		p = pieces.NewPiece(DEFAULTFIRSTRANK[i], addr, "Player 2", true)
		b.board[addr.Index()] = p //

		// add the pawns of player 2
		addr = address.MakeAddr(address.Rank(SEVENTHRANK), address.File(i))
		p = pieces.NewPiece(pieces.Pawn, addr, "Player 2", true)
		b.board[addr.Index()] = p
	}

	// return board
	return &b
}

// print a board
func Printer(b Board) {
	for i := 1; i < BOARDSIZE+1; i++ {
		j := i - 1
		fmt.Printf("%v", b.board[j])

		if (i > 0) && (i%8 == 0) {
			fmt.Println()
		}
	}
}
