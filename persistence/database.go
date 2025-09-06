package persistence

import (
	"context"
	"fon/config"
	"log"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Lesson struct {
	ID        int8   `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	VideoLink string `json:"video_link"`
}

var conn *pgxpool.Pool

func Start() {
	var err error
	conn, err = pgxpool.New(context.Background(), config.GetEnv("SUPABASE_URL"))
	if err != nil {
		log.Fatalf("Supabase connection failed! %v", err)
	}
}

func GetLessonByID(id int) *Lesson {
	var lesson Lesson
	query := `SELECT * FROM lessons WHERE id = $1`

	err := pgxscan.Get(context.Background(), conn, &lesson, query, id)
	if err != nil {
		log.Printf("Error querying lesson by ID %d: %v", id, err)
		return nil
	}
	return &lesson
}
