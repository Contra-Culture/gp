package sjson

import "github.com/Contra-Culture/gp"

var (
	pattern = gp.Repeat(
		"value",
		gp.Variant(
			"JSON pattern",
			JSONString,
			JSONNumber,
			JSONArray,
			JSONObject,
			JSONBool,
			JSONNull,
		),
	)
	JSONNull  = gp.String("null")
	JSONBool  = gp.Variant("bool", gp.String("true"), gp.String("false"))
	JSONArray = gp.Seq(
		"array",
		gp.Symbol("opening bracket", '['),
		gp.Symbol("closing bracket", ']'),
	)
	JSONString = gp.Seq(
		"string",
		gp.Symbol("opening quote", '"'),
		gp.Symbol("closing quote", '"'),
	)
	JSONObject = gp.Seq(
		"string",
		gp.Symbol("opening bracket", '{'),
		gp.Symbol("closing bracket", '}'),
	)
	JSONNumber = gp.Seq(
		"number",
		gp.Optional(gp.Symbol("minus", '-')),
		gp.Variant(
			"int",
			gp.Seq(
				"zero-beginning number",
				gp.Symbol("zero", '0'),
				gp.Optional(
					gp.Seq(
						"floating point part",
						gp.Symbol("floating point", '.'),
						gp.Repeat("fraction", gp.Digit()),
					),
				),
				gp.Variant(
					"non-zero",
					gp.Symbol("1", '1'),
					gp.Symbol("2", '2'),
					gp.Symbol("3", '3'),
					gp.Symbol("4", '4'),
					gp.Symbol("5", '5'),
					gp.Symbol("6", '6'),
					gp.Symbol("7", '7'),
					gp.Symbol("8", '8'),
					gp.Symbol("9", '9'),
				),
			),
		),
	)
)

func New() {

}
