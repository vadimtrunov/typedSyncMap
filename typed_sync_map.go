package common

import "sync"

// TypedSyncMap is a thread-safe map with typed keys and values. It is a wrapper around sync.Map that provides type safety.
type TypedSyncMap[Key any, Item any] struct {
	internalMap sync.Map
}

// NewTypedSyncMap creates a new TypedSyncMap. It has to be used instead of direct initialization to avoid type errors.
func NewTypedSyncMap[Key any, Item any]() TypedSyncMap[Key, Item] {
	return TypedSyncMap[Key, Item]{
		internalMap: sync.Map{},
	}
}

// Load returns the value stored in the map for a key, or nil if no value is present.
// The ok result indicates whether value was found in the map.
// It is safe for concurrent use.
func (m *TypedSyncMap[Key, Item]) Load(key Key) (val Item, ok bool) {
	v, ok := m.internalMap.Load(key)

	if !ok {
		return *new(Item), ok
	}

	return v.(Item), ok
}

// Store sets the value for a key.
func (m *TypedSyncMap[Key, Item]) Store(key Key, value Item) {
	m.internalMap.Store(key, value)
}

// Delete deletes the value for a key.
func (m *TypedSyncMap[Key, Item]) Delete(key Key) {
	m.internalMap.Delete(key)
}

// Range calls f sequentially for each key and value present in the map.
func (m *TypedSyncMap[Key, Item]) Range(f func(key Key, value Item) bool) {
	m.internalMap.Range(func(key, value interface{}) bool {
		return f(key.(Key), value.(Item))
	})
}

// CopyFrom copies all key-value pairs from the source map to the destination map.
func (m *TypedSyncMap[Key, Item]) CopyFrom(src *TypedSyncMap[Key, Item]) {
	src.Range(func(key Key, value Item) bool {
		m.Store(key, value)
		return true
	})
}
