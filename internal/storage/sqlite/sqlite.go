package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/MohitKumar2217/Students-api/internal/config"
	"github.com/MohitKumar2217/Students-api/internal/types"
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

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM Students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Age, &student.Email, &student.Enroll)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("No student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query Error: %w", err)
	}

	return student, nil
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, age, email, enroll from Students")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Email, &student.Enroll)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (s *Sqlite) DeleteStudent(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM Students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(
		&student.Id,
		&student.Name,
		&student.Age,
		&student.Email,
		&student.Enroll,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %d", id)
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}

	stmt2, err := s.Db.Prepare("DELETE FROM Students WHERE id = ?")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt2.Close()

	_, err = stmt2.Exec(id)
	if err != nil {
		return types.Student{}, fmt.Errorf("delete error: %w", err)
	}

	return student, nil
}