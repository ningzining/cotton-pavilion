package timerutil

import "time"

func RunPeriodTask(f func(), duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	f()
	for range ticker.C {
		f()
	}
}
