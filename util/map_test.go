package util

import (
	"sort"
	"testing"
)

func TestGetMapKeys(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	keys := GetMapKeys(m)
	sort.Strings(keys)

	expected := []string{"a", "b", "c"}
	if len(keys) != len(expected) {
		t.Fatalf("Expected %d keys, got %d", len(expected), len(keys))
	}

	for i, k := range keys {
		if k != expected[i] {
			t.Errorf("Expected key %s, got %s", expected[i], k)
		}
	}
}

func TestGetOrFill(t *testing.T) {
	m := map[string]int{
		"exists": 1,
	}

	// Test existing key
	val := GetOrFill(m, "exists", 100)
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	// Test non-existing key with default
	val = GetOrFill(m, "new", 100)
	if val != 100 {
		t.Errorf("Expected 100, got %d", val)
	}
	if m["new"] != 100 {
		t.Errorf("Expected map to be updated with 100")
	}

	// Test non-existing key without default (zero value)
	val = GetOrFill(m, "zero")
	if val != 0 {
		t.Errorf("Expected 0, got %d", val)
	}
	if m["zero"] != 0 {
		t.Errorf("Expected map to be updated with 0")
	}
}
