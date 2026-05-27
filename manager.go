package event

import (
	"sync"
)

const (
	defaultChannelSize = 100
	defaultConsumerNum = 3
)

// Manager event manager definition. for manage events and listeners
type Manager struct {
	Options
	sync.Mutex

	wg  sync.WaitGroup
	ch  chan Event
	oc  sync.Once
	err error // latest error

	// name of the manager
	name string
	// pool sync.Pool
	// is a sample for new BasicEvent
	sample *BasicEvent

	// storage user custom Event instance. you can pre-define some Event instances.
	events map[string]Event
	// storage user pre-defined event factory func.
	eventFc map[string]FactoryFunc

	// storage all event name and ListenerQueue map
	listeners map[string]*ListenerQueue
	// storage all event names by listened
	listenedNames map[string]int
}

// NewM create event manager. alias of the NewManager()
func NewM(name string, fns ...OptionFn) *Manager { _ = "STUB: not implemented"; return nil }

// NewManager create event manager
func NewManager(name string, fns ...OptionFn) *Manager { _ = "STUB: not implemented"; return nil }

// ctx:  context.Background(),
// sample event

// events storage

// listeners

// em.EnableLock = true
// for async fire by goroutine

// apply options

// WithOptions create event manager with options
func (em *Manager) WithOptions(fns ...OptionFn) *Manager { _ = "STUB: not implemented"; return nil }

/*************************************************************
 * region Register listeners
 *************************************************************/

// AddListener register an event handler/listener. alias of the method On()
func (em *Manager) AddListener(name string, listener Listener, priority ...int) {
	_ = "STUB: not implemented"
	return
}

// Listen register an event handler/listener. alias of the On()
func (em *Manager) Listen(name string, listener Listener, priority ...int) {
	_ = "STUB: not implemented"
	return
}

// Once register an event handler/listener. trigger once.
func (em *Manager) Once(name string, listener Listener, priority ...int) {
	_ = "STUB: not implemented"
	return
}

// On register a event handler/listener. can setting priority.
//
// Usage:
//
//	em.On("evt0", listener)
//	em.On("evt0", listener, High)
func (em *Manager) On(name string, listener Listener, priority ...int) {
	_ = "STUB: not implemented"
	return
}

// Subscribe add events by subscriber interface. alias of the AddSubscriber()
func (em *Manager) Subscribe(sbr Subscriber) { _ = "STUB: not implemented"; return }

// AddSubscriber add events by subscriber interface.
//
// you can register multi event listeners in a struct func.
// more usage please see README or tests.
func (em *Manager) AddSubscriber(sbr Subscriber) { _ = "STUB: not implemented"; return }

// case ListenerFunc:
// 	em.On(name, lt)

func (em *Manager) addListenerItem(name string, li *ListenerItem) {
	_ = "STUB: not implemented"
	return
}

// exists, append it.

// first add.

/*************************************************************
 * region Event Manage
 *************************************************************/

// AddEvent add a pre-defined event instance to manager.
func (em *Manager) AddEvent(e Event) error { _ = "STUB: not implemented"; return nil }

// AddEventFc add a pre-defined event factory func to manager.
func (em *Manager) AddEventFc(name string, fc FactoryFunc) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (em *Manager) addEventFc(name string, fc FactoryFunc) { _ = "STUB: not implemented"; return }

// GetEvent get a pre-defined event instance by name
func (em *Manager) GetEvent(name string) (e Event, ok bool) {
	_ = "STUB: not implemented"
	return *new(Event), false
}

// HasEvent has pre-defined event check
func (em *Manager) HasEvent(name string) bool { _ = "STUB: not implemented"; return false }

// RemoveEvent delete pre-define Event by name
func (em *Manager) RemoveEvent(name string) { _ = "STUB: not implemented"; return }

// RemoveEvents remove all registered events
func (em *Manager) RemoveEvents() { _ = "STUB: not implemented"; return }

/*************************************************************
 * region Helper Methods
 *************************************************************/

// newBasicEvent create new BasicEvent by clone em.sample
func (em *Manager) newBasicEvent(name string, data M) *BasicEvent {
	_ = "STUB: not implemented"
	return nil
}

// HasListeners check has direct listeners for the event name.
func (em *Manager) HasListeners(name string) bool { _ = "STUB: not implemented"; return false }

// Listeners get all listeners
func (em *Manager) Listeners() map[string]*ListenerQueue {
	_ = "STUB: not implemented"

	// ListenersByName get listeners by given event name
	return nil
}

func (em *Manager) ListenersByName(name string) *ListenerQueue {
	_ = "STUB: not implemented"
	return nil

	// ListenersCount get listeners number for the event name.
}

func (em *Manager) ListenersCount(name string) int { _ = "STUB: not implemented"; return 0 }

// ListenedNames get listened event names
func (em *Manager) ListenedNames() map[string]int { _ = "STUB: not implemented"; return nil }

// RemoveListener remove a given listener, you can limit event name.
//
// Usage:
//
//	RemoveListener("", listener)
//	RemoveListener("name", listener) // limit event name.
func (em *Manager) RemoveListener(name string, listener Listener) {
	_ = "STUB: not implemented"
	return
}

// delete from manager

// name is empty. find all listener and remove matched.

// delete from manager

// RemoveListeners remove listeners by given name
func (em *Manager) RemoveListeners(name string) { _ = "STUB: not implemented"; return }

// delete from manager

// Clear alias of the Reset()
func (em *Manager) Clear() {
	_ = "STUB: not implemented"

	// Close event channel, deny to fire new event.
	return
}

func (em *Manager) Close() error { _ = "STUB: not implemented"; return nil }

// Reset the manager, clear all data.
func (em *Manager) Reset() {
	_ = "STUB: not implemented"
	// clear all listeners
	return
}

// reset all
