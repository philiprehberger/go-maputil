package maputil

import (
	"testing"
)

type user struct {
	Name       string
	Department string
}

func TestGroupBy(t *testing.T) {
	users := []user{
		{"Alice", "Engineering"},
		{"Bob", "Marketing"},
		{"Charlie", "Engineering"},
		{"Diana", "Marketing"},
		{"Eve", "Engineering"},
	}
	got := GroupBy(users, func(u user) string { return u.Department })
	if len(got) != 2 {
		t.Fatalf("GroupBy: got %d groups, want 2", len(got))
	}
	if len(got["Engineering"]) != 3 {
		t.Errorf("GroupBy[Engineering]: got %d, want 3", len(got["Engineering"]))
	}
	if len(got["Marketing"]) != 2 {
		t.Errorf("GroupBy[Marketing]: got %d, want 2", len(got["Marketing"]))
	}
}

func TestGroupBy_Empty(t *testing.T) {
	var empty []user
	got := GroupBy(empty, func(u user) string { return u.Department })
	if len(got) != 0 {
		t.Errorf("GroupBy on empty slice: got %v, want empty", got)
	}
}

func TestCountBy(t *testing.T) {
	users := []user{
		{"Alice", "Engineering"},
		{"Bob", "Marketing"},
		{"Charlie", "Engineering"},
		{"Diana", "Marketing"},
		{"Eve", "Engineering"},
	}
	got := CountBy(users, func(u user) string { return u.Department })
	if got["Engineering"] != 3 {
		t.Errorf("CountBy[Engineering] = %d, want 3", got["Engineering"])
	}
	if got["Marketing"] != 2 {
		t.Errorf("CountBy[Marketing] = %d, want 2", got["Marketing"])
	}
}

func TestUniqueBy(t *testing.T) {
	users := []user{
		{"Alice", "Engineering"},
		{"Bob", "Marketing"},
		{"Charlie", "Engineering"},
	}
	got := UniqueBy(users, func(u user) string { return u.Department })
	if len(got) != 2 {
		t.Fatalf("UniqueBy: got %d entries, want 2", len(got))
	}
	// Last element per group wins
	if got["Engineering"].Name != "Charlie" {
		t.Errorf("UniqueBy[Engineering] = %s, want Charlie", got["Engineering"].Name)
	}
	if got["Marketing"].Name != "Bob" {
		t.Errorf("UniqueBy[Marketing] = %s, want Bob", got["Marketing"].Name)
	}
}
