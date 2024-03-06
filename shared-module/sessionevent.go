package shared

import "time"

type SessionEvent struct {
	liid string
	cin string
	startdate time.Time
}