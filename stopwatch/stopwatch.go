package stopwatch

import (
	"errors"
	"fmt"
	"time"
)

type stopWatch struct {
	isRunning   bool
	startTime   time.Time
	elapsedTime time.Duration
	l           laps
}

func newStopWatch(t time.Time, running bool, elapsed int64) *stopWatch {
	return &stopWatch{
		isRunning:   running,
		startTime:   t,
		elapsedTime: time.Duration(elapsed),
		l:           make(laps),
	}
}

func New() *stopWatch {
	return newStopWatch(time.Time{}, false, 0)
}

func NewFrom(t time.Time) *stopWatch {
	e := time.Since(t).Nanoseconds()
	return newStopWatch(t, false, e)
}

func NewStarted() *stopWatch {
	return newStopWatch(time.Now(), true, 0)
}

func NewStartedFrom(t time.Time) *stopWatch {
	e := time.Since(t).Nanoseconds()
	return newStopWatch(t, true, e)
}

func (s *stopWatch) IsRunning() bool {
	return s.isRunning
}

func (s *stopWatch) StartTime() time.Time {
	return s.startTime
}

func (s *stopWatch) Start() (*stopWatch, error) {
	if s.isRunning {
		return s, errors.New("stopwatch is already running")
	}
	s.isRunning = true
	if s.startTime.IsZero() {
		s.startTime = time.Now()
	}
	return s, nil
}

func (s *stopWatch) Stop() (*stopWatch, error) {
	if !s.isRunning {
		return s, errors.New("stopwatch is already stoped")
	}
	s.isRunning = false
	s.elapsedTime = time.Since(s.startTime)
	return s, nil
}

func (s *stopWatch) Reset() *stopWatch {
	s.isRunning = false
	s.startTime = time.Time{}
	s.elapsedTime = time.Duration(0)
	s.l = make(laps)
	return s
}

func (s *stopWatch) Elapsed(u time.Duration) int64 {
	return s.elapsed(u)
}

func (s *stopWatch) ElapsedNanos() int64 {
	return s.elapsed(time.Nanosecond)
}

func (s *stopWatch) ElapsedTime() time.Duration {
	return s.elapsedTime
}

func (s *stopWatch) ElapsedString() string {
	var duration time.Duration

	if s.isRunning {
		if s.startTime.IsZero() {
			return "0"
		}
		duration = time.Since(s.startTime)
	} else {
		duration = s.elapsedTime
	}
	return duration.String()
}

func (s *stopWatch) ElapsedDuration() time.Duration {
	var duration time.Duration
	if s.isRunning {
		if s.startTime.IsZero() {
			return time.Duration(0)
		}
		duration = time.Since(s.startTime)
	} else {
		duration = s.elapsedTime
	}
	return duration
}

func (s *stopWatch) elapsed(u time.Duration) int64 {
	duration := s.ElapsedDuration()
	return toTimeUnit(duration, u)
}

func (s *stopWatch) PrintString() {
	fmt.Println("Elasped Time:", s.ElapsedString())
}

func (s *stopWatch) Print(u time.Duration) {
	time := s.elapsed(u)
	fmt.Println("Elasped Time:", time)
}

func toTimeUnit(duration time.Duration, u time.Duration) int64 {
	nanos := duration.Nanoseconds()
	if u == time.Nanosecond {
		return int64(nanos)
	} else if u == time.Microsecond {
		return int64(nanos / 1000)
	} else if u == time.Millisecond {
		return int64(nanos / 1000000)
	} else if u == time.Second {
		return int64(duration.Seconds())
	} else if u == time.Minute {
		return int64(duration.Minutes())
	} else if u == time.Hour {
		return int64(duration.Hours())
	}
	return int64(0)
}
