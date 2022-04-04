package main

import (
	"dev11/handler"
	"dev11/middleware"
	"log"
	"net/http"
)

func configureRoutes(serveMux *http.ServeMux, storeServer *handler.StoreServer) {
	serveMux.HandleFunc("/create_event", storeServer.HandlerCreateEvent)
	serveMux.HandleFunc("/update_event", storeServer.HandlerUpdateEvent)
	serveMux.HandleFunc("/delete_event", storeServer.HandlerDeleteEvent)
	serveMux.HandleFunc("/events_for_day", storeServer.HandlerEventsForDay)
	serveMux.HandleFunc("/events_for_week", storeServer.HandlerEventsForWeek)
	serveMux.HandleFunc("/events_for_month", storeServer.HandlerEventsForMonth)
}

func main() {

	serveMux := http.NewServeMux()
	storeServer := handler.NewStoreServer()
	configureRoutes(serveMux, storeServer)

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./tmpl/index.html")
	})

	handler := middleware.Logging(serveMux)

	log.Fatal(http.ListenAndServe("localhost:8081", handler))
}
