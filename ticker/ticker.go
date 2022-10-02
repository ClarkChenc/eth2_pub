package ticker

import (
	"fmt"
	"time"
)

func TestAdjustTicker() {
	var interval time.Duration
	interval = 5 * time.Second

	now := time.Now()
	baseTime := time.Unix(0, 0)
	diff := interval - now.Sub(baseTime)%interval
	fmt.Println("now", now, "diff", diff)

	ticker := time.NewTicker(diff)

	for {
		select {
		case <-ticker.C:
			fmt.Println("get", time.Now())
			ticker.Reset(interval)
		default:
			break
		}
	}
}
