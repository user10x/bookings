package dbrepo

import (
	"context"
	"errors"
	"github.com/nickhalden/mynicceprogram/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// InsertReservation inserts a reservation into the database
func (m *postgresDBRepo) InsertReservation(ctx context.Context, res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

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
		return 0, err
	}
	return newID, nil
}

// InsertRoomRestriction inserts a room reservation
func (m *postgresDBRepo) InsertRoomRestriction(ctx context.Context, r models.RoomRestriction) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

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

// FindUserById Finds user by Id
func (m *postgresDBRepo) FindUserById(ctx context.Context, id int) (*models.User, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT id, first_name, last_name, email from users where id=$1`

	row := m.DB.QueryRowContext(ctx, query, id)
	user := &models.User{}
	err := row.Scan(
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.AccessLevel,
		user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserById updates a user
func (m *postgresDBRepo) UpdateUserById(ctx context.Context, u models.User) error {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `UPDATE users set first_name = $1, last_name = $2, email = $3, access_level = $4, updated_at = $5 where id = $6`

	_, err := m.DB.ExecContext(ctx, query, u.FirstName, u.LastName, u.Email, u.AccessLevel, time.Now(), u.ID)

	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) AllUsers(ctx context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var users []models.User

	query := `SELECT id, first_name, last_name, email, access_level, created_at, updated_at from users`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		u := models.User{}
		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Email,
			&u.AccessLevel,
			&u.CreatedAt,
			&u.UpdateAt,
		)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (m *postgresDBRepo) InsertPasswordHashForUser(ctx context.Context, id int, hash string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `UPDATE users_hash set password_hash = $1 where id = $2`

	hash, err := m.generatePasswordHash(ctx, hash)
	if err != nil {
		return err
	}
	_, err = m.DB.ExecContext(ctx, query, hash, id)

	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) generatePasswordHash(ctx context.Context, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

var ErrInvalidCredentials = errors.New("invalid login credentials")

// Authenticate authenticates a user
func (m *postgresDBRepo) Authenticate(ctx context.Context, email, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var id int
	var hashedPassword string

	query := `SELECT id, password_hash from users where email = $1`

	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return 0, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return 0, "", ErrInvalidCredentials
		}
		return 0, "", err
	}

	return id, hashedPassword, nil
}

func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID(ctx context.Context, start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT count(id) from room_restrictions where $1 < end_date and $2 > start_date and room_id = $3`

	var numRows int

	row := m.DB.QueryRowContext(ctx, query, start, end, roomID)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	return numRows == 0, nil
}

/// generate golang code for generating catchy website domain names for your business
