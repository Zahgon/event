package event

import (
	"regexp"
)

// MatchNodePath check for a string.
//
// From the gookit/goutil/strutil.MatchNodePath()
//
// Use on pattern:
//   - `*` match any to sep
//   - `**` match any to end. only allow at start or end on pattern.
func matchNodePath(pattern, s string, sep string) bool { _ = "STUB: not implemented"; return false }

// at start

// eg: "eve.some.*.*" -> match "eve.some.thing.run" "eve.some.thing.do"

// regex for check good event name.
var goodNameReg = regexp.MustCompile(`^[a-zA-Z][\w-.*]*$`)

// goodName check event name is valid.
func goodName(name string, isReg bool) string { _ = "STUB: not implemented"; return "" }

// goodNameOrErr check event name is valid.
func goodNameOrErr(name string, isReg bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// on add listener

func panicf(format string, args ...any) { _ = "STUB: not implemented"; return }
