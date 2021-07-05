package main

import (
	"github.com/Contra-Culture/gp"
	t "github.com/Contra-Culture/gp/tester"
)

func main() {
	separator := gp.Token("sep")
	doubleDotParser := gp.ExactParserMaker(separator, ":")
	literal := gp.Token("literal")
	nullNode := gp.BasicNode("null literal", gp.ExactParserMaker(literal, "null"))
	numberNode := gp.BasicNode("number", gp.ContinuousParserMaker(literal))
	boolNode := gp.OrNode("bool",
		gp.BasicNode("true", gp.ExactParserMaker(literal, "true")),
		gp.BasicNode("false", gp.ExactParserMaker(literal, "false")),
	)
	stringNode := gp.BasicNode("string literal",
		gp.ContinuousParserMaker(literal,
			t.DoubleQuote(),
			t.PrintableWithExtended(),
			t.DoubleQuote(),
		),
	)
	literalNode := gp.OrNode("literal", nullNode, numberNode, boolNode, stringNode)
	keyNode := stringNode
	kvNode := gp.SeqNode("key-value", keyNode, gp.BasicNode("objseparator"))
	gp.BeginNode(gp.OrNode("value", nullNode, numberNode, boolNode, stringNode))
}
