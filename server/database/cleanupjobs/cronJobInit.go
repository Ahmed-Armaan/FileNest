package cleanupjobs

import (
	"time"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/storage"
	"github.com/robfig/cron"
)

func CronInit(db database.DatabaseStore, s storage.StorageStore) {
	c := cron.New()
	c.AddFunc("0 0 */2 * * *", DeleterCron(db, s))
	c.Start()
}

func DeleterCron(db database.DatabaseStore, s storage.StorageStore) cron.FuncJob {
	var job cron.FuncJob
	job = func() {
		deadline := time.Now().Add(30 * time.Minute)

		for time.Now().Before(deadline) {
			done, err := deleteNodes(db, s)
			if err != nil {
				return
			}
			if done {
				return
			}
			time.Sleep(5 * time.Minute)
		}
	}

	return job
}
