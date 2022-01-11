package gp

import "github.com/Contra-Culture/report"

type (
	universe struct {
		rules map[string]interface{}
	}
	UniverseCfgr struct {
		universe     universe
		namesToCheck []string
		report       report.Node
	}
)

func New(cfg func(*UniverseCfgr)) (p *Parser, err error) {
	uc := &UniverseCfgr{
		universe: universe{
			rules: map[string]interface{}{},
		},
	}
	cfg(uc)
	ok := uc.check()
	if !ok {
		return
	}
	p = &Parser{
		universe: uc.universe,
	}
	return
}
func (c *UniverseCfgr) check() (ok bool) {
	_, ok = c.universe.rules[TOP_NAME]
	if !ok {
		c.report.Error("top-level parser is not specified")
		return false
	}
outer:
	for _, nameToCheck := range c.namesToCheck {
		for name := range c.universe.rules {
			if nameToCheck == name {
				continue outer
			}
		}
		c.report.Error("wrong parser name \"%s\"", nameToCheck)
		return false
	}
	return true
}
func (c *UniverseCfgr) Top(p Parser) {
	_, exists := c.universe.rules[TOP_NAME]
	if exists {
		c.report.Error("top parser already specified")
		return
	}
	c.universe.rules[TOP_NAME] = p
}
func (c *UniverseCfgr) Define(n string, p Parser) {
	_, exists := c.universe.rules[n]
	if exists {
		c.report.Error("parser \"%s\" already specified", n)
		return
	}
	c.universe.rules[n] = p
}
