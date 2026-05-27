package simpleevent

import (
	"sync"
)

// Wildcard event name
const Wildcard = "*"

// HandlerFunc event handler func define
type HandlerFunc func(e *EventData) error

/*************************************************************
 * Event Manager
 *************************************************************/

// EventManager struct
type EventManager struct {
	pool  sync.Pool
	names map[string]int
	// storage event handlers
	handlers map[string][]HandlerFunc
}

// NewEventManager create EventManager instance
func NewEventManager() *EventManager { _ = "STUB: not implemented"; return nil }

// set pool creator

// On register a event handler
func (em *EventManager) On(name string, handler HandlerFunc) { _ = "STUB: not implemented"; return }

// first add.

// MustFire fire handlers by name. will panic on error
func (em *EventManager) MustFire(name string, args ...any) { _ = "STUB: not implemented"; return }

// Fire handlers by name
func (em *EventManager) Fire(name string, args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// call event handlers

func (em *EventManager) doFire(e *EventData, handlers []HandlerFunc) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// group listen "app.*"
// groupName :=

// Wildcard event handler

func (em *EventManager) callHandlers(e *EventData, handlers []HandlerFunc) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// HasEvent has event check
func (em *EventManager) HasEvent(name string) bool { _ = "STUB: not implemented"; return false }

// GetEventHandlers get handlers and handlers by name
func (em *EventManager) GetEventHandlers(name string) (es []HandlerFunc) {
	_ = "STUB: not implemented"
	return nil
}

// EventHandlers get all event handlers
func (em *EventManager) EventHandlers() map[string][]HandlerFunc {
	_ = "STUB: not implemented"
	return nil

	// EventNames get all event names
}

func (em *EventManager) EventNames() map[string]int {
	_ = "STUB: not implemented"

	// String convert to string.
	return nil
}

func (em *EventManager) String() string { _ = "STUB: not implemented"; return "" }

// ClearHandlers clear handlers by name
func (em *EventManager) ClearHandlers(name string) bool { _ = "STUB: not implemented"; return false }

// Clear all handlers info.
func (em *EventManager) Clear() { _ = "STUB: not implemented"; return }
