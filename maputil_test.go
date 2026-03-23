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

func TestAny(t *testing.T) {
	m := map[string]int{"a": 1, "b": 5, "c": 3}
	if !Any(m, func(_ string, v int) bool { return v > 4 }) {
		t.Error("Any: expected true for v > 4")
	}
	if Any(m, func(_ string, v int) bool { return v > 10 }) {
		t.Error("Any: expected false for v > 10")
	}
}

func TestAny_Empty(t *testing.T) {
	m := map[string]int{}
	if Any(m, func(_ string, v int) bool { return true }) {
		t.Error("Any on empty map: expected false")
	}
}

func TestAll(t *testing.T) {
	m := map[string]int{"a": 2, "b": 4, "c": 6}
	if !All(m, func(_ string, v int) bool { return v%2 == 0 }) {
		t.Error("All: expected true for all even")
	}
	if All(m, func(_ string, v int) bool { return v > 3 }) {
		t.Error("All: expected false for v > 3")
	}
}

func TestAll_Empty(t *testing.T) {
	m := map[string]int{}
	if !All(m, func(_ string, v int) bool { return false }) {
		t.Error("All on empty map: expected true (vacuous truth)")
	}
}

func TestGetOrDefault(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	if got := GetOrDefault(m, "a", 99); got != 1 {
		t.Errorf("GetOrDefault existing key: got %d, want 1", got)
	}
	if got := GetOrDefault(m, "z", 99); got != 99 {
		t.Errorf("GetOrDefault missing key: got %d, want 99", got)
	}
}

func TestGetOrDefault_NilMap(t *testing.T) {
	var m map[string]int
	if got := GetOrDefault(m, "a", 42); got != 42 {
		t.Errorf("GetOrDefault nil map: got %d, want 42", got)
	}
}

func TestGetOrDefault_ZeroValue(t *testing.T) {
	m := map[string]int{"a": 0}
	if got := GetOrDefault(m, "a", 99); got != 0 {
		t.Errorf("GetOrDefault zero value key: got %d, want 0", got)
	}
}

func TestFind(t *testing.T) {
	m := map[string]int{"a": 1, "b": 5, "c": 3}
	k, v, ok := Find(m, func(_ string, v int) bool { return v > 4 })
	if !ok {
		t.Fatal("Find: expected to find a match")
	}
	if k != "b" || v != 5 {
		t.Errorf("Find: got (%s, %d), want (b, 5)", k, v)
	}
}

func TestFind_NoMatch(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	k, v, ok := Find(m, func(_ string, v int) bool { return v > 10 })
	if ok {
		t.Error("Find: expected no match")
	}
	if k != "" || v != 0 {
		t.Errorf("Find no match: got (%s, %d), want zero values", k, v)
	}
}

func TestFind_Empty(t *testing.T) {
	m := map[string]int{}
	_, _, ok := Find(m, func(_ string, v int) bool { return true })
	if ok {
		t.Error("Find on empty map: expected no match")
	}
}

func TestPartition(t *testing.T) {
	m := map[string]int{"a": 1, "b": 5, "c": 3, "d": 10}
	matching, rest := Partition(m, func(_ string, v int) bool { return v > 3 })

	if len(matching) != 2 {
		t.Fatalf("Partition matching: got %d entries, want 2", len(matching))
	}
	if matching["b"] != 5 || matching["d"] != 10 {
		t.Errorf("Partition matching: got %v", matching)
	}
	if len(rest) != 2 {
		t.Fatalf("Partition rest: got %d entries, want 2", len(rest))
	}
	if rest["a"] != 1 || rest["c"] != 3 {
		t.Errorf("Partition rest: got %v", rest)
	}
}

func TestPartition_Empty(t *testing.T) {
	m := map[string]int{}
	matching, rest := Partition(m, func(_ string, v int) bool { return true })
	if len(matching) != 0 || len(rest) != 0 {
		t.Errorf("Partition empty: got matching=%v, rest=%v", matching, rest)
	}
}

func TestPartition_AllMatch(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	matching, rest := Partition(m, func(_ string, _ int) bool { return true })
	if len(matching) != 2 || len(rest) != 0 {
		t.Errorf("Partition all match: got matching=%v, rest=%v", matching, rest)
	}
}

func TestPartition_NoneMatch(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	matching, rest := Partition(m, func(_ string, _ int) bool { return false })
	if len(matching) != 0 || len(rest) != 2 {
		t.Errorf("Partition none match: got matching=%v, rest=%v", matching, rest)
	}
}

func TestDiff(t *testing.T) {
	a := map[string]int{"x": 1, "y": 2, "z": 3}
	b := map[string]int{"y": 2, "z": 30, "w": 4}

	added, removed, changed := Diff(a, b)

	if len(added) != 1 || added["w"] != 4 {
		t.Errorf("Diff added: got %v, want {w: 4}", added)
	}
	if len(removed) != 1 || removed["x"] != 1 {
		t.Errorf("Diff removed: got %v, want {x: 1}", removed)
	}
	if len(changed) != 1 || changed["z"] != 30 {
		t.Errorf("Diff changed: got %v, want {z: 30}", changed)
	}
}

func TestDiff_IdenticalMaps(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2}
	b := map[string]int{"a": 1, "b": 2}
	added, removed, changed := Diff(a, b)
	if len(added) != 0 || len(removed) != 0 || len(changed) != 0 {
		t.Errorf("Diff identical: added=%v, removed=%v, changed=%v", added, removed, changed)
	}
}

func TestDiff_EmptyMaps(t *testing.T) {
	a := map[string]int{}
	b := map[string]int{}
	added, removed, changed := Diff(a, b)
	if len(added) != 0 || len(removed) != 0 || len(changed) != 0 {
		t.Errorf("Diff empty: added=%v, removed=%v, changed=%v", added, removed, changed)
	}
}

func TestDiff_FirstEmpty(t *testing.T) {
	a := map[string]int{}
	b := map[string]int{"a": 1, "b": 2}
	added, removed, changed := Diff(a, b)
	if len(added) != 2 || len(removed) != 0 || len(changed) != 0 {
		t.Errorf("Diff first empty: added=%v, removed=%v, changed=%v", added, removed, changed)
	}
}

func TestDiff_SecondEmpty(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2}
	b := map[string]int{}
	added, removed, changed := Diff(a, b)
	if len(added) != 0 || len(removed) != 2 || len(changed) != 0 {
		t.Errorf("Diff second empty: added=%v, removed=%v, changed=%v", added, removed, changed)
	}
}
