package common

import (
	"regexp"
	"strconv"
)

// UI struct
type UI struct {
	Analyzer              Analyzer
	IgnoreDirPaths        map[string]struct{}
	IgnoreDirPathPatterns *regexp.Regexp
	IgnoreHidden          bool
	UseColors             bool
	ShowProgress          bool
	ShowApparentSize      bool
	EnableGC              bool
}

// file size constants
const (
	_          = iota
	Ki float64 = 1 << (10 * iota)
	Mi
	Gi
	Ti
	Pi
	Ei
)

// file count constants
const (
	K int = 1e3
	M int = 1e6
	G int = 1e9
)

// FormatNumber returns number as a string with thousands separator
func FormatNumber(n int64) string {
	in := []byte(strconv.FormatInt(n, 10))

	var out []byte
	if i := len(in) % 3; i != 0 {
		if out, in = append(out, in[:i]...), in[i:]; len(in) > 0 {
			out = append(out, ',')
		}
	}
	for len(in) > 0 {
		if out, in = append(out, in[:3]...), in[3:]; len(in) > 0 {
			out = append(out, ',')
		}
	}
	return string(out)
}
