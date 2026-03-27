# go-maputil

[![CI](https://github.com/philiprehberger/go-maputil/actions/workflows/ci.yml/badge.svg)](https://github.com/philiprehberger/go-maputil/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/philiprehberger/go-maputil.svg)](https://pkg.go.dev/github.com/philiprehberger/go-maputil) [![License](https://img.shields.io/github/license/philiprehberger/go-maputil)](LICENSE) [![Sponsor](https://img.shields.io/badge/sponsor-GitHub%20Sponsors-ec6cb9)](https://github.com/sponsors/philiprehberger)

Generic map utilities for Go. Filter, transform, merge, group, and more

## Installation

```bash
go get github.com/philiprehberger/go-maputil
```

## Usage

### Filter & Transform

```go
import "github.com/philiprehberger/go-maputil"

scores := map[string]int{"alice": 90, "bob": 45, "charlie": 72}

// Keep entries where value > 50
high := maputil.Filter(scores, func(_ string, v int) bool {
    return v > 50
})
// {"alice": 90, "charlie": 72}

// Double all values
doubled := maputil.Map(scores, func(_ string, v int) int {
    return v * 2
})
// {"alice": 180, "bob": 90, "charlie": 144}
```

### Merge

```go
defaults := map[string]int{"timeout": 30, "retries": 3}
overrides := map[string]int{"timeout": 60}

config := maputil.Merge(defaults, overrides)
// {"timeout": 60, "retries": 3}

// Merge with conflict resolution
summed := maputil.MergeWith(func(_ string, a, b int) int {
    return a + b
}, defaults, overrides)
// {"timeout": 90, "retries": 3}
```

### Pick & Omit

```go
m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

picked := maputil.Pick(m, "a", "c")
// {"a": 1, "c": 3}

omitted := maputil.Omit(m, "b", "d")
// {"a": 1, "c": 3}
```

### Any & All

```go
scores := map[string]int{"alice": 90, "bob": 45, "charlie": 72}

hasHigh := maputil.Any(scores, func(_ string, v int) bool {
    return v >= 90
})
// true

allPassing := maputil.All(scores, func(_ string, v int) bool {
    return v >= 50
})
// false (bob has 45)
```

### GetOrDefault

```go
config := map[string]int{"timeout": 30, "retries": 3}

timeout := maputil.GetOrDefault(config, "timeout", 60)
// 30 (key exists)

maxConns := maputil.GetOrDefault(config, "max_conns", 10)
// 10 (key missing, returns fallback)
```

### Find

```go
scores := map[string]int{"alice": 90, "bob": 45, "charlie": 72}

name, score, ok := maputil.Find(scores, func(_ string, v int) bool {
    return v >= 90
})
// name="alice", score=90, ok=true
```

### Partition

```go
scores := map[string]int{"alice": 90, "bob": 45, "charlie": 72}

passing, failing := maputil.Partition(scores, func(_ string, v int) bool {
    return v >= 50
})
// passing: {"alice": 90, "charlie": 72}
// failing: {"bob": 45}
```

### Diff

```go
v1 := map[string]int{"a": 1, "b": 2, "c": 3}
v2 := map[string]int{"b": 2, "c": 30, "d": 4}

added, removed, changed := maputil.Diff(v1, v2)
// added:   {"d": 4}
// removed: {"a": 1}
// changed: {"c": 30}
```

### GroupBy

```go
type User struct {
    Name       string
    Department string
}

users := []User{
    {"Alice", "Engineering"},
    {"Bob", "Marketing"},
    {"Charlie", "Engineering"},
}

grouped := maputil.GroupBy(users, func(u User) string {
    return u.Department
})
// {"Engineering": [{Alice, Engineering}, {Charlie, Engineering}], "Marketing": [{Bob, Marketing}]}

counts := maputil.CountBy(users, func(u User) string {
    return u.Department
})
// {"Engineering": 2, "Marketing": 1}
```

## API

| Function | Description |
|---|---|
| `Filter` | Return map with entries matching predicate |
| `Map` | Transform values |
| `MapKeys` | Transform keys |
| `Merge` | Merge maps, last wins |
| `MergeWith` | Merge with conflict resolution |
| `Pick` | Select only given keys |
| `Omit` | Exclude given keys |
| `Invert` | Swap keys and values |
| `Keys` | Extract keys (unordered) |
| `SortedKeys` | Extract keys sorted |
| `Values` | Extract values |
| `Contains` | Check if key exists |
| `Size` | Return map size |
| `Any` | True if any entry matches predicate |
| `All` | True if all entries match predicate |
| `GetOrDefault` | Get value or return fallback |
| `Find` | First entry matching predicate |
| `Partition` | Split map by predicate |
| `Diff` | Compare two maps (added, removed, changed) |
| `GroupBy` | Group slice elements by key |
| `CountBy` | Count elements per group |
| `UniqueBy` | Last element per group |

## Development

```bash
go test ./...
go vet ./...
```

## License

MIT
