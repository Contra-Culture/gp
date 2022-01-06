package json

import "github.com/Contra-Culture/gp"

func New() (gp.Parser, error) {
	return gp.New(func(c *gp.UnivCfgr) {
		c.Top(gp.Repeat(c.Get("value")))
		c.Define("value", gp.Seq(
			c.Get("ws"),
			gp.Variant(
				c.Get("string"),
				c.Get("number"),
				c.Get("array"),
				c.Get("object"),
				c.Get("bool"),
				c.Get("null"),
			),
			c.Get("ws"),
		))
		c.Define("hex-digit", gp.AnyOneOfRunes('0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'))
		c.Define("ws", gp.Repeat(gp.AnyOneOfRunes(' ', '\n', '\t')))
		c.Define("null", gp.String("null"))
		c.Define("bool", gp.Variant(gp.String("true"), gp.String("false")))
		c.Define("array", gp.Seq(gp.Symbol('['), c.Get("ws"), gp.Optional(gp.Seq(c.Get("value"), gp.Repeat(gp.Seq(gp.Symbol(','), c.Get("value"))))), gp.Symbol(']')))
		c.Define("string", gp.Seq(
			gp.Symbol('"'),
			gp.Variant(
				gp.Seq(
					gp.Symbol('\\'),
					gp.Variant(
						gp.AnyOneOfRunes('"', '\\', '/', 'b', 'f', 'n', 'r', 't'),
						gp.Seq(gp.Symbol('u'), c.Get("hex-digit"), c.Get("hex-digit"), c.Get("hex-digit"), c.Get("hex-digit")),
					),
				),
				gp.RuneExcept('"', '\\'),
			),
			gp.Symbol('"'),
		))
		c.Define("object", gp.Seq(
			gp.Symbol('{'), c.Get("ws"),
			gp.Optional(
				gp.Seq(
					c.Get("ws"), c.Get("string"), c.Get("ws"), gp.Symbol(':'), c.Get("value"),
					gp.Repeat(gp.Seq(gp.Symbol(','), c.Get("string"), c.Get("ws"), gp.Symbol(':'), c.Get("value"))),
				),
			),
			gp.Symbol('}'),
		))
		c.Define("digit", gp.AnyOneOfRunes('0', '1', '2', '3', '4', '5', '6', '7', '8', '9'))
		c.Define("number", gp.Seq(
			gp.Optional(gp.Symbol('-')),
			gp.Variant(
				gp.Seq(gp.AnyOneOfRunes('1', '2', '3', '4', '5', '6', '7', '8', '9'), gp.Repeat(c.Get("digit"))),
				gp.Symbol('0'),
			),
			gp.Optional(gp.Seq(gp.Symbol('.'), c.Get("digit"), gp.Repeat(c.Get("digit")))),
			gp.Optional(gp.Seq(gp.AnyOneOfRunes('e', 'E'), gp.Optional(gp.AnyOneOfRunes('+', '-')), c.Get("digit"), gp.Repeat(c.Get("digit")))),
		))
	})
}
