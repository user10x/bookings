package repository

import (
	"context"
	"github.com/nickhalden/mynicceprogram/pkg/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(ctx context.Context, res models.Reservation) (int, error)
	InsertRoomRestriction(ctx context.Context, res models.RoomRestriction) (int, error)
	FindUserById(ctx context.Context, id string) (*models.User, error)
}
