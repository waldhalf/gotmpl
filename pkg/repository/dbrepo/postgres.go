package dbrepo

import (
	"context"
	"time"

	"github.com/waldhalf/gotmpl/pkg/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}
// InsertReservation inserts a reservation into the database
func (m *postgresDBRepo)InsertReservation(res models.Reservation) (int, error) {
	// On s'assure que si pb lors de la transaction on timeout au bout de 3 secondes
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newId int

	stmt := `INSERT INTO reservations (first_name, last_name,
			email, phone, start_date, end_date, room_id, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

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
 	).Scan(&newId)

	 if err != nil {
		 return 0, err
	 }

	return newId, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *postgresDBRepo)InsertRoomRestriction(res models.RoomRestriction) (error) {
	// On s'assure que si pb lors de la transaction on timeout au bout de 3 secondes
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO room_restrictions (
		start_date,
		end_date,
		room_id,
		reservation_id,
		created_at,
		updated_at,
		restriction_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		`

	_, err := m.DB.ExecContext(ctx, stmt, 
		res.StartDate,
		res.EndDate,
		res.RoomID,
		res.ReservationID,
		time.Now(),
		time.Now(),
		res.RestrictionID,
	)

	if err != nil {
		return err
	}

	return nil
}

// SearchAvailabilityByDatesByRoomID returns true is availability exists for room id
func (m *postgresDBRepo)SearchAvailabilityByDatesByRoomID(start, end time.Time, roomId int) (bool, error) {
	// On s'assure que si pb lors de la transaction on timeout au bout de 3 secondes
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var numRows int
	query := `SELECT COUNT(ID) 
			FROM 
				room_restrictions
			WHERE 
				room_id = $1
			AND
				$2 < end_date and $3 > start_date;`
	
	row := m.DB.QueryRowContext(ctx, query, roomId, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}

	return false, nil
}

// SearchAvailabilityForAllRooms return a slice of available romms if any for given range date
func (m *postgresDBRepo)SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	// On s'assure que si pb lors de la transaction on timeout au bout de 3 secondes
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room
	query := `SELECT 
				r.id, r.room_name
			FROM
				rooms r
			WHERE r.id NOT IN
				(SELECT room_id 
					FROM room_restrictions rr
				WHERE 
					$1 < rr.end_date AND $2 > rr.start_date)`
	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID, 
			&room.RoomName,
		)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

// GetRoomByID gets a room by Id
func (m *postgresDBRepo)GetRoomByID(id int)(models.Room, error){
	// On s'assure que si pb lors de la transaction on timeout au bout de 3 secondes
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var room models.Room
	query := `SELECT id, room_name, created_at, updated_at 
				FROM rooms
				WHERE id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	if err != nil {
		return room, err
	}

	return room, nil
			
}