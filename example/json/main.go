package main

import (
	"github.com/Contra-Culture/gp"
)

func main() {
	literal := gp.Token("literal")
	nullNode := gp.BasicNode("null literal", gp.ExactParserMaker(literal, "null"))
	numberNode := gp.BasicNode("number", gp.ContinuousParserMaker(literal))
	boolNode := gp.OrNode("bool",
		gp.BasicNode("true", gp.ExactParserMaker(literal, "true")),
		gp.BasicNode("false", gp.ExactParserMaker(literal, "false")),
	)
	stringNode := gp.SeqNode("string")
	gp.BeginNode(gp.OrNode("value", nullNode, numberNode, boolNode))
}
