package repository

import (
	"context"
	"github.com/nickhalden/mynicceprogram/pkg/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(ctx context.Context, res models.Reservation) (int, error)
	InsertRoomRestriction(ctx context.Context, res models.RoomRestriction) (int, error)
	FindUserById(ctx context.Context, id int) (*models.User, error)
	UpdateUserById(ctx context.Context, user models.User) error
}
