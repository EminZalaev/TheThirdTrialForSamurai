package service

import (
	"sync"
	"time"
)

type EventCalendar struct {
	Date time.Time
	Mes  string
}

type StoreServer struct {
	m     sync.Mutex
	store map[int]EventCalendar
}
