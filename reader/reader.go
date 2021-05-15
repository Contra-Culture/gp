package reader

import (
	"fmt"

	symbolsStore "github.com/Contra-Culture/gp/store"
)

type BaseSymbolReader struct {
	beginIdx int
	cursor   int
	store    *symbolsStore.SymbolsStore
	parent   *BaseSymbolReader
}

func New(store *symbolsStore.SymbolsStore, beginIdx int) *BaseSymbolReader {
	return &BaseSymbolReader{
		beginIdx: beginIdx,
		cursor:   beginIdx,
		store:    store,
	}
}
func (sr *BaseSymbolReader) ReadSymbol() (s symbolsStore.Symbol, err error) {
	s, err = sr.store.GetSymbol(sr.cursor)
	if err != nil {
		return
	}
	sr.cursor++
	return
}
func (sr *BaseSymbolReader) Frame() []symbolsStore.Symbol {
	symbols, _ := sr.store.GetRange(sr.beginIdx, sr.cursor-1)
	return symbols
}
func (sr *BaseSymbolReader) Continuation() *BaseSymbolReader {
	return &BaseSymbolReader{
		store:    sr.store,
		beginIdx: sr.cursor,
		cursor:   sr.cursor,
		parent:   sr,
	}
}
func (sr *BaseSymbolReader) ReadRune() (r rune, size int, err error) {
	var s symbolsStore.Symbol
	s, err = sr.ReadSymbol()
	if err != nil {
		fmt.Printf("\n\t\terr: %s", err.Error())
		return
	}
	r = s.Rune
	size = s.Size
	fmt.Printf("\n\t\treader.ReadRune() symbol: `%s` -> %#v\n", string(s.Rune), s)
	return
}
func (sr *BaseSymbolReader) CommitToParent() {
	sr.parent.cursor = sr.cursor
}
