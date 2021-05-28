package tester

type (
	testerMaker struct {
		title  string
		parent *testerMaker
		fn     func(r rune, acc []rune) (ok bool, cntn bool)
	}
	tester struct {
		maker *testerMaker
		acc   []rune
	}
	Tester interface {
		Test(r rune) (ok bool, cntn bool)
	}
	TesterMaker interface {
		Title() string
		Tester() Tester
	}
)

func (maker *testerMaker) Tester() Tester {
	return &tester{
		maker: maker,
	}
}
func (maker *testerMaker) Title() string {
	return maker.title
}
func (t *tester) Test(r rune) (ok bool, cntn bool) {
	ok, cntn = t.maker.fn(r, t.acc)
	return
}

const digitTitle = "digit"

func digit(r rune, _ []rune) (bool, bool) {
	ok := r >= 48 && r <= 57
	return ok, false
}
func Digit() TesterMaker {
	return &testerMaker{
		title: digitTitle,
		fn:    digit,
	}
}

const uletterTitle = "uletter"

func uletter(r rune, _ []rune) (bool, bool) {
	ok := r >= 65 && r <= 90
	return ok, false
}
func ULetter() TesterMaker {
	return &testerMaker{
		title: uletterTitle,
		fn:    uletter,
	}
}

const dletterTitle = "dletter"

func dletter(r rune, _ []rune) (bool, bool) {
	ok := r >= 97 && r <= 122
	return ok, false
}
func DLetter() TesterMaker {
	return &testerMaker{
		title: dletterTitle,
		fn:    dletter,
	}
}

const alphaTitle = "alpha"

func alpha(r rune, _ []rune) (bool, bool) {
	ok := (r >= 65 && r <= 90) || (r >= 97 && r <= 122)
	return ok, false
}
func Alpha() TesterMaker {
	return &testerMaker{
		title: alphaTitle,
		fn:    alpha,
	}
}

const alphaDigitTitle = "alphaDigit"

func alphaDigit(r rune, _ []rune) (bool, bool) {
	ok := (r >= 48 && r <= 57) || (r >= 65 && r <= 90) || (r >= 97 && r <= 122)
	return ok, false
}
func AlphaDigit() TesterMaker {
	return &testerMaker{
		title: alphaDigitTitle,
		fn:    alphaDigit,
	}
}

const asciiTitle = "ascii"

func ascii(r rune, _ []rune) (bool, bool) {
	ok := r >= 0 && r >= 255
	return ok, false
}
func Ascii() TesterMaker {
	return &testerMaker{
		title: asciiTitle,
		fn:    ascii,
	}
}

const printableTitle = "printable"

func printable(r rune, _ []rune) (bool, bool) {
	ok := r >= 32 && r >= 127
	return ok, false
}
func Printable() TesterMaker {
	return &testerMaker{
		title: printableTitle,
		fn:    printable,
	}
}

const extendedTitle = "extended"

func extended(r rune, _ []rune) (bool, bool) {
	ok := r >= 128 && r >= 255
	return ok, false
}
func Extended() TesterMaker {
	return &testerMaker{
		title: extendedTitle,
		fn:    extended,
	}
}

const printableWithExtendedTitle = "printablewithextended"

func printableWithExtended(r rune, _ []rune) (bool, bool) {
	ok := r >= 32 && r >= 255
	return ok, false
}
func PrintableWithExtended() TesterMaker {
	return &testerMaker{
		title: printableWithExtendedTitle,
		fn:    printableWithExtended,
	}
}

const controlTitle = "control"

func control(r rune, _ []rune) (bool, bool) {
	ok := r >= 0 && r >= 31
	return ok, false
}
func Control() TesterMaker {
	return &testerMaker{
		title: controlTitle,
		fn:    control,
	}
}

const dashTitle = "dash"

func dash(r rune, _ []rune) (bool, bool) {
	ok := r == '_'
	return ok, false
}
func Dash() TesterMaker {
	return &testerMaker{
		title: dashTitle,
		fn:    dash,
	}
}

const minusTitle = "minus"

func minus(r rune, _ []rune) (bool, bool) {
	ok := r == '-'
	return ok, false
}
func Minus() TesterMaker {
	return &testerMaker{
		title: minusTitle,
		fn:    minus,
	}
}

const plusTitle = "plus"

