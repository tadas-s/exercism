package splitsecondstopwatch

import (
  "errors"
  "fmt"
  "strconv"
  "strings"
  "time"
)

type State int

const (
  StateReady State = iota
  StateRunning
  StateStopped
)

type SplitSecondStopwatch struct {
  state   State
  thisLap time.Duration
  laps    []time.Duration
}

func (sss *SplitSecondStopwatch) Start() error {
  if sss.state == StateRunning {
    return errors.New("cannot start an already running stopwatch")
  }

  sss.state = StateRunning
  return nil
}

func (sss *SplitSecondStopwatch) Stop() error {
  if sss.state == StateReady || sss.state == StateStopped {
    return errors.New("cannot stop a stopwatch that is not running")
  }

  sss.state = StateStopped
  return nil
}

func (sss *SplitSecondStopwatch) Reset() error {
  if sss.state != StateStopped {
    return errors.New("cannot reset a stopwatch that is not stopped")
  }

  sss.state = StateReady
  sss.thisLap = 0
  sss.laps = []time.Duration{}

  return nil
}

func (sss *SplitSecondStopwatch) Lap() error {
  if sss.state != StateRunning {
    return errors.New("cannot lap a stopwatch that is not running")
  }

  sss.laps = append(sss.laps, sss.thisLap)
  sss.thisLap = 0

  return nil
}

func (sss *SplitSecondStopwatch) AdvanceTime(by string) {
  if sss.state != StateRunning {
    return
  }

  durationNumbers := strings.Split(by, ":")

  hours, _ := strconv.Atoi(durationNumbers[0])
  minutes, _ := strconv.Atoi(durationNumbers[1])
  seconds, _ := strconv.Atoi(durationNumbers[2])

  sss.thisLap += time.Duration(hours*60*60+minutes*60+seconds) * time.Second

  return
}

func (sss *SplitSecondStopwatch) State() string {
  switch sss.state {
  case StateReady:
    return "ready"
  case StateRunning:
    return "running"
  case StateStopped:
    return "stopped"
  default:
    return "invalid state"
  }
}

func (sss *SplitSecondStopwatch) CurrentLap() string {
  return formatDuration(sss.thisLap)
}

func (sss *SplitSecondStopwatch) Total() string {
  total := sss.thisLap

  for _, lap := range sss.laps {
    total += lap
  }

  return formatDuration(total)
}

func (sss *SplitSecondStopwatch) PreviousLaps() []string {
  var formattedLaps []string

  for _, lap := range sss.laps {
    formattedLaps = append(formattedLaps, formatDuration(lap))
  }

  return formattedLaps
}

func NewSplitSecondStopwatch() *SplitSecondStopwatch {
  return &SplitSecondStopwatch{
    state:   StateReady,
    thisLap: 0,
    laps:    []time.Duration{},
  }
}

func formatDuration(d time.Duration) string {
  return fmt.Sprintf(
    "%02d:%02d:%02d",
    int(d.Hours())%24,
    int(d.Minutes())%60,
    int(d.Seconds())%60,
  )
}
