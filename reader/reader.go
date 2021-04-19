package reader

import (
	"fmt"

	symbolsStore "github.com/Contra-Culture/gp/store"
)

type BaseSymbolReader struct {
	beginIdx int
	store    *symbolsStore.SymbolsStore
	frame    []symbolsStore.Symbol
}

func New(store *symbolsStore.SymbolsStore, beginIdx int) *BaseSymbolReader {
	return &BaseSymbolReader{
		beginIdx: beginIdx,
		store:    store,
		frame:    []symbolsStore.Symbol{},
	}
}
func (sr *BaseSymbolReader) ReadSymbol() (s symbolsStore.Symbol, err error) {
	cursor := len(sr.frame)
	s, err = sr.store.GetSymbol(cursor)
	if err != nil {
		return
	}
	sr.frame = append(sr.frame, s)
	return
}
func (sr *BaseSymbolReader) Frame() []symbolsStore.Symbol {
	return sr.frame[sr.beginIdx:]
}
func (sr *BaseSymbolReader) Continuation() *BaseSymbolReader {
	return &BaseSymbolReader{
		store:    sr.store,
		beginIdx: len(sr.frame),
		frame:    sr.frame,
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
