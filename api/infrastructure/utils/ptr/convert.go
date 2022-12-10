package ptr

import "time"

func String(s string) *string {
	return &s
}

func Int(s int) *int {
	return &s
}

func Int64(s int64) *int64 {
	return &s
}

func Time(s time.Time) *time.Time {
	return &s
}
