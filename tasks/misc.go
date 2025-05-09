package tasks

import "time"

var tokyoLoc *time.Location

func InitTasks() {
	tokyoLoc, _ = time.LoadLocation("Asia/Tokyo")
}
