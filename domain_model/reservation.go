package domain_model

import (
	"time"
)

type Reservation interface {
	At() time.Time
	Email() string
	Name() string
	Quantity() int
	Equals(obj Reservation) bool
}

type reservation struct {
	at       time.Time
	email    string
	name     string
	quantity int
}

func (r reservation) At() time.Time {
	return r.at
}

func (r reservation) Email() string {
	return r.email
}

func (r reservation) Name() string {
	return r.name
}

func (r reservation) Quantity() int {
	return r.quantity
}

func (r reservation) Equals(obj Reservation) bool {
	return obj == r &&
		obj.At() == r.at &&
		obj.Email() == r.email &&
		obj.Name() == r.name &&
		obj.Quantity() == r.quantity
}

func NewReservation(
	at time.Time,
	email string,
	name string,
	quantity int,
) Reservation {
	return &reservation{
		at:       at,
		email:    email,
		name:     name,
		quantity: quantity,
	}
}
