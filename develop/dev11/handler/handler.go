package handler

import (
	"dev11/request"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type StoreServer struct {
	store *request.StoreServer
}

func NewStoreServer() *StoreServer {
	store := request.NewStore()
	return &StoreServer{store: store}
}

func (ss *StoreServer) HandlerCreateEvent(w http.ResponseWriter, r *http.Request) {
	_, date, mes, err := HandlerDataPost(w, r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	id := ss.store.CreateEvent(date, mes)
	RenderJSON(w, id)
}

func (ss *StoreServer) HandlerUpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, date, mes, err := HandlerDataPost(w, r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	event, err := ss.store.UpdateEvent(id, date, mes)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	RenderJSON(w, event)
}

func (ss *StoreServer) HandlerDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, _, _, err := HandlerDataPost(w, r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	errDelete := ss.store.DeleteEvent(id)
	if errDelete != nil {
		http.Error(w, errDelete.Error(), 503)
		return
	}

	RenderJSON(w, "Element deleted")

}

func (ss *StoreServer) HandlerEventsForDay(w http.ResponseWriter, r *http.Request) {

	date := HandlerDataGet(r)

	events, err := ss.store.EventsForDay(date, 0)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	RenderJSON(w, events)
}

func (ss *StoreServer) HandlerEventsForWeek(w http.ResponseWriter, r *http.Request) {
	date := HandlerDataGet(r)

	events, err := ss.store.EventsForDay(date, 7)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	RenderJSON(w, events)
}

func (ss *StoreServer) HandlerEventsForMonth(w http.ResponseWriter, r *http.Request) {
	date := HandlerDataGet(r)

	events, err := ss.store.EventsForDay(date, 30)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	RenderJSON(w, events)
}

func HandlerDataPost(w http.ResponseWriter, r *http.Request) (int, time.Time, string, error) {
	var id int
	var date time.Time
	var mes string

	idString := r.FormValue("id")
	if idString != "" {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			return 0, time.Time{}, "", errors.New("400: invalid int")
		}

		id = idInt
	}

	dateString := r.FormValue("date")
	if dateString != "" {
		dateString += "T00:00:00Z"
		dateTime, err := time.Parse(time.RFC3339, dateString)
		if err != nil {
			return 0, time.Time{}, "", errors.New("400: invalid date")
		}

		date = dateTime
	}

	mes = r.FormValue("mes")

	return id, date, mes, nil
}

func HandlerDataGet(r *http.Request) time.Time {
	dateF := r.FormValue("date") + "T00:00:00Z"
	date, err := time.Parse(time.RFC3339, dateF)
	if err != nil {
		fmt.Println(err)
	}
	return date
}

func RenderJSON(w http.ResponseWriter, v interface{}) {
	fmt.Printf("request: %v\n", v)

	resultJSON := struct {
		Result interface{}
	}{Result: v}

	js, err := json.Marshal(&resultJSON)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}