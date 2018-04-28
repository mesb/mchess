/*
   The mchess engine sees the chessboard as a 1 dimensional
   array of size 64. The players sees a 2D or 3D array.

   The address package handles the addressing scheme/translation
*/
package address

import (
	"fmt"
)

const (
	BLOCKSIZE = 8
)

// rank and file are tupes used for the address
type Rank uint
type File uint

// the address of a piece
type Addr struct {
	Rank
	File
}

// string representation of address
func (a Addr) String() string {
	f := (int(a.File) + 97)
	return fmt.Sprintf("%c%v::%2d", f, int(a.Rank)+1, a.Index())
}

// the pieces are stored in memory in a 1d array of size 64
// a.Index() returns the memory location or effective address
// of a piece
func (a Addr) Index() (index int) {
	r := int(a.Rank)
	f := int(a.File)

	index = (r * BLOCKSIZE) + f

	return
}

func TranslateIndex(index int) Addr {
	rank := (index / 8)
	file := (index % 8)

	return MakeAddr(Rank(rank), File(file))

}

// construct a new address from a given rank and file
func MakeAddr(r Rank, f File) Addr {
	a := Addr{r, f}
	return a
}
