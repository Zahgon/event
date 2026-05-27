// Package event is lightweight event manager and dispatcher implements by Go.
package event

import (
	"context"
)

// wildcard event name
const (
	Wildcard = "*"
	AnyNode  = "*"
	AllNode  = "**"
)

const (
	// ModeSimple old mode, simple match group listener.
	//
	//  - "*" only allow one and must at end
	//
	// Support: "user.*" -> match "user.created" "user.updated"
	ModeSimple uint8 = iota

	// ModePath path mode.
	//
	//  - "*" matches any sequence of non . characters (like at path.Match())
	//  - "**" match all characters to end, only allow at start or end on pattern.
	//
	// Support like this:
	// 	"eve.some.*.*"       -> match "eve.some.thing.run" "eve.some.thing.do"
	// 	"eve.some.*.run"     -> match "eve.some.thing.run", but not match "eve.some.thing.do"
	// 	"eve.some.**"        -> match any start with "eve.some.". eg: "eve.some.thing.run" "eve.some.thing.do"
	// 	"**.thing.run"       -> match any ends with ".thing.run". eg: "eve.some.thing.run"
	ModePath
)

// M is short name for map[string]...
type M = map[string]any

// ManagerFace event manager interface
type ManagerFace interface {
	// AddEvent events: add event
	AddEvent(Event) error
	// On listeners: add listeners
	On(name string, listener Listener, priority ...int)
	// Fire event
	Fire(name string, params M) (error, Event)
}

// Options event manager config options
type Options struct {
	// EnableLock enable lock on fire event. default is False.
	EnableLock bool
	// ChannelSize for fire events by goroutine. default: 100
	ChannelSize int
	// ConsumerNum for fire events by goroutine. default: 3
	ConsumerNum int
	// MatchMode event name match mode. default is ModeSimple
	MatchMode uint8
}

// OptionFn event manager config option func
type OptionFn func(o *Options)

// UsePathMode set event name match mode to ModePath
func UsePathMode(o *Options) { _ = "STUB: not implemented"; return }

// WithChannelSize set channel size for async fire event.
func WithChannelSize(size int) OptionFn { _ = "STUB: not implemented"; return *new(OptionFn) }

// WithConsumerNum set consumer num for async fire
func WithConsumerNum(num int) OptionFn { _ = "STUB: not implemented"; return *new(OptionFn) }

// EnableLock enable lock on fire event.
func EnableLock(enable bool) OptionFn { _ = "STUB: not implemented"; return *new(OptionFn) }

// Event interface
type Event interface {
	Name() string
	Get(key string) any
	Set(key string, val any)
	Add(key string, val any)
	Data() map[string]any
	SetData(M) Event
	Abort(bool)
	IsAborted() bool
}

// Cloneable interface. event can be cloned.
//
// Check and convert:
//
//	if ec, ok := e.(Cloneable); ok {}
type Cloneable interface {
	Event
	Clone() Event
}

// ContextAble context-able event interface
//
// Check and convert in listener:
//
//	if ec, ok := e.(ContextAble); ok {}
type ContextAble interface {
	Event
	Context() context.Context
	WithContext(ctx context.Context)
}

// FactoryFunc for create event instance.
type FactoryFunc func() Event

// BasicEvent a built-in implements Event interface
type BasicEvent struct {
	// event name
	name string
	// user data.
	data map[string]any
	// target
	target any
	// mark is aborted
	aborted bool
}

// New create an event instance
func New(name string, data M) *BasicEvent { _ = "STUB: not implemented"; return nil }

// NewEvent create an event instance
func NewEvent(name string, data M) *BasicEvent { _ = "STUB: not implemented"; return nil }

// NewBasic new a basic event instance
func NewBasic(name string, data M) *BasicEvent { _ = "STUB: not implemented"; return nil }

// Abort event loop exec
func (e *BasicEvent) Abort(abort bool) {
	_ = "STUB: not implemented"

	// Fill event data
	return
}

func (e *BasicEvent) Fill(target any, data M) *BasicEvent { _ = "STUB: not implemented"; return nil }

// AttachTo add current event to the event manager.
func (e *BasicEvent) AttachTo(em ManagerFace) error { _ = "STUB: not implemented"; return nil }

// Get data by index
func (e *BasicEvent) Get(key string) any { _ = "STUB: not implemented"; return *new(any) }

// Add value by key
func (e *BasicEvent) Add(key string, val any) { _ = "STUB: not implemented"; return }

// Set value by key
func (e *BasicEvent) Set(key string, val any) { _ = "STUB: not implemented"; return }

// Name get event name
func (e *BasicEvent) Name() string {
	_ = "STUB: not implemented"

	// Data get all data
	return ""
}

func (e *BasicEvent) Data() map[string]any {
	_ = "STUB: not implemented"

	// IsAborted check.
	return nil
}

func (e *BasicEvent) IsAborted() bool {
	_ = "STUB: not implemented"

	// Target get target
	return false
}

func (e *BasicEvent) Target() any {
	_ = "STUB: not implemented"

	// SetName set event name
	return *new(any)
}

func (e *BasicEvent) SetName(name string) *BasicEvent { _ = "STUB: not implemented"; return nil }

// SetData set data to the event
func (e *BasicEvent) SetData(data M) Event { _ = "STUB: not implemented"; return *new(Event) }

// SetTarget set event target
func (e *BasicEvent) SetTarget(target any) *BasicEvent { _ = "STUB: not implemented"; return nil }

// ContextTrait event context trait
type ContextTrait struct {
	// context
	ctx context.Context
}

// Context get context
func (t *ContextTrait) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// WithContext set context
func (t *ContextTrait) WithContext(ctx context.Context) {
	_ = "STUB: not implemented"

	// ContextEvent event with context
	return
}

type contextEvent struct {
	Event
	ContextTrait
}

func newContextEvent(ctx context.Context, e Event) ContextAble {
	_ = "STUB: not implemented"
	return *new(ContextAble)
}
