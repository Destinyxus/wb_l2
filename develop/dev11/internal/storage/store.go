package storage

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Init() error {
	addr := os.Getenv("POSTGRES")
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) CreateEvent(eventName string, eventDate time.Time, userId int) error {
	query := "INSERT INTO events (event_name, event_date, user_id) VALUES ($1, $2, $3)"
	_, err := s.db.Exec(query, eventName, eventDate, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateEvent(userId int, newEventDate time.Time, newEventName, oldEvent string) error {
	query := "UPDATE events SET event_date = $1, event_name = $2 WHERE user_id = $3 AND event_name = $4"
	_, err := s.db.Exec(query, newEventDate, newEventName, userId, oldEvent)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteEvent(userId int, eventName string) error {
	query := "DELETE FROM events WHERE user_id = $1 AND event_name = $2"
	_, err := s.db.Exec(query, userId, eventName)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetForDay(date time.Time, userId int) ([]string, error) {
	query := "SELECT event_name FROM events WHERE event_date = $1 AND user_id = $2"
	rows, err := s.db.Query(query, date, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventNames []string

	for rows.Next() {
		var eventName string
		if err := rows.Scan(&eventName); err != nil {
			return nil, err
		}
		eventNames = append(eventNames, eventName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return eventNames, nil
}

func (s *Store) GetForWeekAndMonth(startDate, endDate time.Time, userId int) ([]string, error) {
	query := "SELECT event_name FROM events WHERE event_date BETWEEN $1 AND $2 AND user_id = $3"
	rows, err := s.db.Query(query, startDate, endDate, userId)
	if err != nil {
		return nil, err
	}

	var eventNames []string
	for rows.Next() {
		var eventName string
		err := rows.Scan(&eventName)
		if err != nil {
			return nil, err
		}
		eventNames = append(eventNames, eventName)
	}

	return eventNames, nil
}
