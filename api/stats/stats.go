package stats

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type key int

const (
	startTimeKey key = 1
)

var total int
var execTimes []float64

type clientStats struct {
	Total   int     `json:"total"`
	Average float64 `json:"average"`
}

// func AddToTotal() {
// 	total = total + 1
// }

// func AddTime(time int) {
// 	execTimes = append(execTimes, time)
// }

func StartTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := context.WithValue(r.Context(), startTimeKey, start)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CalcDuration() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime, ok := r.Context().Value(startTimeKey).(time.Time)
		if !ok {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		}
		diff := time.Now().Sub(startTime)
		diffMicro := diff.Seconds() * 1000000
		fmt.Println("Duration is")
		fmt.Println(diff.Seconds() * 1000000)
		execTimes = append(execTimes, diffMicro)

	})
}

func GetStats() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sum float64
		for _, num := range execTimes {
			sum += num
		}
		average := sum / float64(len(execTimes))
		cStats := clientStats{
			Total:   total,
			Average: average,
		}
		clientJSON, err := json.Marshal(cStats)
		if err != nil {
			// handle error here
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(clientJSON)
	})
}
