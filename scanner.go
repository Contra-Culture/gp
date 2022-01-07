package gp

import (
	"errors"
	"fmt"
)

type (
	// GPRuneScanner interface is not suitable because it allows to unread only the last read rune
	// so that we need *gp.GPRuneScanner.
	GPRuneScanner struct {
		runes  []rune
		cursor int
	}
)

var ErrNoRuneToRead = errors.New("no rune to read")

func NewRuneScanner(s string) *GPRuneScanner {
	return &GPRuneScanner{
		runes:  []rune(s),
		cursor: -1,
	}
}
func (rs *GPRuneScanner) ReadOne() (r rune, cur int, err error) {
	rs.cursor++
	cur = rs.cursor
	if rs.cursor >= len(rs.runes) {
		err = ErrNoRuneToRead
		return
	}
	r = rs.runes[rs.cursor]
	return
}
func (rs *GPRuneScanner) Unread(n int) (int, error) {
	if n < 0 || rs.cursor-n < -1 {
		return rs.cursor, fmt.Errorf("can't unread %d rune(s)", n)
	}
	rs.cursor = rs.cursor - n
	return rs.cursor, nil
}
