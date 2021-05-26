package repository

import (
	"github.com/nickhalden/mynicceprogram/pkg/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int,error)
	InsertRoomRestriction(res models.RoomRestriction) (int, error)
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
}
