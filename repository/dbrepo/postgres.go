package dbrepo

import (
	"context"
	"github.com/nickhalden/mynicceprogram/pkg/models"
	"log"
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

// SearchAvailabilityByDates returns true of availability and false
func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID( start, end time.Time, roomID int) (bool, error)  {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int

	query := `
		select 
			count(id)
		from 
			room_restrictions
		where room_id = $1 and $1 > start_date and $2 < end_date
		
	`
	 row := m.DB.QueryRowContext(ctx, query, roomID, start, end)

	 err := row.Scan(&numRows)

	 if err != nil {
	 	return false, err
	 }

	 if numRows == 0 {
	 	return true, nil
	 }


	return false, nil
}

//SearchAvailabilityAllRoomsByDates searches all rooms available and returns room names
func (m *postgresDBRepo) SearchAvailabilityAllRoomsByDates() string {

	query := `
	select 
		room_name
	from 
		rooms as r, 
		restrictions rs 
	where 
		r.room_id = rs. 
	`
	log.Println(query)

	return ""
}


