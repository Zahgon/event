package event

import (
	"context"
)

// std default event manager
var std = NewManager("default")

// Std get default event manager
func Std() *Manager {
	_ = "STUB: not implemented"

	// Config set default event manager options
	return nil
}

func Config(fn ...OptionFn) { _ = "STUB: not implemented"; return }

/*************************************************************
 * region Listener
 *************************************************************/

// On register a listener to the event. alias of Listen()
func On(name string, listener Listener, priority ...int) { _ = "STUB: not implemented"; return }

// Once register a listener to the event. trigger once
func Once(name string, listener Listener, priority ...int) { _ = "STUB: not implemented"; return }

// Listen register a listener to the event
func Listen(name string, listener Listener, priority ...int) { _ = "STUB: not implemented"; return }

// Subscribe register a listener to the event
func Subscribe(sbr Subscriber) {
	_ = "STUB: not implemented"

	// AddSubscriber register a listener to the event
	return
}

func AddSubscriber(sbr Subscriber) { _ = "STUB: not implemented"; return }

// HasListeners has listeners for the event name.
func HasListeners(name string) bool { _ = "STUB: not implemented"; return false }

// Reset the default event manager
func Reset() {
	_ = "STUB: not implemented"

	// CloseWait close chan and wait for all async events done.
	return
}

func CloseWait() error { _ = "STUB: not implemented"; return nil }

/*************************************************************
 * region Trigger
 *************************************************************/

// AsyncFire simple async fire event by 'go' keywords
func AsyncFire(e Event) {
	_ = "STUB: not implemented"

	// Async fire event by channel
	return
}

func Async(name string, params M) { _ = "STUB: not implemented"; return }

// FireAsync fire event by channel
func FireAsync(e Event) {
	_ = "STUB: not implemented"

	// FireAsyncCtx async fire event by go channel, and with context TODO need?
	// func FireAsyncCtx(ctx context.Context, e Event)
	return
}

// Trigger alias of Fire
func Trigger(name string, params M) (error, Event) {
	_ = "STUB: not implemented"
	return nil, *

	// Fire listeners by name.
	new(Event)
}

func Fire(name string, params M) (error, Event) {
	_ = "STUB: not implemented"
	return nil,

		// FireCtx listeners by name with context.
		*new(Event)
}

func FireCtx(ctx context.Context, name string, params M) (error, Event) {
	_ = "STUB: not implemented"
	return nil, *new(Event)
}

// FireEvent fire listeners by Event instance.
func FireEvent(e Event) error { _ = "STUB: not implemented"; return nil }

// FireEventCtx fire listeners by Event instance with context.
func FireEventCtx(ctx context.Context, e Event) error { _ = "STUB: not implemented"; return nil }

// TriggerEvent alias of FireEvent
func TriggerEvent(e Event) error { _ = "STUB: not implemented"; return nil }

// MustFire fire event by name. will panic on error
func MustFire(name string, params M) Event { _ = "STUB: not implemented"; return *new(Event) }

// MustTrigger alias of MustFire
func MustTrigger(name string, params M) Event { _ = "STUB: not implemented"; return *new(Event) }

// FireBatch fire multi event at once.
func FireBatch(es ...any) []error { _ = "STUB: not implemented"; return nil }

/*************************************************************
 * region Event
 *************************************************************/

// AddEvent add a pre-defined event.
func AddEvent(e Event) error { _ = "STUB: not implemented"; return nil }

// AddEventFc add a pre-defined event factory func to manager.
func AddEventFc(name string, fc FactoryFunc) error { _ = "STUB: not implemented"; return nil }

// GetEvent get event by name.
func GetEvent(name string) (Event, bool) {
	_ = "STUB: not implemented"
	return *

	// HasEvent has event check.
	new(Event), false
}

func HasEvent(name string) bool { _ = "STUB: not implemented"; return false }