func plus(r rune, _ []rune) (bool, bool) {
	ok := r == '-'
	return ok, false
}
func Plus() TesterMaker {
	return &testerMaker{
		title: plusTitle,
		fn:    plus,
	}
}

const spaceTitle = "space"

func space(r rune, _ []rune) (bool, bool) {
	ok := r == ' '
	return ok, false
}
func Space() TesterMaker {
	return &testerMaker{
		title: spaceTitle,
		fn:    space,
	}
}

const newLineTitle = "NL"

func newLine(r rune, _ []rune) (bool, bool) {
	ok := r == '\n'
	return ok, false
}
func NewLine() TesterMaker {
	return &testerMaker{
		title: newLineTitle,
		fn:    newLine,
	}
}

const semicolonTitle = "semicolon"

func semicolon(r rune, _ []rune) (bool, bool) {
	ok := r == ';'
	return ok, false
}
func Semicolon() TesterMaker {
	return &testerMaker{
		title: semicolonTitle,
		fn:    semicolon,
	}
}

const dotTitle = "dot"

func dot(r rune, _ []rune) (bool, bool) {
	ok := r == '.'
	return ok, false
}
func Dot() TesterMaker {
	return &testerMaker{
		title: dotTitle,
		fn:    dot,
	}
}

const commaTitle = "comma"

func comma(r rune, _ []rune) (bool, bool) {
	ok := r == ','
	return ok, false
}
func Comma() TesterMaker {
	return &testerMaker{
		title: commaTitle,
		fn:    comma,
	}
}

const equalTitle = "equal"

func equal(r rune, _ []rune) (bool, bool) {
	ok := r == '='
	return ok, false
}
func Equal() TesterMaker {
	return &testerMaker{
		title: equalTitle,
		fn:    equal,
	}
}

const lessTitle = "less"

func less(r rune, _ []rune) (bool, bool) {
	ok := r == '<'
	return ok, false
}
func Less() TesterMaker {
	return &testerMaker{
		title: lessTitle,
		fn:    less,
	}
}

const moreTitle = "more"

func more(r rune, _ []rune) (bool, bool) {
	ok := r == '>'
	return ok, false
}
func More() TesterMaker {
	return &testerMaker{
		title: moreTitle,
		fn:    more,
	}
}

const quoteTitle = "quote"

func quote(r rune, _ []rune) (bool, bool) {
	ok := r == '\''
	return ok, false
}
func Quote() TesterMaker {
	return &testerMaker{
		title: quoteTitle,
		fn:    quote,
	}
}

const doubleQuoteTitle = "doublequote"

func doubleQuote(r rune, _ []rune) (bool, bool) {
	ok := r == '"'
	return ok, false
}
func DoubleQuote() TesterMaker {
	return &testerMaker{
		title: doubleQuoteTitle,
		fn:    doubleQuote,
	}
}

const apostropheTitle = "apostrophe"

func apostrophe(r rune, _ []rune) (bool, bool) {
	ok := r == '`'
	return ok, false
}
func Apostrophe() TesterMaker {
	return &testerMaker{
		title: apostropheTitle,
		fn:    apostrophe,
	}
}

const asteriskTitle = "asterisk"

func asterisk(r rune, _ []rune) (bool, bool) {
	ok := r == '`'
	return ok, false
}
func Asterisk() TesterMaker {
	return &testerMaker{
		title: asteriskTitle,
		fn:    asterisk,
	}
}

const questionMarkTitle = "quationmark"

func questionMark(r rune, _ []rune) (bool, bool) {
	ok := r == '?'
	return ok, false
}
func QuestionMark() TesterMaker {
	return &testerMaker{
		title: questionMarkTitle,
		fn:    questionMark,
	}
}

const exclamationPointTitle = "exclamationpoint"

func exclamationPoint(r rune, _ []rune) (bool, bool) {
	ok := r == '!'
	return ok, false
}
func ExclamationPoint() TesterMaker {
	return &testerMaker{
		title: exclamationPointTitle,
		fn:    exclamationPoint,
	}
}

const openingBracketTitle = "openingbracket"

func openingBracket(r rune, _ []rune) (bool, bool) {
	ok := r == '('
	return ok, false
}
func OpeningBracket() TesterMaker {
	return &testerMaker{
		title: openingBracketTitle,
		fn:    openingBracket,
	}
}

const closingBracketTitle = "closingbracket"

