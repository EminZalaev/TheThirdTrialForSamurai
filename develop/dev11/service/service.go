package service

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

func (ss *StoreServer) UpdateEvent(id int, date time.Time, mes string) (EventCalendar, error) {
	ss.m.Lock()
	defer ss.m.Unlock()
	// вернем ошибку если элемента нет

	if reflect.DeepEqual(ss.store[id], EventCalendar{}) {
		return EventCalendar{}, errors.New("503: invalid element")
	}

	event := EventCalendar{date, mes}

	ss.store[id] = event

	return ss.store[id], nil

}

func (ss *StoreServer) DeleteEvent(id int) error {
	ss.m.Lock()
	defer ss.m.Unlock()

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

	if len(result) == 0 {
		return []EventCalendar{}, errors.New("503: Invalid event")
	}

	return result, nil
}
