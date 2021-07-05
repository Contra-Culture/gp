package ast

type (
	Pos struct {
		InStream, Line, InLine int
	}
	Range struct {
		Start, End Pos
		RawParsed  []rune
		Parsed     []rune
	}
	Meaning struct {
		Token       string
		Description string
	}
	Namespace struct {
		typ    string
		names  []string
		parent *Namespace
	}
	Node struct {
		Meaning
		Range
		Namespace
		Children []Node
	}
	Scanner interface {
		Read() (rune, error)
		Unread() error
	}
	FragmentParserMaker interface {
		Make(*Node) FragmentParser
	}
	FragmentParser interface {
		Parse(Scanner) (*Node, error)
	}
	Grammar struct {
		FragmentParserMakers []FragmentParserMaker
	}
	Parser interface {
		SyntaxTree(Scanner) (*Node, error)
	}
	Mode       string
	ModeParams []string
)

const (
	NOTOKEN  = ""
	NOLIMIT  = 0
	UNDEFPOS = 0
)

func Beginning(conts ...FragmentParserMaker) Parser {
	return &Grammar{FragmentParserMakers: conts}
}
func (g *Grammar) SyntaxTree(scanner Scanner) (node *Node, err error) {
	node = &Node{
		Meaning{
			Token:       NOTOKEN,
			Description: "beginning",
		},
		Range{
			Start: Pos{
				InStream: 1,
				Line:     1,
				InLine:   1,
			},
			End: Pos{
				InStream: UNDEFPOS,
				Line:     UNDEFPOS,
				InLine:   UNDEFPOS,
			},
		},
		Namespace{},
		[]Node{},
	}
	var child *Node
	for _, fpm := range g.FragmentParserMakers {
		fp := fpm.Make(node)
		child, err = fp.Parse(scanner)
		if err != nil {
			return
		}
		if child != nil {
			break
		}
	}
	node.Children = append(node.Children, *child)
	return
}
