package util

import (
	"fmt"
	"math"
	"time"
)

const Billion = 1000000000

func ConvertNanoToString(nanoseconds float64) string {
	output := ""

	if nanoseconds > float64(time.Hour) {
		hours := math.Floor(nanoseconds / float64(time.Hour))
		output += fmt.Sprintf("%.fh", hours)
		nanoseconds -= hours * float64(time.Hour)
	}

	if nanoseconds > float64(time.Minute) {
		mins := math.Floor(nanoseconds / float64(time.Minute))
		output += fmt.Sprintf("%.fm", mins)
		nanoseconds -= mins * float64(time.Minute)
	}

	val := nanoseconds / float64(time.Second)
	val = val * Billion
	val = math.Floor(val)
	val = val / Billion
	output += fmt.Sprintf("%.9fs", val)

	return output
}

// GetMinMaxTimes gets the new min and max.
// It assumes that when there is only one time, it will return it as min and set max as nil
func GetMinMaxTimes(min, max, newOne *time.Time) (*time.Time, *time.Time) {
	if newOne == nil {
		return min, max
	}

	if min == nil {
		return newOne, nil
	}

	compMin := min.Compare(*newOne)
	if max == nil {

		if compMin > 0 {
			return newOne, min
		} else if compMin < 0 {
			return min, newOne
		} else {
			// newOne is the same as min, so its still just one time provided
			return min, nil
		}
	}

	compMax := max.Compare(*newOne)

	if compMin > 0 {
		return newOne, max
	}

	if compMax < 0 {
		return min, newOne
	}
	return min, max

}
