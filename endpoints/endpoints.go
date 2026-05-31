package endpoints

import (
	"language-backend/database"
)

type Handler struct {
	db *database.Database
}

func NewEndpointHandler(db *database.Database) *Handler {
	return &Handler{
		db: db,
	}
}
