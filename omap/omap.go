package omap

import (
	"errors"
)

// OMap provides ordered map.
type OMap struct {
	keys    []string
	baseMap map[string]interface{}
}

// Entry provides an entry in OMap.  Returned by Entries().
type Entry struct {
	Key string
	Value interface{}
}

// ErrorKeyExisting raised by Add(), when key is existing.
var ErrorKeyExisting = errors.New("key is already existing")

// New returns a OMap.
func New() *OMap {
	return &OMap{
		keys:    make([]string, 0),
		baseMap: make(map[string]interface{}, 0),
	}
}

// Count of items in OMap.
func (m *OMap) Count() int {
	return len(m.keys)
}

// Keys in OMap.
func (m *OMap) Keys() []string {
	return m.keys
}

// Values in OMap
func (m *OMap) Values() []interface{} {
	values := make([]interface{}, len(m.keys), 0)
	for _, k := range m.keys {
		values = append(values, m.baseMap[k])
	}
	return values
}

// Entries enumerate all pairs of key and value in OMap.
func (m *OMap) Entries() []Entry {
	entries := make([]Entry, len(m.keys), 0)
	for _, k := range m.keys {
		entry := Entry{Key: k, Value: m.baseMap[k]}
		entries = append(entries, entry)
	}
	return entries
}

// HasKey check a key existing or not.
func (m *OMap) HasKey(k string) bool {
	_, ok := m.baseMap[k]
	return ok
}

// Key for n'th in order.
func (m *OMap) Key(n int) string {
	return m.keys[n]
}

// Get value for the key.
func (m *OMap) Get(s string) interface{} {
	if n, ok := m.baseMap[s]; ok {
		return n
	}
	return nil
}

// Add a value with key.
func (m *OMap) Add(k string, v interface{}) error {
	if _, ok := m.baseMap[k]; ok {
		return ErrorKeyExisting
	}
	m.add(k, v)
	return nil
}

// Put a value for key, if key is unknown, add a new.
func (m *OMap) Put(k string, v interface{}) bool {
	if !m.HasKey(k) {
		m.add(k, v)
		return false
	}
	// move k to last of m.keys
	n := indexOf(k, m.keys)
	m.keys = append(append(m.keys[:n], m.keys[n+1:]...), k)
	// set v for k in m.baseMap
	m.baseMap[k] = v
	return true
}

// Remove key and value.
func (m *OMap) Remove(k string) bool {
	if !m.HasKey(k) {
		return false
	}
	// remove k from m.keys
	n := indexOf(k, m.keys)
	m.keys = append(m.keys[:n], m.keys[n+1:]...)
	// remove v from m.baseMap.
	delete(m.baseMap, k)
	return true
}

func (m *OMap) add(k string, v interface{}) {
	m.keys = append(m.keys, k)
	m.baseMap[k] = v
}

func indexOf(s string, array []string) int {
	for i, v := range array {
		if v == s {
			return i
		}
	}
	return -1
}
