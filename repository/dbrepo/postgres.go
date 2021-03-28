package dbrepo

import (
	"context"
	"github.com/nickhalden/mynicceprogram/pkg/models"
	"time"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	var newID int
	defer cancel()

	stmt := `INSERT INTO reservations (first_name, last_name, email, phone, start_date,
			 end_date, room_id, created_at, updated_at)
			 values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0,err
	}
	return newID,nil
}
// InsertRoomRestriction inserts a room reservation
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) (int , error) {
	ctx, cancel := context.WithTimeout(context.Background(),3*time.Second)

	var newID int
	defer cancel()

	stmt := `INSERT INTO room_restrictions (start_date, end_date, room_id,reservation_id, 
            created_at, updated_at,restriction_id) values ($1, $2, $3, $4, $5, $6, $7) returning id
	`

	err := m.DB.QueryRowContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		time.Now(),
		time.Now(),
		r.RestrictionID,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}
	return newID, nil
}


