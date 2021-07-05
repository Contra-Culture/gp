package ast

type (
	withReplacementsFPM struct {
		runesfpm *runesFPM
		reps     ReplacementsMapping
	}
	withReplacementsFP struct {
		runesfpm *runesFPM
		reps     ReplacementsMapping
		context  *Node
	}
	ReplacementsMapping struct {
		special rune
		mapping map[rune]rune
	}
)

func WITH_REPLACEMENTS(runesfpm *runesFPM, reps ReplacementsMapping) FragmentParserMaker {
	return &withReplacementsFPM{
		runesfpm,
		reps,
	}
}
func (fpm *withReplacementsFPM) Make(context *Node) FragmentParser {
	runesfpm := fpm.runesfpm
	reps := fpm.reps
	return &withReplacementsFP{
		runesfpm,
		reps,
		context,
	}
}
func (fp *withReplacementsFP) Parse(scanner Scanner) (node *Node, err error) {
	var runes []rune
	runesfp := fp.runesfpm.Make(fp.context)
	node, err = runesfp.Parse(scanner)
	replace := false
	if err != nil {
		return
	}
	for _, r := range node.RawParsed {
		if !replace {
			if fp.reps.special == r {
				replace = true
				continue
			}
			runes = append(runes, r)
			continue
		}
		runes = append(runes, fp.reps.mapping[r])
	}
	node.Parsed = runes
	return
}
