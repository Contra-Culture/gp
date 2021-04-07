package iterator

import (
	store "github.com/Contra-Culture/gp/store"
)

type SymbolsIterator struct {
	store    *store.SymbolsStore
	beginIdx int
	cursor   int
}

func New(store *store.SymbolsStore, beginIdx int) *SymbolsIterator {
	return &SymbolsIterator{
		store:    store,
		beginIdx: beginIdx,
		cursor:   beginIdx,
	}
}
func (iter *SymbolsIterator) Next() (s store.Symbol, err error) {
	s, err = iter.store.GetSymbol(iter.cursor)
	if err != nil {
		return
	}
	iter.cursor++
	return
}
func (iter *SymbolsIterator) Fork() *SymbolsIterator {
	return &SymbolsIterator{
		store:    iter.store,
		beginIdx: iter.cursor,
		cursor:   iter.cursor,
	}
}
