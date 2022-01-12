package json

import "github.com/Contra-Culture/gp"

func New() (*gp.Parser, error) {
	return gp.New(func(c *gp.SyntaxCfgr) {
		c.TopRule("value")
		c.Rule("value", gp.Seq(
			gp.Use("ws"),
			gp.Variant(
				gp.Use("string"),
				gp.Use("number"),
				gp.Use("array"),
				gp.Use("object"),
				gp.Use("bool"),
				gp.Use("null"),
			),
			gp.Use("ws"),
		))
		c.Rule("hex-digit", gp.AnyOfRunes('0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'))
		c.Rule("ws", gp.Repeat(gp.AnyOfRunes(' ', '\n', '\t')))
		c.Rule("null", gp.String("null"))
		c.Rule("bool", gp.Variant(gp.String("true"), gp.String("false")))
		c.Rule("array", gp.Seq(gp.Symbol('['), gp.Use("ws"), gp.Optional(gp.Seq(gp.Use("value"), gp.Repeat(gp.Seq(gp.Symbol(','), gp.Use("value"))))), gp.Symbol(']')))
		c.Rule("string", gp.Seq(
			gp.Symbol('"'),
			gp.Variant(
				gp.Seq(
					gp.Symbol('\\'),
					gp.Variant(
						gp.AnyOfRunes('"', '\\', '/', 'b', 'f', 'n', 'r', 't'),
						gp.Seq(gp.Symbol('u'), gp.Use("hex-digit"), gp.Use("hex-digit"), gp.Use("hex-digit"), gp.Use("hex-digit")),
					),
				),
				gp.RuneExcept('"', '\\'),
			),
			gp.Symbol('"'),
		))
		c.Rule("object", gp.Seq(
			gp.Symbol('{'), gp.Use("ws"),
			gp.Optional(
				gp.Seq(
					gp.Use("ws"), gp.Use("string"), gp.Use("ws"), gp.Symbol(':'), gp.Use("value"),
					gp.Repeat(gp.Seq(gp.Symbol(','), gp.Use("string"), gp.Use("ws"), gp.Symbol(':'), gp.Use("value"))),
				),
			),
			gp.Symbol('}'),
		))
		c.Rule("digit", gp.AnyOfRunes('0', '1', '2', '3', '4', '5', '6', '7', '8', '9'))
		c.Rule("number", gp.Seq(
			gp.Optional(gp.Symbol('-')),
			gp.Variant(
				gp.Seq(gp.AnyOfRunes('1', '2', '3', '4', '5', '6', '7', '8', '9'), gp.Repeat(gp.Use("digit"))),
				gp.Symbol('0'),
			),
			gp.Optional(gp.Seq(gp.Symbol('.'), gp.Use("digit"), gp.Repeat(gp.Use("digit")))),
			gp.Optional(gp.Seq(gp.AnyOfRunes('e', 'E'), gp.Optional(gp.AnyOfRunes('+', '-')), gp.Use("digit"), gp.Repeat(gp.Use("digit")))),
		))
	})
}
