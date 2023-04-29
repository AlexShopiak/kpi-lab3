package test

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }
type LoopSuite struct{}
var _ = Suite(&LoopSuite{})

//TODO
func (s *LoopSuite) TestFailer(c *C) {
	c.Assert("Alex", Equals, "Alexandr")
}
