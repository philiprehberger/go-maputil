package maputil

import (
	"slices"
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	m := map[string]int{"a": 1, "b": 5, "c": 3, "d": 10}
	got := Filter(m, func(_ string, v int) bool { return v > 3 })
	want := map[string]int{"b": 5, "d": 10}
	if len(got) != len(want) {
		t.Fatalf("Filter: got %v, want %v", got, want)
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("Filter[%s] = %d, want %d", k, got[k], v)
		}
	}
}

func TestFilter_Empty(t *testing.T) {
	m := map[string]int{}
	got := Filter(m, func(_ string, v int) bool { return v > 0 })
	if len(got) != 0 {
		t.Errorf("Filter on empty map: got %v, want empty", got)
	}
}

func TestMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	got := Map(m, func(_ string, v int) int { return v * 2 })
	want := map[string]int{"a": 2, "b": 4, "c": 6}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("Map[%s] = %d, want %d", k, got[k], v)
		}
	}
}

func TestMapKeys(t *testing.T) {
	m := map[string]int{"hello": 1, "world": 2}
	got := MapKeys(m, strings.ToUpper)
	if got["HELLO"] != 1 || got["WORLD"] != 2 {
		t.Errorf("MapKeys: got %v", got)
	}
	if len(got) != 2 {
		t.Errorf("MapKeys: expected 2 entries, got %d", len(got))
	}
}

func TestMerge(t *testing.T) {
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"y": 20, "z": 30}
	got := Merge(a, b)
	if got["x"] != 1 {
		t.Errorf("Merge[x] = %d, want 1", got["x"])
	}
	if got["y"] != 20 {
		t.Errorf("Merge[y] = %d, want 20 (last wins)", got["y"])
	}
	if got["z"] != 30 {
		t.Errorf("Merge[z] = %d, want 30", got["z"])
	}
}

func TestMergeWith(t *testing.T) {
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"y": 20, "z": 30}
	got := MergeWith(func(_ string, existing, incoming int) int {
		return existing + incoming
	}, a, b)
	if got["x"] != 1 {
		t.Errorf("MergeWith[x] = %d, want 1", got["x"])
	}
	if got["y"] != 22 {
		t.Errorf("MergeWith[y] = %d, want 22", got["y"])
	}
	if got["z"] != 30 {
		t.Errorf("MergeWith[z] = %d, want 30", got["z"])
	}
}

func TestPick(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	got := Pick(m, "a", "c")
	if len(got) != 2 {
		t.Fatalf("Pick: got %d entries, want 2", len(got))
	}
	if got["a"] != 1 || got["c"] != 3 {
		t.Errorf("Pick: got %v", got)
	}
}

func TestPick_MissingKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	got := Pick(m, "a", "z", "missing")
	if len(got) != 1 {
		t.Fatalf("Pick with missing keys: got %d entries, want 1", len(got))
	}
	if got["a"] != 1 {
		t.Errorf("Pick[a] = %d, want 1", got["a"])
	}
}

func TestOmit(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	got := Omit(m, "b", "d")
	if len(got) != 2 {
		t.Fatalf("Omit: got %d entries, want 2", len(got))
	}
	if got["a"] != 1 || got["c"] != 3 {
		t.Errorf("Omit: got %v", got)
	}
}

func TestInvert(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	got := Invert(m)
	if got[1] != "a" || got[2] != "b" || got[3] != "c" {
		t.Errorf("Invert: got %v", got)
	}
}

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	got := Keys(m)
	if len(got) != 3 {
		t.Fatalf("Keys: got %d keys, want 3", len(got))
	}
	slices.Sort(got)
	if got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Errorf("Keys: got %v", got)
	}
}

func TestSortedKeys(t *testing.T) {
	m := map[string]int{"c": 3, "a": 1, "b": 2}
	got := SortedKeys(m)
	want := []string{"a", "b", "c"}
	if !slices.Equal(got, want) {
		t.Errorf("SortedKeys: got %v, want %v", got, want)
	}
}

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	got := Values(m)
	if len(got) != 3 {
		t.Fatalf("Values: got %d values, want 3", len(got))
	}
	slices.Sort(got)
	if got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("Values: got %v", got)
	}
}

func TestContains(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	if !Contains(m, "a") {
		t.Error("Contains(m, 'a') = false, want true")
	}
	if Contains(m, "z") {
		t.Error("Contains(m, 'z') = true, want false")
	}
}
