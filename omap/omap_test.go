package omap

import "testing"

func TestOMapCount(t *testing.T) {
	m := New()
	if n := m.Count(); n != 0 {
		t.Errorf("Count() should return 0 for initial OMap: %d", n)
	}
	m.Add("1st", 10)
	if n := m.Count(); n != 1 {
		t.Errorf("Count() should return 1: %d", n)
	}
	m.Add("2st", 20)
	m.Add("3rd", 30)
	m.Add("4th", 40)
	m.Add("5th", 50)
	if n := m.Count(); n != 5 {
		t.Errorf("Count() should return 5: %d", n)
	}
}

func TestOMapKey(t *testing.T) {
	m := New()
	m.Add("foo", 10)
	m.Add("bar", 20)
	m.Add("abc", 30)
	if k := m.Key(0) ; k != "foo" {
		t.Errorf("Key(0) must foo: %s", k)
	}
	if k := m.Key(1) ; k != "bar" {
		t.Errorf("Key(1) must bar: %s", k)
	}
	if k := m.Key(2) ; k != "abc" {
		t.Errorf("Key(2) must abc: %s", k)
	}
}

func TestOMapKeys(t *testing.T) {
	m := New()
	m.Add("foo", 10)
	m.Add("bar", 20)
	m.Add("abc", 30)
	keys := m.Keys();
	if len(keys) != 3 {
		t.Errorf("Keys() returns not 3: %d", len(keys))
	}
	if keys[0] != "foo" {
		t.Errorf("keys[0] not foo: %s", keys[0])
	}
	if keys[1] != "bar" {
		t.Errorf("keys[1] not bar: %s", keys[1])
	}
	if keys[2] != "abc" {
		t.Errorf("keys[2] not abc: %s", keys[2])
	}
}

func TestOMapGet(t *testing.T) {
	m := New()
	m.Add("foo", 10)
	m.Add("bar", 20)
	m.Add("abc", 30)
	if v := m.Get("foo").(int) ; v != 10 {
		t.Errorf("Get(foo) not 10: %d", v)
	}
	if v := m.Get("bar").(int) ; v != 20 {
		t.Errorf("Get(bar) not 20: %d", v)
	}
	if v := m.Get("abc").(int) ; v != 30 {
		t.Errorf("Get(abc) not 30: %d", v)
	}
	if v := m.Get("xxx") ; v != nil {
		t.Errorf("Get(xxx) not nil: %v", nil)
	}
}
