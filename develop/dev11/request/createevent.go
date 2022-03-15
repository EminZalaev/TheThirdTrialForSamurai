package request

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

func NewStore() *StoreServer {
	ss := &StoreServer{}
	ss.m = sync.Mutex{}
	ss.store = make(map[int]EventCalendar)
	return ss
}

type EventCalendar struct {
	Date time.Time
	Mes  string
}

type StoreServer struct {
	m     sync.Mutex
	store map[int]EventCalendar
}

func (ss *StoreServer) CreateEvent(date time.Time, mes string) int {

	event := EventCalendar{date, mes}

	ss.m.Lock()
	defer ss.m.Unlock()

	id := len(ss.store)

	for {
		if reflect.DeepEqual(ss.store[id], EventCalendar{}) {
			ss.store[id] = event
			return id
		}
		id++
	}
}

func (ss *StoreServer) DeleteEvent(id int) error {
	ss.m.Lock()
	defer ss.m.Unlock()
	// вернем ошибку если элемента нет
	if reflect.DeepEqual(ss.store[id], EventCalendar{}) {
		return errors.New("503: No event for delete")
	}

	delete(ss.store, id)
	return nil
}

func (ss *StoreServer) EventsForDay(date time.Time, days int) ([]EventCalendar, error) {
	var result []EventCalendar
	for _, event := range ss.store {
		if event.Date.Sub(date) >= time.Duration(days*time.Now().Day()) {
			result = append(result, event)
		}
	}
	// вернем ошибку если элементы не были надены
	if len(result) == 0 {
		return []EventCalendar{}, errors.New("503: Invalid event")
	}

	return result, nil
}
