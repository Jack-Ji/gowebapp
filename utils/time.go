package utils

import "time"

// 归一化时间范围，从起始天的00:00:00到最后一天的23:59:59
func GetTimeRange(s, e time.Time) (start, end time.Time) {
	loc := s.Location()
	yy, mm, dd := s.Date()
	start = time.Date(yy, mm, dd, 0, 0, 0, 0, loc)
	yy, mm, dd = e.Date()
	end = time.Date(yy, mm, dd, 23, 59, 59, 0, loc)
	return
}
