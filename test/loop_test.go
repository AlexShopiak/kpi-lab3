package test

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }
type MySuite struct{}
var _ = Suite(&MySuite{})

func (s *MySuite) TestFailer(c *C) {
	c.Assert("Alex", Equals, "Alexandr")
}
