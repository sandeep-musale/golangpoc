package shared

import "time"

type SessionEvent struct {
	LIID string
	CIN string
	StartDate time.Time
}