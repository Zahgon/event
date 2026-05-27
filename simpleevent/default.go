package simpleevent

// DefaultEM default event manager
var DefaultEM = NewEventManager()

// On register a event and handler
func On(name string, handler HandlerFunc) { _ = "STUB: not implemented"; return }

// Has event check.
func Has(name string) bool { _ = "STUB: not implemented"; return false }

// Fire handlers by name.
func Fire(name string, args ...any) error { _ = "STUB: not implemented"; return nil }

// MustFire fire event by name. will panic on error
func MustFire(name string, args ...any) { _ = "STUB: not implemented"; return }

func funcName(f any) string { _ = "STUB: not implemented"; return "" }
