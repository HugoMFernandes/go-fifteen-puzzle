package stats

import (
	"time"
)

// Stats structure for displaying game stats at the end
type Stats struct {
	elapsedTime string
	numMoves    int
}

// Stats getters
func (s *Stats) ElapsedTime() string {
	return s.elapsedTime
}

func (s *Stats) MoveCount() int {
	return s.numMoves
}

// Stats monitor to actually monitor stuff and create the Stats structure at the end
type StatsMonitor struct {
	startTime time.Time
	endTime   time.Time
	numMoves  int
}

func CreateStatsMonitor() *StatsMonitor {
	return &StatsMonitor{
		startTime: time.Now(),
		endTime:   time.Now(),
		numMoves:  0,
	}
}

func (s *StatsMonitor) StartTimer() {
	s.startTime = time.Now()
}

func (s *StatsMonitor) StopTimer() {
	s.endTime = time.Now()
}

func (s *StatsMonitor) CountMove() {
	s.numMoves++
}

func (s *StatsMonitor) Stats() Stats {
	// This format is pretty ugly sometimes, but it doesn't bother me enough to format
	// it myself
	duration := (time.Since(s.startTime) - time.Since(s.endTime))
	elapsedTime := duration.String()

	return Stats{
		elapsedTime: elapsedTime,
		numMoves:    s.numMoves,
	}
}
