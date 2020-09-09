package hotel

import "github.com/rbpermadi/bobobox/internal/repository"

type AccessProvider struct {
	HotelRepo repository.IHotelRepo
}
