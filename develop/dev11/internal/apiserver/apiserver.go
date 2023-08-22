package apiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"dev11/internal/storage"
	"dev11/pkg"
)

type ApiServer struct {
	router *http.ServeMux
	db     *storage.Store
}

func NewApiServer() *ApiServer {
	return &ApiServer{
		router: http.NewServeMux(),
		db:     storage.NewStore(),
	}
}
func (s *ApiServer) Run() error {
	s.configureRouting()
	if err := s.configureStore(); err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Server is running...")
	err := http.ListenAndServe(":8080", s.router)
	if err != nil {
		return err
	}
	return nil
}

func (s *ApiServer) configureRouting() {
	s.router.HandleFunc("/create_event", s.withLogger(s.createEvent))
	s.router.HandleFunc("/update_event", s.withLogger(s.updateEvent))
	s.router.HandleFunc("/delete_event", s.withLogger(s.deleteEvent))
	s.router.HandleFunc("/events_for_day", s.withLogger(s.getForDay))
	s.router.HandleFunc("/events_for_week", s.withLogger(s.GetForWeek))
	s.router.HandleFunc("/events_for_month", s.withLogger(s.GetForMonth))

}

func (s *ApiServer) withLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		duration := time.Since(start)
		log.Printf("[%s] %s %s %s", r.Method, r.RemoteAddr, r.URL.Path, duration)
	}
}

func (s *ApiServer) configureStore() error {
	if err := s.db.Init(); err != nil {
		return err
	}
	return nil
}

func (s *ApiServer) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	eventName := r.FormValue("event_name")
	eventDate := r.FormValue("event_date")
	userId := r.FormValue("user_id")
	id, _ := strconv.Atoi(userId)

	da, ok := pkg.ValidateDate(eventDate)
	if !ok {
		writeToJson(w, http.StatusBadRequest, "Not a valid format date")
		return
	}

	if err := s.db.CreateEvent(eventName, da, id); err != nil {
		writeToJson(w, http.StatusInternalServerError, "db error")
		return
	}

	if err := writeToJson(w, http.StatusCreated, "created"); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

}

func (s *ApiServer) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "provide right method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		log.Fatal(err)
	}
	newEventName := r.FormValue("new_event_name")
	newEventDate := r.FormValue("new_event_date")
	oldName := r.FormValue("old_name")

	d, ok := pkg.ValidateDate(newEventDate)
	if !ok {
		writeToJson(w, http.StatusBadRequest, "Not a valid format date")
		return
	}

	if err := s.db.UpdateEvent(userID, d, newEventName, oldName); err != nil {
		writeToJson(w, http.StatusInternalServerError, "")
		return
	}

	writeToJson(w, http.StatusCreated, "updated")

}

func (s *ApiServer) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "provide right method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		log.Fatal(err)
	}

	eventName := r.FormValue("event_name")

	if err := s.db.DeleteEvent(userID, eventName); err != nil {
		writeToJson(w, http.StatusInternalServerError, "")
		return
	}

	writeToJson(w, http.StatusOK, "deleted")

}

func (s *ApiServer) getForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "provide right method", http.StatusMethodNotAllowed)
		return
	}
	queryParams := r.URL.Query()
	date := queryParams.Get("event_date")
	userID := queryParams.Get("user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		log.Fatal(err)
		return
	}

	d, ok := pkg.ValidateDate(date)
	if !ok {
		writeToJson(w, http.StatusBadRequest, "Not a valid format date")
		return
	}

	events, err := s.db.GetForDay(d, id)
	if err != nil {
		writeToJson(w, http.StatusInternalServerError, events)
		return
	}
	writeToJson(w, http.StatusOK, events)
}

func (s *ApiServer) GetForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	queryParams := r.URL.Query()
	startDate := queryParams.Get("start_date")
	userID := queryParams.Get("user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		log.Fatal(err)
		return
	}

	d, ok := pkg.ValidateDate(startDate)
	if !ok {
		writeToJson(w, http.StatusBadRequest, "Not a valid format date")
		return
	}

	endDate := pkg.GetEventsForWeek(d)
	events, err := s.db.GetForWeekAndMonth(d, endDate, id)
	if err != nil {
		writeToJson(w, http.StatusInternalServerError, events)
	}

	writeToJson(w, http.StatusOK, events)
}

func (s *ApiServer) GetForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()
	startDate := queryParams.Get("start_date")
	userID := queryParams.Get("user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		log.Fatal(err)
		return
	}
	d, ok := pkg.ValidateDate(startDate)
	if !ok {
		writeToJson(w, http.StatusBadRequest, "Not a valid format date")
		return
	}
	endDate := pkg.GetEventsForMonth(d)

	events, err := s.db.GetForWeekAndMonth(d, endDate, id)
	if err != nil {
		writeToJson(w, http.StatusInternalServerError, events)
	}

	writeToJson(w, http.StatusOK, events)
}

func writeToJson(w http.ResponseWriter, code int, value interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(value)
}
