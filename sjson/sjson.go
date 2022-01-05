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
	hexDigit = gp.AnyOneOfRunes("hex", '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f')

	whitespace = gp.Repeat("whitespace", gp.AnyOneOfRunes("single-whitespace", ' ', '\n', '\t'))
	JSONNull   = gp.String("null")
	JSONBool   = gp.Variant("bool", gp.String("true"), gp.String("false"))
	JSONArray  = gp.Seq(
		"array",
		gp.Symbol("opening bracket", '['),
		whitespace,
		gp.Optional(
			gp.Seq(
				"",
				JSONValue,
				gp.Repeat(
					"",
					gp.Symbol("", ','),
					JSONValue,
				),
			),
		),
		gp.Symbol("closing bracket", ']'),
	)
	JSONString = gp.Seq(
		"string",
		gp.Symbol("opening quote", '"'),
		gp.Variant(
			"",
			gp.Seq(
				"escape",
				gp.Symbol("", '\\'),
				gp.Variant(
					"",
					gp.AnyOneOfRunes("", '"', '\\', '/', 'b', 'f', 'n', 'r', 't'),
					gp.Seq(
						"",
						gp.Symbol("", 'u'),
						hexDigit,
						hexDigit,
						hexDigit,
						hexDigit,
					),
				),
			),
			gp.RuneExcept('"', '\\'),
		),
		gp.Symbol("closing quote", '"'),
	)
	JSONObject = gp.Seq(
		"string",
		gp.Symbol("opening bracket", '{'),
		whitespace,
		gp.Optional(
			gp.Seq(
				"",
				whitespace,
				JSONString,
				whitespace,
				gp.Symbol("", ':'),
				JSONValue,
				gp.Repeat(
					"",
					gp.Seq(
						gp.Symbol("", ','),
						JSONString,
						whitespace,
						gp.Symbol("", ':'),
						JSONValue,
					),
				),
			),
		),
		gp.Symbol("closing bracket", '}'),
	)
	JSONValue = gp.Seq(
		"value",
		whitespace,
		gp.Variant(
			"literal",
			JSONString,
			JSONNumber,
			JSONObject,
			JSONArray,
			JSONBool,
			JSONNull,
		),
		whitespace,
	)
	numMinusSign = gp.Optional(gp.Symbol("num-sign", '-'))
	numSign      = gp.Optional(gp.AnyOneOfRunes("num-sign", '+', '-'))
	zero         = gp.Symbol("digit", '0')
	nonZeroDigit = gp.AnyOneOfRunes("digit", '1', '2', '3', '4', '5', '6', '7', '8', '9')
	digit        = gp.AnyOneOfRunes("digit", '0', '1', '2', '3', '4', '5', '6', '7', '8', '9')
	JSONNumber   = gp.Seq(
		"number",
		numMinusSign,
		gp.Variant(
			"",
			gp.Seq(
				"",
				nonZeroDigit,
				gp.Repeat("", digit),
			),
			zero,
		),
		gp.Optional(
			gp.Seq(
				"",
				gp.Symbol("dot", '.'),
				digit,
				gp.Repeat("", digit),
			),
		),
		gp.Optional(
			gp.Seq(
				"mantissa",
				gp.AnyOneOfRunes("", 'e', 'E'),
				numSign,
				digit,
				gp.Repeat("", digit),
			),
		),
	)
)

func New() {

}
