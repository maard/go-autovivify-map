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

func (suite *AVMapTestSuite) TestSetIfMissing() {
	var mi = map[string]int{
		"a": 2,
	}

	SetIfMissing(mi, "a", 3)
	SetIfMissing(mi, "a2", 3)

	assert.Equal(suite.T(), 2, mi["a"])
	assert.Equal(suite.T(), 3, mi["a2"])

	var mb = map[int]bool{
		1: false,
	}

	SetIfMissing(mb, 1, true)
	SetIfMissing(mb, 2, true)

	assert.Equal(suite.T(), false, mb[1])
	assert.Equal(suite.T(), true, mb[2])
}

func (suite *AVMapTestSuite) TestGetOrCreateRef() {
	v2 := 2
	var m1 = map[string]*int{
		"a": &v2,
	}

	assert.Equal(suite.T(), 2, *GetOrCreateRef(m1, "a"))
	assert.Equal(suite.T(), 0, *GetOrCreateRef(m1, "b"))

	_, exists := m1["b"]
	assert.Equal(suite.T(), true, exists)

	_, exists = m1["c"]
	assert.Equal(suite.T(), false, exists)

	var m2 = map[string]*[]int{
		"a": {1},
	}

	array := GetOrCreateRef(m2, "a")
	assert.Equal(suite.T(), 1, len(*array))
	*array = append(*array, 2)
	assert.Equal(suite.T(), 2, len(*array))
	assert.Equal(suite.T(), 2, len(*m2["a"]))

	array = GetOrCreateRef(m2, "b")
	assert.Equal(suite.T(), 0, len(*array))
	*array = append(*array, 2)
	assert.Equal(suite.T(), 1, len(*array))
	assert.Equal(suite.T(), 1, len(*m2["b"]))

	type S struct{ name string }
	var m3 = map[string]*S{
		"a": {name: "a"},
	}

	sp := GetOrCreateRef(m3, "a")
	assert.Equal(suite.T(), "a", sp.name)
	sp.name = "c"
	assert.Equal(suite.T(), "c", m3["a"].name)

	sp = GetOrCreateRef(m3, "b")
	assert.Equal(suite.T(), "", sp.name)
	sp.name = "c"
	assert.Equal(suite.T(), "c", m3["b"].name)
}

func (suite *AVMapTestSuite) TestInc() {
	var m = map[string]int{
		"a": 1,
		"b": 2,
	}

	Inc(m, "a")
	Inc(m, "a2")

	assert.Equal(suite.T(), 2, m["a"])
	assert.Equal(suite.T(), 1, m["a2"])
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
