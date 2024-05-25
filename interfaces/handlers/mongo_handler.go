package handlers

import "net/http"

type MongoHandler interface {
	AddTestData(w http.ResponseWriter, r *http.Request)
	GetTestData(w http.ResponseWriter, r *http.Request)
}
