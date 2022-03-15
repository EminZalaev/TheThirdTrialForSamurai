package main

import (
	"dev11/request"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type storeServer struct {
	store *request.StoreServer
}

// func инициализация хранилища
func newStoreServer() *storeServer {
	store := request.NewStore()
	return &storeServer{store: store}
}

// обработчики пост запросов
func (ss *storeServer) handlerCreateEvent(w http.ResponseWriter, r *http.Request) {
	_, date, mes, err := handlerDataPost(w, r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	id := ss.store.CreateEvent(date, mes)
	renderJSON(w, id)
}

func (ss *storeServer) handlerUpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, date, mes, err := handlerDataPost(w, r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	event, err := ss.store.UpdateEvent(id, date, mes)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	renderJSON(w, event)
}

func (ss *storeServer) handlerDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, _, _, err := handlerDataPost(w, r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	errDelete := ss.store.DeleteEvent(id)
	if errDelete != nil {
		http.Error(w, errDelete.Error(), 503)
		return
	}

	renderJSON(w, "Element deleted")

}

// обработчики гет запросов
func (ss *storeServer) handlerEventsForDay(w http.ResponseWriter, r *http.Request) {
	// обрабатываем данные с формы
	date := handlerDataGet(r)
	// обращаемся к хранилищу для обработки
	events, err := ss.store.EventsForDay(date, 0)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	// выводим на экран
	renderJSON(w, events)
}
func (ss *storeServer) handlerEventsForWeek(w http.ResponseWriter, r *http.Request) {
	date := handlerDataGet(r)

	events, err := ss.store.EventsForDay(date, 7)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	renderJSON(w, events)
}
func (ss *storeServer) handlerEventsForMonth(w http.ResponseWriter, r *http.Request) {
	date := handlerDataGet(r)

	events, err := ss.store.EventsForDay(date, 30)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	renderJSON(w, events)
}

// обработчик данных пост запроса
func handlerDataPost(w http.ResponseWriter, r *http.Request) (int, time.Time, string, error) {
	var id int
	var date time.Time
	var mes string

	// сначала проверяем вернула ли форма хоть что-то, а далее пробуем привести в нужный фoрмат
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

// обработчик данных гетзапроса
func handlerDataGet(r *http.Request) time.Time {
	dateF := r.FormValue("date") + "T00:00:00Z"
	date, err := time.Parse(time.RFC3339, dateF)
	if err != nil {
		fmt.Println(err)
	}
	return date
}

// вывод джейсона на страницу
func renderJSON(w http.ResponseWriter, v interface{}) {
	fmt.Printf("v: %v\n", v)

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

func configureRoutes(mux *http.ServeMux, ss *storeServer) {
	mux.HandleFunc("/create_event", ss.handlerCreateEvent)
	mux.HandleFunc("/update_event", ss.handlerUpdateEvent)
	mux.HandleFunc("/delete_event", ss.handlerDeleteEvent)
	mux.HandleFunc("/events_for_day", ss.handlerEventsForDay)
	mux.HandleFunc("/events_for_week", ss.handlerEventsForWeek)
	mux.HandleFunc("/events_for_month", ss.handlerEventsForMonth)
}

func main() {

	mux := http.NewServeMux()
	ss := newStoreServer()
	configureRoutes(mux, ss)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
