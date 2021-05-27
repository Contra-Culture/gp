package reader

import (
	"errors"
	"fmt"

	symbolsStore "github.com/Contra-Culture/gp/store"
)

type (
	SymbolReader interface {
		ReadSymbol() (s symbolsStore.Symbol, err error)
	}
	Reader struct {
		beginIdx int
		endIdx   int
		cursor   int
		store    *symbolsStore.SymbolsStore
		parent   *Reader
	}
)

const NOLIMIT = 0

func New(store *symbolsStore.SymbolsStore, beginIdx int) *Reader {
	return &Reader{
		beginIdx: beginIdx,
		cursor:   beginIdx,
		store:    store,
	}
}
func (sr *Reader) ReadSymbol() (s symbolsStore.Symbol, err error) {
	if sr.endIdx > 0 && sr.cursor > sr.endIdx {
		err = fmt.Errorf("index %d is out of range (limit: %d)", sr.cursor, sr.endIdx)
		return
	}
	s, err = sr.store.GetSymbol(sr.cursor)
	if err != nil {
		return
	}
	sr.cursor++
	return
}
func (st *Reader) UnreadSymbol() (err error) {
	cursor := st.cursor - 1
	if cursor < st.beginIdx {
		err = errors.New("you're trying to unread out of uncommited range")
		return
	}
	st.cursor--
	return
}
func (sr *Reader) UncommittedRead() symbolsStore.Symbols {
	symbols, _ := sr.store.GetRange(sr.beginIdx, sr.cursor-1)
	return symbols
}
func (sr *Reader) Continuation(limit int) *Reader {
	var endIdx int
	if limit > 0 {
		endIdx = sr.cursor + limit
	}
	return &Reader{
		store:    sr.store,
		beginIdx: sr.cursor,
		endIdx:   endIdx,
		cursor:   sr.cursor,
		parent:   sr,
	}
}
func (sr *Reader) ReadRune() (r rune, size int, err error) {
	var s symbolsStore.Symbol
	s, err = sr.ReadSymbol()
	if err != nil {
		return
	}
	r = s.Rune
	size = s.Size
	return
}
func (sr *Reader) CommitToParent() {
	sr.parent.cursor = sr.cursor
}
