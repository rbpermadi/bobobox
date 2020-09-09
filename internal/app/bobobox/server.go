package bobobox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/rbpermadi/bobobox/internal/config"
	"github.com/rbpermadi/bobobox/internal/repository"
	"github.com/rbpermadi/bobobox/internal/service/hotel"
	"github.com/rs/cors"

	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//ServerUp to run server
func ServerUp() {
	cfg := config.GetInstance()

	router := httprouter.New()
	router.GET("/healthz", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		fmt.Println("test")
		resp := Response{
			Data:    cfg.DB.Stats(),
			Message: "OK",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(resp)
	})
	router.GET("/hotels/search-availabilities", SearchHotelAvailabilities)

	co := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         86400,
	})

	fmt.Println("Listening to port: 7171...")
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), co.Handler(router))
}

func SearchHotelAvailabilities(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cfg := config.GetInstance()

	limit := getRequestLimit(r)
	offset := getRequestOffset(r)
	checkinDate := r.URL.Query().Get("checkin_date")
	checkoutDate := r.URL.Query().Get("checkout_date")
	hotelIds := r.URL.Query().Get("hotel_ids")

	hotelRepo := repository.NewHotelRepo(cfg.DB)
	ac := &hotel.AccessProvider{
		HotelRepo: hotelRepo,
	}

	hotels, err, httpStatus := ac.SearchAvailabilities(limit, offset, checkinDate, checkoutDate, hotelIds)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatus)
		response := ResponseError{
			Code:    httpStatus,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	response := Response{
		Data: hotels,
	}

	json.NewEncoder(w).Encode(response)
}

//getRequestLimit returns a limit value from a http request query param
func getRequestLimit(r *http.Request) int {
	queryValues := r.URL.Query()
	limit, err := strconv.Atoi(queryValues.Get("limit"))
	if err != nil {
		limit = 10
	}

	if limit != -1 && limit < 0 {
		limit = 10
	}

	return limit
}

//getRequestOffset returns an offset value from a http request query param
func getRequestOffset(r *http.Request) int {
	queryValues := r.URL.Query()
	offset, err := strconv.Atoi(queryValues.Get("offset"))
	if err != nil {
		offset = 0
	}

	if offset != 0 && offset < 0 {
		offset = 0
	}

	return offset
}
