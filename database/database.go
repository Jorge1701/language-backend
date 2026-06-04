package database

import (
	"database/sql"
	"errors"
	"fmt"
	"language-backend/model"
	"math/rand"
	"strings"
)

const (
	db_file = "language.db"
)

var (
	ErrNotFoundInDB = errors.New("Not found in DB")
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("sqlite", db_file)
	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func (db *Database) Close() {
	db.db.Close()
}

func (db *Database) FindConjugation(tense model.Tense, pronoun model.Pronoun, verb string) (string, error) {
	query := fmt.Sprintf("SELECT %s FROM verbs WHERE infinitive = ?", getColumnName(tense, pronoun))
	row := db.db.QueryRow(query, verb)

	var conjugation string
	err := row.Scan(&conjugation)

	return conjugation, mapError(err)
}

func getColumnName(tense model.Tense, pronoun model.Pronoun) string {
	p := ""
	switch pronoun {
	case model.FirstPersonalSingular:
		p = "fps"
	case model.SecondPersonalSingular:
		p = "sps"
	case model.ThirdPersonalSingular:
		p = "tps"
	case model.FirstPersonalPlural:
		p = "fpp"
	case model.SecondPersonalPlural:
		p = "spp"
	case model.ThirdPersonalPlural:
		p = "tpp"
	}
	return fmt.Sprintf("%s_%s", strings.ToLower(string(tense)), p)
}

func (db *Database) FindVerb(verb string) (VerbRow, error) {
	query := fmt.Sprintf("SELECT %s FROM verbs WHERE infinitive = ?", verbsProperties())
	row := db.db.QueryRow(query, verb)

	return extractVerbRow(row)
}

func (db *Database) ListVerbs(limit int) ([]string, error) {
	query := `SELECT infinitive FROM verbs LIMIT ?`
	rows, err := db.db.Query(query, limit)
	if err != nil {
		return nil, err
	}

	verbs := []string{}
	for rows.Next() {
		var verb string
		if err := rows.Scan(&verb); err != nil {
			return nil, err
		}
		verbs = append(verbs, verb)
	}

	return verbs, rows.Err()
}

func (db *Database) FindRandomVerb(topSearched int) (VerbRow, error) {
	verbId := rand.Intn(topSearched) + 1

	query := fmt.Sprintf("SELECT %s FROM verbs WHERE id = ?", verbsProperties())
	row := db.db.QueryRow(query, verbId)

	return extractVerbRow(row)
}
