package tasks

import "time"

const ChannelGeneralID = "1370092141732302960"

var tokyoLoc *time.Location

func InitTasks() {
	tokyoLoc, _ = time.LoadLocation("Asia/Tokyo")
}
