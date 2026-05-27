package simpleevent

/*************************************************************
 * Event Data
 *************************************************************/

// EventData struct
type EventData struct {
	aborted bool
	// event name
	name string
	// user data.
	Data []any
}

// Name get
func (e *EventData) Name() string {
	_ = "STUB: not implemented"

	// Abort abort event exec
	return ""
}

func (e *EventData) Abort() {
	_ = "STUB: not implemented"

	// IsAborted check.
	return
}

func (e *EventData) IsAborted() bool { _ = "STUB: not implemented"; return false }

func (e *EventData) init(name string, data []any) {
	e.name = name
	e.Data = data
}

func (e *EventData) reset() { _ = "STUB: not implemented"; return }
