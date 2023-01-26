package main

import (
	"context"
	"time"
)

// Time struct
type Time struct {
	ctx context.Context
}

const (
	layoutTime = "15:04:05"
)

// NewTime creates a new App application struct
func NewTime() *Time {
	return &Time{}
}

// Calculates the time to work
func (a *Time) CalculateWorktime(hr, min, mustHr, mustMin int) string {
	timeNow := time.Now()
	if hr == 0 && min == 0 {
		t := timeNow.Add(time.Hour*time.Duration(mustHr) + time.Minute*time.Duration(mustMin))
		return ("You have to work till: " + t.Format(layoutTime))
	} else {
		then := time.Date(0000, 0, 0, hr, min, timeNow.Second(), 0, time.UTC)
		t := then.Add(time.Hour*time.Duration(mustHr) + time.Minute*time.Duration(mustMin))
		return ("You have to work till: " + t.Format(layoutTime))
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *Time) startup(ctx context.Context) {
	a.ctx = ctx
}
