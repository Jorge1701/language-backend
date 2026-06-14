package endpoints

import (
	"language-backend/database"
)

type Handler struct {
	db     *database.Database
	isProd bool
}

func NewEndpointHandler(db *database.Database, isProd bool) *Handler {
	return &Handler{
		db:     db,
		isProd: isProd,
	}
}
