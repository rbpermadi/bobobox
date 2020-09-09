package hotel

import (
	"fmt"
	"time"

	"github.com/rbpermadi/bobobox/internal/repository"
)

func (ap *AccessProvider) SearchAvailabilities(limit int, offset int, checkinDate string, checkoutDate string, hotelIds string) ([]repository.HotelEntity, int, error) {
	if checkinDate == "" {
		err := fmt.Errorf("Checkin Date Must Not Be NULL")
		return nil, 0, err
	}

	if checkoutDate == "" {
		err := fmt.Errorf("Checkout Date Must Not Be NULL")
		return nil, 0, err
	}

	checkinTime, err := time.Parse(time.RFC3339, checkinDate+"T00:00:00.000Z")
	if err != nil {
		err := fmt.Errorf("Format Checkin Date no yyyy-mm-dd")
		return nil, 0, err
	}

	checkoutTime, err := time.Parse(time.RFC3339, checkoutDate+"T00:00:00.000Z")
	if err != nil {
		err := fmt.Errorf("Format Checkout Date no yyyy-mm-dd")
		return nil, 0, err
	}

	if diff := checkoutTime.Sub(checkinTime); diff < 0 {
		err := fmt.Errorf("Checkout Date Must Be Greater than Checkin Date")
		return nil, 0, err
	}

	hotelIdsCondition := ``
	if hotelIds != "" {
		hotelIdsCondition = " AND h.id IN (" + hotelIds + ") "
	}

	count := ap.HotelRepo.CountHotelRoomAvailabilities(checkinDate, checkoutDate, hotelIdsCondition)
	if count == 0 {
		return []repository.HotelEntity{}, 0, nil
	}

	hotels := ap.HotelRepo.SearchHotelRoomAvailabilities(limit, offset, checkinDate, checkoutDate, hotelIdsCondition)

	return hotels, count, nil
}
