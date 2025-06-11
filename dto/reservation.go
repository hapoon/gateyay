package dto

import "time"

type Reservation struct {
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	ReservationDate time.Time `json:"reservation_date"`
	Quantity        int       `json:"quantity"`
}
