package internal

import (
	"fmt"
	"time"
)

type TimeRange struct {
	from time.Time
	to   time.Time
}

func NewTimeRange(from time.Time, to time.Time) TimeRange {
	return TimeRange{from: from, to: to}
}

func (n TimeRange) From() time.Time {
	return n.from
}

func (n TimeRange) To() time.Time {
	return n.to
}

func (n TimeRange) IsSameNight(o time.Time) bool {
	return o.Equal(n.from) || o.Equal(n.to) || (o.After(n.from) && o.Before(n.to))
}

func (n TimeRange) String() string {
	from := n.from.Format("2006-01-02")
	to := n.to.Format("2006-01-02")

	return fmt.Sprintf("%s - %s", from, to)
}
