package stats

import (
	"reflect"
	"testing"
	"time"
)

func TestSuccessfulRequest(t *testing.T) {
	tests := []struct {
		expectedTotal   int
		expectedExec    []float64
		expectedAverage float64
	}{
		{
			expectedTotal:   1,
			expectedExec:    []float64{1000000},
			expectedAverage: 1000000,
		}, {
			expectedTotal:   2,
			expectedExec:    []float64{1000000, 1000000},
			expectedAverage: 1000000,
		}, {
			expectedTotal:   3,
			expectedExec:    []float64{1000000, 1000000, 1000000},
			expectedAverage: 1000000,
		},
	}

	sStats := New()
	mockStartTime := time.Now()
	mockEndTime := mockStartTime.Add(1 * time.Second)
	for _, tc := range tests {
		sStats.SuccessfulRequest(mockStartTime, mockEndTime)
		//Test the Total
		if sStats.Total != tc.expectedTotal {
			t.Errorf("SuccessfulRequest returned wrong Total: got %v want %v",
				sStats.Total, tc.expectedTotal)
		}
		//Test ExecTimes
		if !reflect.DeepEqual(tc.expectedExec, sStats.ExecTimes) {
			t.Errorf("SuccessfulRequest returned wrong ExecTimes: got %v want %v",
				sStats.ExecTimes, tc.expectedExec)
		}

		//Test Average
		if sStats.Average != tc.expectedAverage {
			t.Errorf("SuccessfulRequest returned wrong Average: got %v want %v",
				sStats.Average, tc.expectedAverage)
		}
	}
}
