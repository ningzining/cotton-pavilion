package timerutil

import "time"

func RunPeriodTask(f func(), duration time.Duration) {
	ticker := time.NewTicker(duration)
	go f()
	go func() {
		for range ticker.C {
			f()
		}
	}()
}
