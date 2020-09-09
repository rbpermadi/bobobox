package bobobox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rbpermadi/bobobox/internal/config"
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

	co := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         86400,
	})

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), co.Handler(router))
}
