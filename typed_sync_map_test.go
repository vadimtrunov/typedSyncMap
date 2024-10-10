package sync

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestTypedSyncMap struct {
	suite.Suite
	typedSync TypedSyncMap[string, string]
}

func (suite *TestTypedSyncMap) SetupTest() {
	suite.typedSync = NewTypedSyncMap[string, string]()
}

func TestTypedSyncMapSuite(t *testing.T) {
	suite.Run(t, new(TestTypedSyncMap))
}

func (suite *TestTypedSyncMap) TestStoreAndLoad() {
	suite.typedSync.Store("key1", "value1")
	value, ok := suite.typedSync.Load("key1")
	suite.True(ok)
	suite.Equal("value1", value)
}

func (suite *TestTypedSyncMap) TestDelete() {
	suite.typedSync.Store("key1", "value1")
	suite.typedSync.Delete("key1")
	_, ok := suite.typedSync.Load("key1")
	suite.False(ok)
}

func (suite *TestTypedSyncMap) TestRange() {
	suite.typedSync.Store("key1", "value1")
	suite.typedSync.Store("key2", "value2")

	keys := make(map[string]bool)
	suite.typedSync.Range(func(key, value string) bool {
		keys[key] = true
		return true
	})

	suite.Len(keys, 2)
	suite.True(keys["key1"])
	suite.True(keys["key2"])
}

func (suite *TestTypedSyncMap) TestLoadNotExisting() {
	v, ok := suite.typedSync.Load("key1")
	suite.False(ok)
	suite.Empty(v)
}

func (suite *TestTypedSyncMap) TestCopyFrom() {
	src := NewTypedSyncMap[string, string]()
	src.Store("key1", "value1")
	src.Store("key2", "value2")

	suite.typedSync.CopyFrom(&src)

	value, ok := suite.typedSync.Load("key1")
	suite.True(ok)
	suite.Equal("value1", value)

	value, ok = suite.typedSync.Load("key2")
	suite.True(ok)
	suite.Equal("value2", value)
}
