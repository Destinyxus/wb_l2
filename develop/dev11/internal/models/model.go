package models

type Event struct {
	Id     uint
	Name   string `json:"event_name"`
	Date   string `json:"event_date"`
	UserId int    `json:"user_id"`
}
