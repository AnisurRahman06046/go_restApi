package storage

import "github.com/AnisurRahman06046/go_restApi/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudent(id int64, name string, email string, age int) (types.Student, error)
	DeleteStudentById(id int64) (int64, error)
}