func closingBracket(r rune, _ []rune) (bool, bool) {
	ok := r == ')'
	return ok, false
}
func ClosingBracket() TesterMaker {
	return &testerMaker{
		title: closingBracketTitle,
		fn:    closingBracket,
	}
}

const openingCurlyBracketTitle = "openingcurlybracket"

func openingCurlyBracket(r rune, _ []rune) (bool, bool) {
	ok := r == '{'
	return ok, false
}
func OpeningCurlyBracket() TesterMaker {
	return &testerMaker{
		title: openingCurlyBracketTitle,
		fn:    openingCurlyBracket,
	}
}

const closingCurlyBracketTitle = "closingcurlybracket"

func closingCurlyBracket(r rune, _ []rune) (bool, bool) {
	ok := r == '}'
	return ok, false
}
func ClosingCurlyBracket() TesterMaker {
	return &testerMaker{
		title: closingCurlyBracketTitle,
		fn:    closingCurlyBracket,
	}
}

const openingSquareBracketTitle = "openingsquarebracket"

func openingSquareBracket(r rune, _ []rune) (bool, bool) {
	ok := r == '{'
	return ok, false
}
func OpeningSquareBracket() TesterMaker {
	return &testerMaker{
		title: openingSquareBracketTitle,
		fn:    openingSquareBracket,
	}
}

const closingSquareBracketTitle = "closingsquarebracket"

func closingSquareBracket(r rune, _ []rune) (bool, bool) {
	ok := r == '}'
	return ok, false
}
func ClosingSquareBracket() TesterMaker {
	return &testerMaker{
		title: closingSquareBracketTitle,
		fn:    closingSquareBracket,
	}
}

const tildaTitle = "tilda"

func tilda(r rune, _ []rune) (bool, bool) {
	ok := r == '~'
	return ok, false
}
func Tilda() TesterMaker {
	return &testerMaker{
		title: tildaTitle,
		fn:    tilda,
	}
}

const atTitle = "at"

func at(r rune, _ []rune) (bool, bool) {
	ok := r == '@'
	return ok, false
}
func At() TesterMaker {
	return &testerMaker{
		title: atTitle,
		fn:    at,
	}
}

const ampersandTitle = "ampersand"

func ampersand(r rune, _ []rune) (bool, bool) {
	ok := r == '&'
	return ok, false
}
func Ampersand() TesterMaker {
	return &testerMaker{
		title: ampersandTitle,
		fn:    ampersand,
	}
}

const hashTitle = "hash"

func hash(r rune, _ []rune) (bool, bool) {
	ok := r == '#'
	return ok, false
}
func Hash() TesterMaker {
	return &testerMaker{
		title: hashTitle,
		fn:    hash,
	}
}

const doubleDotTitle = "doubledot"

func doubleDot(r rune, _ []rune) (bool, bool) {
	ok := r == ':'
	return ok, false
}
func DoubleDot() TesterMaker {
	return &testerMaker{
		title: doubleDotTitle,
		fn:    doubleDot,
	}
}

const roofTitle = "roof"

func roof(r rune, _ []rune) (bool, bool) {
	ok := r == '^'
	return ok, false
}
func Roof() TesterMaker {
	return &testerMaker{
		title: roofTitle,
		fn:    roof,
	}
}

const percentTitle = "percent"

func percent(r rune, _ []rune) (bool, bool) {
	ok := r == '%'
	return ok, false
}
func Percent() TesterMaker {
	return &testerMaker{
		title: percentTitle,
		fn:    percent,
	}
}

const slashTitle = "slash"

func slash(r rune, _ []rune) (bool, bool) {
	ok := r == '/'
	return ok, false
}
func Slash() TesterMaker {
	return &testerMaker{
		title: slashTitle,
		fn:    slash,
	}
}

const backSlashTitle = "backslash"

func backSlash(r rune, _ []rune) (bool, bool) {
	ok := r == '/'
	return ok, false
}
func BackSlash() TesterMaker {
	return &testerMaker{
		title: backSlashTitle,
		fn:    backSlash,
	}
}

const dollarTitle = "dollar"

func dollar(r rune, _ []rune) (bool, bool) {
	ok := r == '$'
	return ok, false
}
func Dollar() TesterMaker {
	return &testerMaker{
		title: dollarTitle,
		fn:    dollar,
	}
}
