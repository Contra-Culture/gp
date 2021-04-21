package main

import (
	"github.com/Contra-Culture/gp"
)

func main() {
	objOpening := gp.New("object literal opening", gp.ExactTokenParser("litDelim", "{"))
	objClosing := gp.New("object literal closing", gp.ExactTokenParser("litDelim", "}"))
	arrayOpening := gp.New("array literal opening", gp.ExactTokenParser("litDelim", "["))
	arrayClosing := gp.New("array literal closing", gp.ExactTokenParser("litDelim", "]"))
	strOpeningClosing := gp.New("string literal opening/closing", gp.ExactTokenParser("litDelim", "\""))
	escaping := gp.New("escape symbol", gp.ExactTokenParser("escape", "\\"))
	numberParser, err := gp.PatternTokenParser("number", "^\\d+$")
	if err != nil {
		panic(err)
	}
	numberLiteral := gp.New("number literal", numberParser)
	stringLiteral := gp.Seq("string literal", strOpeningClosing, strOpeningClosing)
	boolLiteral := gp.New()
	arrayLiteral := gp.New()
	objectLiteral := gp.New()
}
