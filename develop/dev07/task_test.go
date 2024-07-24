package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	tests := []struct {
		name      string
		channels  []<-chan interface{}
		expectDur time.Duration
	}{

		{
			name: "multiple channels",
			channels: []<-chan interface{}{
				sig(2 * time.Hour),
				sig(5 * time.Minute),
				sig(1 * time.Second),
				sig(1 * time.Hour),
				sig(1 * time.Minute),
			},
			expectDur: 1 * time.Second,
		},
		{
			name: "channels with same duration",
			channels: []<-chan interface{}{
				sig(1 * time.Second),
				sig(1 * time.Second),
				sig(1 * time.Second),
				sig(1 * time.Second),
				sig(1 * time.Second),
				sig(1 * time.Second),
			},
			expectDur: 1,
		},
		{
			name:      "no channels",
			channels:  []<-chan interface{}{},
			expectDur: 0,
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			<-or(tt.channels...)
			elapsed := time.Since(start)

			if elapsed < tt.expectDur || elapsed > tt.expectDur+50*time.Millisecond {
				t.Errorf("expected duration ~%v, got %v", tt.expectDur, elapsed)
			}
		})
	}
}
