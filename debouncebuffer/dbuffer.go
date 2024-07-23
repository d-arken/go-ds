package debouncebuffer

import (
	"sync"
	"time"
)

type DebounceBuffer struct {
	data     []string
	duration time.Duration
	mutex    sync.Mutex
	timer    *time.Timer
	dispatch func([]string)
}

func NewDebounceBuffer(duration time.Duration, dispatch func([]string)) *DebounceBuffer {
	return &DebounceBuffer{
		data:     make([]string, 0),
		duration: duration,
		dispatch: dispatch,
	}
}

func (db *DebounceBuffer) Add(item string) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.data = append(db.data, item)

	if db.timer != nil {
		db.timer.Stop()
	}
	db.timer = time.AfterFunc(db.duration, db.flush)
}

func (db *DebounceBuffer) flush() {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if len(db.data) > 0 {
		db.dispatch(db.data)
		db.data = make([]string, 0)
	}
}
