package endpoints

import (
	"language-backend/database"
	"language-backend/generator"
)

type Handler struct {
	db        *database.Database
	generator *generator.Generator
	isProd    bool
}

func NewEndpointHandler(db *database.Database, generator *generator.Generator, isProd bool) *Handler {
	return &Handler{
		db:        db,
		generator: generator,
		isProd:    isProd,
	}
}
