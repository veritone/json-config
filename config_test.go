package jsonconfig

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type ConfigTestSuite struct{}

var _ = Suite(&ConfigTestSuite{})

func (s *ConfigTestSuite) SetUpSuite(c *C) {
	//	c.Skip("Skipping")
}

func (s *ConfigTestSuite) TestLoadConfigFromFile(c *C) {
	cfg, err := LoadConfigFromFile("test.json")
	c.Assert(err, IsNil)
	c.Assert(cfg, Not(IsNil))

	st, ok := cfg.GetString("key0")
	c.Assert(ok, Equals, false)

	st, ok = cfg.GetString("svalue")
	c.Assert(ok, Equals, true)
	c.Assert(st, Equals, "value1")

	i, ok := cfg.GetInt("ivalue")
	c.Assert(ok, Equals, true)
	c.Assert(i, Equals, 2)

	b, ok := cfg.GetBool("bvalue")
	c.Assert(ok, Equals, true)
	c.Assert(b, Equals, true)

	sa, ok := cfg.GetStringArray("sarray")
	c.Assert(ok, Equals, true)
	c.Assert(len(sa), Equals, 2)
	c.Assert(sa[0], Equals, "a1")
	c.Assert(sa[1], Equals, "a2")

	ia, ok := cfg.GetIntArray("iarray")
	c.Assert(ok, Equals, true)
	c.Assert(len(ia), Equals, 2)
	c.Assert(ia[0], Equals, 6)
	c.Assert(ia[1], Equals, 7)

	obj, ok := cfg.GetObject("oval")
	c.Assert(ok, Equals, true)
	c.Assert(obj, Not(IsNil))

	st, ok = obj.GetString("key0")
	c.Assert(ok, Equals, false)

	st, ok = obj.GetString("svalue")
	c.Assert(ok, Equals, true)
	c.Assert(st, Equals, "value1")

	i, ok = obj.GetInt("ivalue")
	c.Assert(ok, Equals, true)
	c.Assert(i, Equals, 2)

	b, ok = obj.GetBool("bvalue")
	c.Assert(ok, Equals, true)
	c.Assert(b, Equals, true)

	sa, ok = obj.GetStringArray("sarray")
	c.Assert(ok, Equals, true)
	c.Assert(len(sa), Equals, 2)
	c.Assert(sa[0], Equals, "a1")
	c.Assert(sa[1], Equals, "a2")

	ia, ok = obj.GetIntArray("iarray")
	c.Assert(ok, Equals, true)
	c.Assert(len(ia), Equals, 2)
	c.Assert(ia[0], Equals, 6)
	c.Assert(ia[1], Equals, 7)

	obj2, ok := cfg.GetObject("oval")
	c.Assert(ok, Equals, true)
	c.Assert(obj, Not(IsNil))

	st, ok = obj2.GetString("key0")
	c.Assert(ok, Equals, false)

	st, ok = obj2.GetString("svalue")
	c.Assert(ok, Equals, true)
	c.Assert(st, Equals, "value1")
}

func (s *ConfigTestSuite) TestLoadConfigFromString(c *C) {
	cfg, err := LoadConfigFromString(`{"key1":"value1","key2":2,"key3":true,"key4":["a1","a2"],"key5":[6,7]}`)
	c.Assert(err, IsNil)
	c.Assert(cfg, Not(IsNil))

	st, ok := cfg.GetString("key0")
	c.Assert(ok, Equals, false)

	st, ok = cfg.GetString("key1")
	c.Assert(ok, Equals, true)
	c.Assert(st, Equals, "value1")

	i, ok := cfg.GetInt("key2")
	c.Assert(ok, Equals, true)
	c.Assert(i, Equals, 2)

	b, ok := cfg.GetBool("key3")
	c.Assert(ok, Equals, true)
	c.Assert(b, Equals, true)

	sa, ok := cfg.GetStringArray("key4")
	c.Assert(ok, Equals, true)
	c.Assert(len(sa), Equals, 2)
	c.Assert(sa[0], Equals, "a1")
	c.Assert(sa[1], Equals, "a2")

	ia, ok := cfg.GetIntArray("key5")
	c.Assert(ok, Equals, true)
	c.Assert(len(ia), Equals, 2)
	c.Assert(ia[0], Equals, 6)
	c.Assert(ia[1], Equals, 7)
}
