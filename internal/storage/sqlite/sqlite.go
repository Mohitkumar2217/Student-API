package sqlite

import (
	"database/sql"

	"github.com/MohitKumar2217/Students-api/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	age INTEGER,
	email TEXT,
	enroll TEXT
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int, enroll string) (int64, error) {
	smt, err := s.Db.Prepare("INSERT INTO Students (name, email, age, enroll) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer smt.Close()

	result, err := smt.Exec(name, email, age, enroll)
	if err != nil {
		return 0, err
	}

	lastid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastid, nil
}
