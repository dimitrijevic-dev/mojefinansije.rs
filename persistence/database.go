package persistence

import (
	"context"
	"fon/config"
	"log"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmailAttempt struct {
	Email string `json:"email"`
}

type LessonAttempt struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	VideoLink string `json:"video_link"`
}

type Lesson struct {
	ID        int8   `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	VideoLink string `json:"video_link"`
}

type User struct {
	Email       string `json:"email"`
	Displayname string `json:"displayname"`
	Password    string `json:"password"`
}

type UserSignup struct {
	Email       string `json:"email"`
	Displayname string `json:"displayname"`
	Password    string `json:"password"`
}

type UserAttempt struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var conn *pgxpool.Pool

func Start() {
	var err error
	conn, err = pgxpool.New(context.Background(), config.GetEnv("SUPABASE_URL"))
	if err != nil {
		log.Fatalf("Supabase connection failed! %v", err)
	}
}

func GetAll() []Lesson {
	var lessons []Lesson
	query := `SELECT * FROM lessons ORDER BY id ASC`

	err := pgxscan.Select(context.Background(), conn, &lessons, query)
	if err != nil {
		log.Printf("Error in query!")
		return nil
	}

	return lessons
}

func GetBySearch(search string) []Lesson {
	var lessons []Lesson
	query := `SELECT * FROM lessons WHERE title LIKE $1`

	searchPattern := "%" + search + "%"
	err := pgxscan.Select(context.Background(), conn, &lessons, query, searchPattern)
	if err != nil {
		log.Println("Error in query! ", err)
		return nil
	}

	return lessons
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

func GetUser(email string, password string) *User {
	var user User
	query := `SELECT * FROM users WHERE email=$1 AND password=$2`

	err := pgxscan.Get(context.Background(), conn, &user, query, email, config.Sha256(password))
	if err != nil {
		log.Printf("User doesnt exist in DB! %v", err)
	}

	return &user
}

func AddUser(attempt UserSignup) error {
	query := `INSERT INTO users (email, displayname, password) VALUES ($1,$2,$3)`

	_, err := conn.Exec(context.Background(), query, attempt.Email, attempt.Displayname, config.Sha256(attempt.Password))
	return err
}

func GetReadLessons(email string) []Lesson {
	query := `SELECT l.id, l.title, l.body, l.video_link FROM lessons l 
	INNER JOIN user_lessons u ON l.id = u.lessonid WHERE u.email = $1`

	var arr []Lesson
	err := pgxscan.Select(context.Background(), conn, &arr, query, email)
	if err != nil {
		log.Printf("Error reading read lessons: %v", err)
		return nil
	}

	return arr
}

func ReadLesson(attempt EmailAttempt, id int) error {
	query := `INSERT INTO user_lessons (email, lessonid) VALUES ($1, $2)`

	_, err := conn.Exec(context.Background(), query, attempt.Email, id)
	return err
}

func AddLesson(attempt LessonAttempt) error {
	query := `INSERT INTO lessons (title,body,video_link) VALUES ($1,$2,$3)`

	_, err := conn.Exec(context.Background(), query, attempt.Title, attempt.Body, attempt.VideoLink)
	return err
}

func CountAll() int {
	var count int
	query := `SELECT COUNT(*) FROM lessons`
	err := pgxscan.Get(context.Background(), conn, &count, query)
	if err != nil {
		log.Printf("Error counting lessons: %v", err)
		return 0
	}

	return count
}
