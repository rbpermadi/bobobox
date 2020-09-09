package hotel

import (
	"fmt"
	"time"

	"github.com/rbpermadi/bobobox/internal/repository"
)

func (ap *AccessProvider) SearchAvailabilities(limit int, offset int, checkinDate string, checkoutDate string, hotelIds string) ([]repository.HotelEntity, error, int) {
	if checkinDate == "" {
		err := fmt.Errorf("Checkin Date Must Not Be NULL")
		return nil, err, 400
	}

	if checkoutDate == "" {
		err := fmt.Errorf("Checkout Date Must Not Be NULL")
		return nil, err, 400
	}

	checkinTime, err := time.Parse(time.RFC3339, checkinDate+"T00:00:00.000Z")
	if err != nil {
		return nil, err, 400
	}

	checkoutTime, err := time.Parse(time.RFC3339, checkoutDate+"T00:00:00.000Z")
	if err != nil {
		return nil, err, 400
	}

	if diff := checkoutTime.Sub(checkinTime); diff < 0 {
		err := fmt.Errorf("Checkout Date Must Be Greater than Checkin Date")
		return nil, err, 400
	}

	hotels := ap.HotelRepo.SearchRoomAvailabilities(limit, offset, checkinDate, checkoutDate, hotelIds)

	return hotels, nil, 200
}
