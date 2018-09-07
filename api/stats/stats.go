package stats

import (
	"time"
)

type ServerStats struct {
	Total     int       `json:"total"`
	ExecTimes []float64 `json:"-"`
	Average   float64   `json:"average"`
}

func New() *ServerStats {
	sStats := ServerStats{}
	return &sStats
}

func (sStats *ServerStats) SuccessfulRequest(startTime time.Time, endTime time.Time) {
	var sum float64
	diff := endTime.Sub(startTime)
	diffMicro := diff.Seconds() * 1000000
	sStats.ExecTimes = append(sStats.ExecTimes, diffMicro)
	for _, num := range sStats.ExecTimes {
		sum += num
	}
	average := sum / float64(len(sStats.ExecTimes))
	sStats.Average = average
	sStats.Total++
}
