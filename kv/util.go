package kv

import "time"

func TimeStampToEpoch(t time.Time) int {
	return int(t.Unix()/30) - (1598306400 / 30)
}

func EpochToTime(e int) time.Time {
	return time.Unix(int64(e*30)+1598306400, 0)
}