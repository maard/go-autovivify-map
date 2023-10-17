package avmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AVMapTestSuite struct {
	suite.Suite
}

func (suite *AVMapTestSuite) SetupTest() {
}

func TestAVMapTestSuite(t *testing.T) {
	suite.Run(t, new(AVMapTestSuite))
}

func (suite *AVMapTestSuite) TestBasics() {
	var m = map[string]int{
		"a": 1,
		"b": 2,
	}

	Inc(m, "a")
	Inc(m, "a2")

	SetIfMissing(m, "b", 3)
	SetIfMissing(m, "b2", 3)

	assert.Equal(suite.T(), 2, m["a"])
	assert.Equal(suite.T(), 1, m["a2"])

	assert.Equal(suite.T(), 2, m["b"])
	assert.Equal(suite.T(), 3, m["b2"])
}

func (suite *AVMapTestSuite) TestAdd() {
	var m1 = map[string]int{
		"a": 1,
	}
	var m2 = map[string]string{
		"a": "A",
	}

	Add(m1, "a", 2)
	Add(m2, "a", "B")

	Add(m1, "b", 2)
	Add(m2, "b", "B")

	assert.Equal(suite.T(), 3, m1["a"])
	assert.Equal(suite.T(), "AB", m2["a"])

	assert.Equal(suite.T(), 2, m1["b"])
	assert.Equal(suite.T(), "B", m2["b"])
}

func (suite *AVMapTestSuite) TestAppend() {
	var m = map[string][]string{
		"a": {"a"},
		"b": {},
	}

	Append(m, "a", "b")
	Append(m, "b", "b")
	Append(m, "c", "b")

	assert.Equal(suite.T(), 2, len(m["a"]))
	assert.Equal(suite.T(), "b", m["a"][1])

	assert.Equal(suite.T(), 1, len(m["b"]))
	assert.Equal(suite.T(), "b", m["b"][0])

	assert.Equal(suite.T(), 1, len(m["c"]))
	assert.Equal(suite.T(), "b", m["c"][0])
}
