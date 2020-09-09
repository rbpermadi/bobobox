package repository

import (
	"github.com/jmoiron/sqlx"
)

// entity
type HotelEntity struct {
	ID               int    `json:"id" db:"id"`
	HotelName        string `json:"hotel_name" db:"hotel_name"`
	Address          string `json:"address" db:"address"`
	RoomAvailability int    `json:"room_availability,omitempty" db:"room_availability,omitempty"`
}

type IHotelRepo interface {
	SearchRoomAvailabilities(int, int, string, string, string) []HotelEntity
}

type HotelRepo struct {
	db *sqlx.DB
}

//NewMysqlCategory is a function to create implementation of mysql category repository
func NewHotelRepo(db *sqlx.DB) *HotelRepo {
	return &HotelRepo{db}
}

func (hotelRepo *HotelRepo) SearchRoomAvailabilities(limit int, offset int, checkinDate string, checkoutDate string, hotelIds string) []HotelEntity {
	hotelIdsCondition := ``
	he := []HotelEntity{}

	if hotelIds != "" {
		hotelIdsCondition = " AND h.id IN (" + hotelIds + ") "
	}

	query := `
	SELECT 
		rc.id, rc.hotel_name, rc.address, 
		CASE
			WHEN rc.id = rg.id THEN rc.total_rooms - rg.max_guest
			ELSE rc.total_rooms
		END as room_availability
	FROM
		(
			SELECT
				h.*, COUNT(r.id) as total_rooms
			FROM
				hotel h
			JOIN
				room r ON h.id = r.hotel_id
			WHERE
				r.room_status = 'available' ` + hotelIdsCondition + `
			GROUP BY
				h.id
		) as rc
	LEFT OUTER JOIN
		(
			SELECT id, MAX(guests.total_guest) as max_guest
			FROM
				(
					SELECT 
						h.id, sr.date, COUNT(sr.id) as total_guest
					FROM
						hotel h
					JOIN
						room r ON r.hotel_id = h.id
					JOIN
						stay_room sr ON r.id = sr.room_id
					WHERE
						sr.date BETWEEN ? AND ? ` + hotelIdsCondition + `
					GROUP BY
						sr.date, h.id
				) as guests 
			GROUP BY id
		) as rg ON rc.id = rg.id AND rc.total_rooms > rg.max_guest
	LIMIT ?,?
	`

	err := hotelRepo.db.Select(&he, query, checkinDate, checkoutDate, offset, limit)

	if err != nil {
		return []HotelEntity{}
	}

	return he
}
