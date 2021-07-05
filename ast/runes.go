package ast

import "fmt"

type (
	runesFPM struct {
		token string
		runes []rune
		mode  RunesFpMode
		conts []FragmentParserMaker
	}
	runesFP struct {
		token   string
		runes   []rune
		mode    RunesFpMode
		conts   []FragmentParserMaker
		context *Node
	}
	RunesFpMode int
)

const (
	SEPARATOR_TOKEN = "separator"
	OPENING_TOKEN   = "opening"
	CLOSING_TOKEN   = "closing"
)

const (
	NOT_GREEDY = RunesFpMode(iota)
	GREEDY_SAME
	GREEDY_ANY
)

var modeNames = []string{
	"NOT_GREEDY",
	"GREEDY_SAME",
	"GREED_ANY",
}

func RUNES(runes []rune, token string, mode RunesFpMode, conts ...FragmentParserMaker) FragmentParserMaker {
	return &runesFPM{
		token,
		runes,
		mode,
		conts,
	}
}
func (fpm *runesFPM) Make(context *Node) FragmentParser {
	token := fpm.token
	runes := fpm.runes
	conts := fpm.conts
	mode := fpm.mode
	return &runesFP{
		token,
		runes,
		mode,
		conts,
		context,
	}
}
func (fp *runesFP) Parse(scanner Scanner) (node *Node, err error) {
	var rr rune
	runes := []rune{}
	switch fp.mode {
	case NOT_GREEDY:
		rr, err = scanner.Read()
		if err != nil {
			return
		}
		for _, r := range fp.runes {
			if r == rr {
				runes = append(runes, rr)
				break
			} else {
				err = scanner.Unread()
				if err != nil {
					return
				}
			}
		}
	case GREEDY_SAME:
		rr, err = scanner.Read()
		if err != nil {
			return
		}
	outer_same:
		for _, r := range fp.runes {
			if r == rr {
				runes = append(runes, rr)
				for {
					rr, err = scanner.Read()
					if err != nil {
						return
					}
					if r == rr {
						runes = append(runes, rr)
					} else {
						err = scanner.Unread()
						if err != nil {
							return
						}
						break outer_same
					}
				}
			} else {
				err = scanner.Unread()
				if err != nil {
					return
				}
			}
		}
	case GREEDY_ANY:
		rr, err = scanner.Read()
		if err != nil {
			return
		}
	outer_any:
		for {
			l := len(runes)
			for _, r := range fp.runes {
				if r == rr {
					runes = append(runes, rr)
				} else {
					err = scanner.Unread()
					if err != nil {
						return
					}
				}
			}
			if l <= len(runes) {
				break outer_any
			}
		}
	default:
		err = fmt.Errorf("wrong mode: `%d`, expected: `0` for %s, `1` for %s or `2` for %s", int(fp.mode), modeNames[0], modeNames[1], modeNames[2])
		return
	}
	if len(runes) == 0 {
		return
	}
	node = &Node{
		Meaning{
			Token: KEYWORD_TOKEN,
		},
		Range{
			Start: Pos{
				InStream: fp.context.Start.InStream + 1,
				Line:     fp.context.Start.Line,
				InLine:   fp.context.Start.InLine + 1,
			},
			End: Pos{
				InStream: fp.context.Start.InStream + len(runes),
				Line:     fp.context.Start.Line,
				InLine:   fp.context.Start.InLine + len(runes),
			},
			RawParsed: runes,
			Parsed:    runes,
		},
		Namespace{},
		[]Node{},
	}
	var cfp FragmentParser
	var child *Node
	for _, c := range fp.conts {
		cfp = c.Make(node)
		child, err = cfp.Parse(scanner)
		if err != nil {
			return
		}
		if child != nil {
			node.Children = append(node.Children, *child)
			break
		}
	}
	return
}
