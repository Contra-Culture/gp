package store

import (
	"fmt"
	"io"
	"strings"

	"github.com/Contra-Culture/gp/lsrange"
)

type (
	Symbols []Symbol
	Symbol  struct {
		Rune     rune
		Size     int
		Line     int
		Position int
	}
	SymbolsStore struct {
		symbolsIndex Symbols
		linesIndex   []int
	}
)

const newLINE = '\n'

func (ss Symbols) String() string {
	var sb strings.Builder
	for _, s := range ss {
		sb.WriteRune(s.Rune)
	}
	return sb.String()
}
func (ss Symbols) Runes() (runes []rune) {
	for _, s := range ss {
		runes = append(runes, s.Rune)
	}
	return
}

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
		if r == newLINE {
			line++
			store.linesIndex = append(store.linesIndex, i+1)
			position = 1
		} else {
			position++
		}
	}
	return
}
func (s *SymbolsStore) GetRange(start int, end int) (symbols Symbols, err error) {
	if start > end {
		err = fmt.Errorf("start index (given: %d) should be less than the end index (given: %d)", start, end)
		return
	}
	symbolsQnt := len(s.symbolsIndex)
	if start >= symbolsQnt {
		err = fmt.Errorf("%d index is out of symbols range (lenght: %d)", start, symbolsQnt)
		return
	}
	if end >= symbolsQnt {
		err = fmt.Errorf("%d index is out of symbols range (lenght: %d)", end, symbolsQnt)
		return
	}
	for i := start; i <= end; i++ {
		symbols = append(symbols, s.symbolsIndex[i])
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
func (s *SymbolsStore) GetLine(idx int) (symbols Symbols, err error) {
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
func (ss Symbols) First() Symbol {
	return ss[0]
}
func (ss Symbols) Last() Symbol {
	return ss[len(ss)-1]
}
func (ss Symbols) Symbols() Symbols {
	return ss
}
func (ss Symbols) Lines() lsrange.LinesRange {
	return lsrange.LinesRange{
		First: ss.First().Line,
		Last:  ss.Last().Line,
	}
}
