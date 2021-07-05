package ast

type (
	identFPM struct {
		typ   string
		conts []FragmentParserMaker
	}
	identFP struct {
		typ     string
		conts   []FragmentParserMaker
		context *Node
	}
)

const (
	IDENT_TOKEN = "ident"
)

func IDENT(typ string, conts ...FragmentParserMaker) FragmentParserMaker {
	return &identFPM{
		typ,
		conts,
	}
}
func (fpm *identFPM) Make(context *Node) FragmentParser {
	typ := fpm.typ
	conts := fpm.conts
	return &identFP{
		typ,
		conts,
		context,
	}
}
func (fp *identFP) Parse(scanner Scanner) (node *Node, err error) {
	var sr rune
	var runes []rune
	sr, err = scanner.Read()
	if err != nil {
		return
	}
	runes = append(runes, sr)
	for i := 0; true; i++ {
		for _, c := range fp.conts {
			cp := c.Make(fp.context)
			node, err = cp.Parse(scanner)
			if err != nil {
				return
			}
			if node != nil {
				break
			}
		}
		sr, err = scanner.Read()
		if err != nil {
			return
		}
		runes = append(runes, sr)
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
