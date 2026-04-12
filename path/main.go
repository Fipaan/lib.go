package path

import (
	"path/filepath"
)

func ComparePaths(p1, p2 string) bool {
	p1Full, _ := filepath.EvalSymlinks(p1)
	p2Full, _ := filepath.EvalSymlinks(p2)
	return p1Full == p2Full
}
