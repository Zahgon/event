package event

import (
	"context"
)

/*************************************************************
 * region Trigger event
 *************************************************************/

// MustTrigger alias of the method MustFire()
func (em *Manager) MustTrigger(name string, params M) Event {
	_ = "STUB: not implemented"
	return *new(Event)
}

// MustFire fire event by name. will panic on error
func (em *Manager) MustFire(name string, params M) Event {
	_ = "STUB: not implemented"
	return *new(Event)
}

// Trigger alias of the method Fire()
func (em *Manager) Trigger(name string, params M) (error, Event) {
	_ = "STUB: not implemented"
	return nil,

		// Fire trigger event by name. if not found listener, will return (nil, nil)
		*new(Event)
}

func (em *Manager) Fire(name string, params M) (err error, e Event) {
	_ = "STUB: not implemented"
	// call listeners handle event
	return nil, *new(Event)
}

// FireCtx fire event by name with context
func (em *Manager) FireCtx(ctx context.Context, name string, params M) (err error, e Event) {
	_ = "STUB: not implemented"
	// call listeners handle event
	return nil, *new(Event)
}

// Async fire event by go channel.
//
// Note: if you want to use this method, you should
// call the method Close() after all events are fired.
func (em *Manager) Async(name string, params M) { _ = "STUB: not implemented"; return }

// FireC async fire event by go channel. alias of the method Async()
//
// Note: if you want to use this method, you should
// call the method Close() after all events are fired.
func (em *Manager) FireC(name string, params M) { _ = "STUB: not implemented"; return }

// fire event by name.
//
// if useCh is true, will async fire by channel. always return (nil, nil)
//
// On useCh=false:
//   - will call listeners handle event.
//   - if not found listener, will return (nil, nil)
func (em *Manager) fireByName(name string, params M, useCh bool) (Event, error) {
	_ = "STUB: not implemented"
	return *new(Event), nil
}

// fireByNameCtx fire event by name with context
func (em *Manager) fireByNameCtx(ctx context.Context, name string, params M, useCh bool) (e Event, err error) {
	_ = "STUB: not implemented"
	return *new(Event), nil
}

// use pre-defined Event

// make new instance

// create new basic event instance

// warp context

// fire by channel

// call listeners handle event

// FireEvent fire event by given Event instance.
func (em *Manager) FireEvent(e Event) error { _ = "STUB: not implemented"; return nil }

// FireEventCtx fire event by given Event instance with context
func (em *Manager) FireEventCtx(ctx context.Context, e Event) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// FireEventCtx fire event by given Event instance with context
func (em *Manager) fireEvent(e Event) (err error) { _ = "STUB: not implemented"; return nil }

// ensure aborted is false.

// get context

// fire group listeners by wildcard. eg "db.user.*"

// handle mode: ModeSimple

// fire wildcard event listeners

// Check context cancellation

// ModeSimple has group listeners by wildcard. eg "db.user.*"
//
// Example:
//   - event "db.user.add" will trigger listeners on the "db.user.*"
func (em *Manager) fireSimpleMode(ctx context.Context, name string, e Event) (err error) {
	_ = "STUB: not implemented"
	// fire direct matched listeners. eg: db.user.add
	return nil
}

// sort by priority before call.

// Check context cancellation

// exists group

// "app.*"

// Check context cancellation

// firePathMode fire group listeners by ModePath.
//
// Example:
//   - event "db.user.add" will trigger listeners on the "db.**"
//   - event "db.user.add" will trigger listeners on the "db.user.*"
func (em *Manager) firePathMode(ctx context.Context, name string, e Event) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Check context cancellation

/*************************************************************
 * region Fire by channel
 *************************************************************/

// FireAsyncCtx async fire event by go channel, and with context TODO need?
// func (em *Manager) FireAsyncCtx(ctx context.Context, e Event)

// FireAsync async fire event by go channel.
//
// Note: if you want to use this method, you should
// call the method Close() after all events are fired.
//
// Example:
//
//	em := NewManager("test")
//	em.FireAsync("db.user.add", M{"id": 1001})
func (em *Manager) FireAsync(e Event) {
	_ = "STUB: not implemented"
	// once make consumers
	return
}

// dispatch event

// async fire event by 'go' keywords
func (em *Manager) makeConsumers() { _ = "STUB: not implemented"; return }

// make event consumers

// keep running until channel closed

// ignore async fire error

// FireBatch fire multi event at once.
//
// Usage:
//
//	FireBatch("name1", "name2", &MyEvent{})
func (em *Manager) FireBatch(es ...any) (ers []error) { _ = "STUB: not implemented"; return nil }

// ignore invalid param.

// AsyncFire simple async fire event by 'go' keywords
func (em *Manager) AsyncFire(e Event) { _ = "STUB: not implemented"; return }

// AwaitFire async fire event by 'go' keywords, but will wait return result
func (em *Manager) AwaitFire(e Event) (err error) { _ = "STUB: not implemented"; return nil }

/*************************************************************
 * region Helper methods
 *************************************************************/

// MustCloseWait close channel and wait all async event done. panic if error
func (em *Manager) MustCloseWait() { _ = "STUB: not implemented"; return }

// CloseWait close channel and wait all async event done.
func (em *Manager) CloseWait() error { _ = "STUB: not implemented"; return nil }

// Wait wait all async event done.
func (em *Manager) Wait() error { _ = "STUB: not implemented"; return nil }
