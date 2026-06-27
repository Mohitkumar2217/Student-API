package storage

// interface
type Storage interface {
	CreateStudent(name string, email string, age int, enroll string) (int64, error) 
}

