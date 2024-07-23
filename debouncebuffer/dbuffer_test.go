package debouncebuffer_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"go-ds/debouncebuffer"
	"testing"
	"time"
)

type MockHandler struct {
	called int
}

func (m *MockHandler) msgHandler(msg []string) {
	fmt.Println(msg)
	m.called++
}

func maxDuration(a time.Duration, b time.Duration) time.Duration {
	if a < b {
		return b
	}
	return a
}

func TestMessageBufferCache_Handle(t *testing.T) {
	t.Parallel()
	msg1 := "Hello, my name is Matheus"
	msg2 := "I'm from Brazil"
	msg3 := "Investing time im my own micro saas"

	msgs := []string{msg1, msg2, msg3}

	cases := map[string]struct {
		debounceDuration        time.Duration
		incomingMessageInternal time.Duration
		expectedCallsCount      int
	}{
		"1 flush": {
			debounceDuration:        time.Second / 2,
			incomingMessageInternal: time.Second / 4,
			expectedCallsCount:      1,
		},
		"2 flushes": {
			debounceDuration:        time.Second / 2,
			incomingMessageInternal: time.Second / 2,
			expectedCallsCount:      2,
		},
		"3 flushes": {
			debounceDuration:        time.Second / 4,
			incomingMessageInternal: time.Second / 2,
			expectedCallsCount:      3,
		},
	}

	for testName, testCase := range cases {
		t.Run(testName, func(t *testing.T) {
			mock := MockHandler{
				called: 0,
			}

			db := debouncebuffer.NewDebounceBuffer(testCase.debounceDuration, mock.msgHandler)
			for _, m := range msgs {
				db.Add(m)
				time.Sleep(testCase.incomingMessageInternal)
			}

			time.Sleep(maxDuration(testCase.incomingMessageInternal, testCase.debounceDuration))
			require.Equal(t, testCase.expectedCallsCount, mock.called)
		})

	}
}
