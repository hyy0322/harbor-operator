// Code generated by github.com/gobuffalo/mapgen. DO NOT EDIT.

package here

import (
	"sort"
	"sync"
)

// infoMap wraps sync.Map and uses the following types:
// key:   string
// value: Info
type infoMap struct {
	data *sync.Map
}

// Delete the key from the map
func (m *infoMap) Delete(key string) {
	m.data.Delete(key)
}

// Load the key from the map.
// Returns Info or bool.
// A false return indicates either the key was not found
// or the value is not of type Info
func (m *infoMap) Load(key string) (Info, bool) {
	i, ok := m.data.Load(key)
	if !ok {
		return Info{}, false
	}
	s, ok := i.(Info)
	return s, ok
}

// LoadOrStore will return an existing key or
// store the value if not already in the map
func (m *infoMap) LoadOrStore(key string, value Info) (Info, bool) {
	i, _ := m.data.LoadOrStore(key, value)
	s, ok := i.(Info)
	return s, ok
}

// LoadOr will return an existing key or
// run the function and store the results
func (m *infoMap) LoadOr(key string, fn func(*infoMap) (Info, bool)) (Info, bool) {
	i, ok := m.Load(key)
	if ok {
		return i, ok
	}
	i, ok = fn(m)
	if ok {
		m.Store(key, i)
		return i, ok
	}
	return i, false
}

// Range over the Info values in the map
func (m *infoMap) Range(f func(key string, value Info) bool) {
	m.data.Range(func(k, v interface{}) bool {
		key, ok := k.(string)
		if !ok {
			return false
		}
		value, ok := v.(Info)
		if !ok {
			return false
		}
		return f(key, value)
	})
}

// Store a Info in the map
func (m *infoMap) Store(key string, value Info) {
	m.data.Store(key, value)
}

// Keys returns a list of keys in the map
func (m *infoMap) Keys() []string {
	var keys []string
	m.Range(func(key string, value Info) bool {
		keys = append(keys, key)
		return true
	})
	sort.Strings(keys)
	return keys
}
