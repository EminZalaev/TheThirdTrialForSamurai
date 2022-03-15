package request

import (
	"errors"
	"reflect"
	"time"
)

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
