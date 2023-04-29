package test

import (
	. "gopkg.in/check.v1"
)

type ParserSuite struct{}
var _ = Suite(&ParserSuite{})

//TODO
func (s *ParserSuite) TestAccesser(c *C) {
	c.Assert("Alex", Not(Equals), "Alexandr")
}
