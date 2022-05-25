package stopwatch_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"

	"github.com/sendya/pkg/stopwatch"
)

func TestStopWatch(t *testing.T) {
	watch := stopwatch.New()

	assert.Equal(t, watch.IsRunning(), false)
	assert.Equal(t, watch.StartTime(), time.Time{})
	assert.Equal(t, watch.ElapsedString(), "0s")
}

func TestStopWatch2(t *testing.T) {
	watch := stopwatch.New()

	watch.Start()
	time.Sleep(time.Second * 3)
	watch.Stop()

	fmt.Println(watch.ElapsedString())
}
