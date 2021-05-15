package reader

import (
	symbolsStore "github.com/Contra-Culture/gp/store"
)

type Reader struct {
	beginIdx int
	cursor   int
	store    *symbolsStore.SymbolsStore
	parent   *Reader
}

func New(store *symbolsStore.SymbolsStore, beginIdx int) *Reader {
	return &Reader{
		beginIdx: beginIdx,
		cursor:   beginIdx,
		store:    store,
	}
}
func (sr *Reader) ReadSymbol() (s symbolsStore.Symbol, err error) {
	s, err = sr.store.GetSymbol(sr.cursor)
	if err != nil {
		return
	}
	sr.cursor++
	return
}
func (sr *Reader) Frame() []symbolsStore.Symbol {
	symbols, _ := sr.store.GetRange(sr.beginIdx, sr.cursor-1)
	return symbols
}
func (sr *Reader) Continuation() *Reader {
	return &Reader{
		store:    sr.store,
		beginIdx: sr.cursor,
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
