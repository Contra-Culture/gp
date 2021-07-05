package ast

type (
	keywordFPM struct {
		keyword string
		conts   []FragmentParserMaker
	}
	keywordFP struct {
		keyword string
		conts   []FragmentParserMaker
		context *Node
	}
)

const (
	KEYWORD_TOKEN = "keyword"
)

func KEYWORD(keyword string, conts ...FragmentParserMaker) FragmentParserMaker {
	return &keywordFPM{
		keyword,
		conts,
	}
}
func (fpm *keywordFPM) Make(context *Node) FragmentParser {
	keyword := fpm.keyword
	conts := fpm.conts
	return &keywordFP{
		keyword,
		conts,
		context,
	}
}
func (fp *keywordFP) Parse(scanner Scanner) (node *Node, err error) {
	var sr rune
	var runes []rune
	for i, r := range fp.keyword {
		sr, err = scanner.Read()
		if err != nil {
			return
		}
		if sr != r {
			for ; i >= 0; i-- {
				err = scanner.Unread()
				if err != nil {
					return
				}
			}
		}
	}
	parsed := []rune(fp.keyword)
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
				InStream: fp.context.Start.InStream + len(fp.keyword),
				Line:     fp.context.Start.Line,
				InLine:   fp.context.Start.InLine + len(fp.keyword),
			},
			RawParsed: parsed,
			Parsed:    parsed,
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
