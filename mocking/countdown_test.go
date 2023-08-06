package countdown

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type SpyCountdownOperations struct {
	CallCount int
	Calls     []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
	s.CallCount++
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	s.CallCount++
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func assertCallCount(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("not enough calls to sleeper, got %d want %d", got, want)
	}
}

func assertCorrectArray(t testing.TB, got, want []string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got calls %v wanted calls %v", got, want)
	}
}

func assertCorrectDuration(t testing.TB, got, want time.Duration) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertCorrectValue(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCount(t *testing.T) {
	t.Run("sleep before every print", func(t *testing.T) {
		spy := &SpyCountdownOperations{}
		Countdown(spy, spy)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		assertCorrectArray(t, spy.Calls, want)
		assertCallCount(t, spy.CallCount, len(want))
	})
	t.Run("correct output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &DefaultSleeper{}
		Countdown(buffer, sleeper)
		got := buffer.String()
		want := `3
2
1
Go!`
		assertCorrectValue(t, got, want)
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()
	assertCorrectDuration(t, spyTime.durationSlept, sleepTime)
}
