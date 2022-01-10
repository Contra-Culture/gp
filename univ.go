package gp

import "github.com/Contra-Culture/report"

type (
	univParser struct {
		parsers map[string]Parser
	}
	UnivCfgr struct {
		univ         *univParser
		namesToCheck []string
		report       report.Node
	}
)
func (u *univParser) Parse(rs *GPRuneScanner) (*Node, error) {
	return u.parsers[TOP_NAME].Parse(rs)
}

func New(cfg func(*UnivCfgr)) (p Parser, err error) {
	uc := &UnivCfgr{
		univ: &univParser{
			parsers: map[string]Parser{},
		},
	}
	cfg(uc)
	ok := uc.check()
	if !ok {
		return
	}
	p = uc.univ
	return
}
func (c *UnivCfgr) check() (ok bool) {
	_, ok = c.univ.parsers[TOP_NAME]
	if !ok {
		c.report.Error("top-level parser is not specified")
		return false
	}
outer:
	for _, nameToCheck := range c.namesToCheck {
		for name := range c.univ.parsers {
			if nameToCheck == name {
				continue outer
			}
		}
		c.report.Error("wrong parser name \"%s\"", nameToCheck)
		return false
	}
	return true
}
func (c *UnivCfgr) Top(p Parser) {
	_, exists := c.univ.parsers[TOP_NAME]
	if exists {
		c.report.Error("top parser already specified")
		return
	}
	c.univ.parsers[TOP_NAME] = p
}
func (c *UnivCfgr) Define(n string, p Parser) {
	_, exists := c.univ.parsers[n]
	if exists {
		c.report.Error("parser \"%s\" already specified", n)
		return
	}
	c.univ.parsers[n] = p
}

