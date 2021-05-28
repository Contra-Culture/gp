package main

import (
	"github.com/Contra-Culture/gp"
	t "github.com/Contra-Culture/gp/tester"
)

func main() {
	literal := gp.Token("literal")
	nullNode := gp.BasicNode("null literal", gp.ExactParserMaker(literal, "null"))
	numberNode := gp.BasicNode("number", gp.ContinuousParserMaker(literal))
	boolNode := gp.OrNode("bool",
		gp.BasicNode("true", gp.ExactParserMaker(literal, "true")),
		gp.BasicNode("false", gp.ExactParserMaker(literal, "false")),
	)
	stringNode := gp.BasicNode("string", gp.ContinuousParserMaker(literal, t.DoubleQuote(), t.Any(), t.DoubleQuote()))
	gp.BeginNode(gp.OrNode("value", nullNode, numberNode, boolNode, stringNode))
}
