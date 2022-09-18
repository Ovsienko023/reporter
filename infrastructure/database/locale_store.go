package database

import (
	"sync"
)

type ReportLocalStorage struct {
	reports map[string]*Report
	mutex   *sync.Mutex
}

func NewReportLocalStorage() *ReportLocalStorage {
	return &ReportLocalStorage{
		reports: make(map[string]*Report),
		mutex:   new(sync.Mutex),
	}
}
