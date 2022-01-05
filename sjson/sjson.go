package sjson

import "github.com/Contra-Culture/gp"

var (
	pattern = gp.Repeat(
		gp.Variant(
			JSONString,
			JSONNumber,
			JSONArray,
			JSONObject,
			JSONBool,
			JSONNull,
		),
	)
	hexDigit = gp.AnyOneOfRunes('0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f')

	whitespace = gp.Repeat(gp.AnyOneOfRunes(' ', '\n', '\t'))
	JSONNull   = gp.String("null")
	JSONBool   = gp.Variant(gp.String("true"), gp.String("false"))
	JSONArray  = gp.Seq(
		gp.Symbol('['),
		whitespace,
		gp.Optional(
			gp.Seq(
				JSONValue,
				gp.Repeat(
					gp.Symbol(','),
					JSONValue,
				),
			),
		),
		gp.Symbol(']'),
	)
	JSONString = gp.Seq(
		gp.Symbol('"'),
		gp.Variant(
			gp.Seq(
				gp.Symbol('\\'),
				gp.Variant(
					gp.AnyOneOfRunes('"', '\\', '/', 'b', 'f', 'n', 'r', 't'),
					gp.Seq(
						gp.Symbol('u'),
						hexDigit,
						hexDigit,
						hexDigit,
						hexDigit,
					),
				),
			),
			gp.RuneExcept('"', '\\'),
		),
		gp.Symbol('"'),
	)
	JSONObject = gp.Seq(
		gp.Symbol('{'),
		whitespace,
		gp.Optional(
			gp.Seq(
				"",
				whitespace,
				JSONString,
				whitespace,
				gp.Symbol(':'),
				JSONValue,
				gp.Repeat(
					gp.Seq(
						gp.Symbol(','),
						JSONString,
						whitespace,
						gp.Symbol(':'),
						JSONValue,
					),
				),
			),
		),
		gp.Symbol('}'),
	)
	JSONValue = gp.Seq(
		whitespace,
		gp.Variant(
			JSONString,
			JSONNumber,
			JSONObject,
			JSONArray,
			JSONBool,
			JSONNull,
		),
		whitespace,
	)
	numMinusSign = gp.Optional(gp.Symbol('-'))
	numSign      = gp.Optional(gp.AnyOneOfRunes('+', '-'))
	zero         = gp.Symbol('0')
	nonZeroDigit = gp.AnyOneOfRunes('1', '2', '3', '4', '5', '6', '7', '8', '9')
	digit        = gp.AnyOneOfRunes('0', '1', '2', '3', '4', '5', '6', '7', '8', '9')
	JSONNumber   = gp.Seq(
		numMinusSign,
		gp.Variant(
			gp.Seq(
				nonZeroDigit,
				gp.Repeat(digit),
			),
			zero,
		),
		gp.Optional(
			gp.Seq(
				gp.Symbol('.'),
				digit,
				gp.Repeat(digit),
			),
		),
		gp.Optional(
			gp.Seq(
				gp.AnyOneOfRunes('e', 'E'),
				numSign,
				digit,
				gp.Repeat(digit),
			),
		),
	)
)

func New() {

}
