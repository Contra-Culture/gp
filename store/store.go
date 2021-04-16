package store

import (
	"fmt"
	"io"
)

type Symbol struct {
	Rune     rune
	Size     int
	Line     int
	Position int
}

type SymbolsStore struct {
	symbolsIndex []Symbol
	linesIndex   []int
}

const NewLine = '\n'

func New(source io.RuneReader) (store *SymbolsStore, err error) {
	store = &SymbolsStore{
		symbolsIndex: []Symbol{},
		linesIndex:   []int{0},
	}
	var (
		r        rune
		i        int
		size     int
		symbol   Symbol
		line     = 1
		position = 1
	)
	for i = 0; true; i++ {
		r, size, err = source.ReadRune()
		if err != nil {
			if err == io.EOF {
				err = nil
				return
			}
			return
		}
		symbol = Symbol{
			Rune:     r,
			Size:     size,
			Line:     line,
			Position: position,
		}
		store.symbolsIndex = append(store.symbolsIndex, symbol)
		if r == NewLine {
			line++
			store.linesIndex = append(store.linesIndex, i+1)
			position = 1
		} else {
			position++
		}
	}
	return
}

func (s *SymbolsStore) GetSymbol(idx int) (symbol Symbol, err error) {
	symbolsQnt := len(s.symbolsIndex)
	if idx >= symbolsQnt {
		err = fmt.Errorf("%d index is out of symbols range (lenght: %d)", idx, symbolsQnt)
		return
	}
	symbol = s.symbolsIndex[idx]
	return
}
func (s *SymbolsStore) GetLine(idx int) (symbols []Symbol, err error) {
	linesQnt := len(s.linesIndex)
	if idx >= linesQnt {
		err = fmt.Errorf("line index %d is out of lines range (length: %d)", idx, linesQnt)
		return
	}
	line := idx + 1
	lineStart := s.linesIndex[idx]
	var symbol Symbol
	for i := lineStart; true; i++ {
		symbolsQnt := len(s.symbolsIndex)
		if i >= symbolsQnt {
			return
		}
		symbol = s.symbolsIndex[i]
		if symbol.Line != line {
			break
		}
		symbols = append(symbols, symbol)
	}
	return
}
func (s *SymbolsStore) GetLineBySymbolIndex(idx int) (ln int, symbols []Symbol, err error) {
	symbol, err := s.GetSymbol(idx)
	if err != nil {
		return
	}
	symbols, err = s.GetLine(symbol.Line - 1)
	if err != nil {
		return
	}
	ln = symbol.Line
	return
}

func (s *SymbolsStore) LineIndex() []int {
	return s.linesIndex
}

