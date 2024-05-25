package handlers

import (
	"context"
	"github.com/RandySteven/go-kopi/enums"
	"github.com/RandySteven/go-kopi/interfaces/handlers"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type MongoHandler struct {
}

func (m *MongoHandler) AddTestData(w http.ResponseWriter, r *http.Request) {
	var (
		rID = uuid.NewString()
		ctx = context.WithValue(r.Context(), enums.RequestID, rID)
	)
	log.Println(ctx)
}

func (m *MongoHandler) GetTestData(w http.ResponseWriter, r *http.Request) {
}

func NewMongoHandler() *MongoHandler {
	return &MongoHandler{}
}

var _ handlers.MongoHandler = &MongoHandler{}
