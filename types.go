package event

import (
	"reflect"
)

// There are some default priority constants
const (
	Min         = -300
	Low         = -200
	BelowNormal = -100
	Normal      = 0
	AboveNormal = 100
	High        = 200
	Max         = 300
)

/*************************************************************
 * region Listener
 *************************************************************/

// Listener interface
type Listener interface {
	Handle(e Event) error
}

// ListenerFunc func definition.
type ListenerFunc func(e Event) error

// Handle event. implements the Listener interface
func (fn ListenerFunc) Handle(e Event) error {
	_ = "STUB: not implemented"

	// Subscriber event subscriber interface.
	//
	// you can register multi event listeners in a struct func.
	return nil
}

type Subscriber interface {
	// SubscribedEvents register event listeners
	//
	//  - key: is event name. eg "user.created" "user.*" "user.**"
	//  - value: can be Listener or ListenerItem interface
	SubscribedEvents() map[string]any
}

// ListenerItem storage a event listener and it's priority value.
type ListenerItem struct {
	Priority int
	Listener Listener
}

/*************************************************************
 * Listener Queue
 *************************************************************/

// ListenerQueue storage sorted Listener instance.
type ListenerQueue struct {
	items []*ListenerItem
}

// Len get items length
func (lq *ListenerQueue) Len() int { _ = "STUB: not implemented"; return 0 }

// IsEmpty get items length == 0
func (lq *ListenerQueue) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Push get items length
func (lq *ListenerQueue) Push(li *ListenerItem) *ListenerQueue {
	_ = "STUB: not implemented"
	return nil
}

// Sort the queue items by ListenerItem's priority.
//
// Priority:
//
//	High > Low
func (lq *ListenerQueue) Sort() *ListenerQueue {
	_ = "STUB: not implemented"
	//	if lq.IsEmpty() {
	//		return lq
	//	}
	return nil
}

// check items is sorted

// Items get all ListenerItem
func (lq *ListenerQueue) Items() []*ListenerItem {
	_ = "STUB: not implemented"

	// Remove a listener from the queue
	return nil
}

func (lq *ListenerQueue) Remove(listener Listener) { _ = "STUB: not implemented"; return }

// unsafe.Pointer(listener)

// Clear all listeners
func (lq *ListenerQueue) Clear() { _ = "STUB: not implemented"; return }

// getListenCompareKey get listener compare key
func getListenCompareKey(src Listener) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

/*************************************************************
 * Sorted PriorityItems
 *************************************************************/

// ByPriorityItems type. implements the sort.Interface
type ByPriorityItems []*ListenerItem

// Len get items length
func (ls ByPriorityItems) Len() int {
	_ = "STUB: not implemented"

	// Less implements the sort.Interface.Less.
	return 0
}

func (ls ByPriorityItems) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements the sort.Interface.Swap.
func (ls ByPriorityItems) Swap(i, j int) { _ = "STUB: not implemented"; return }
