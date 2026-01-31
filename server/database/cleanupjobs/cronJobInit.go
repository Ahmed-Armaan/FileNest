package cleanupjobs

import (
	"time"

	"github.com/robfig/cron"
)

func CronInit() {
	c := cron.New()
	c.AddFunc("0 */2 * * *", DeleterCron)
	c.Start()
}

func DeleterCron() {
	deadline := time.Now().Add(30 * time.Minute)

	for time.Now().Before(deadline) {
		done, err := deleteNodes()
		if err != nil {
			return
		}
		if done {
			return
		}
		time.Sleep(5 * time.Minute)
	}
}
