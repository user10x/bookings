package repository

import "github.com/nickhalden/mynicceprogram/pkg/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int,error)
	InsertRoomRestriction(res models.RoomRestriction) (int, error)
}
