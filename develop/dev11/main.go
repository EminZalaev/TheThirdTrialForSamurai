package main

import (
	"dev11/handler"
	"log"
	"net/http"
	"time"
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

	handler := Logging(serveMux)

	log.Fatal(http.ListenAndServe("localhost:8081", handler))
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
