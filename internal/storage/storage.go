package storage

import "github.com/MohitKumar2217/Students-api/internal/types"

// interface
type Storage interface {
	CreateStudent(name string, email string, age int, enroll string) (int64, error) 
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
}
