package repository

import (
	"context"
	"jobhun-api/model"
)

type StudentRepository interface {
	Save(ctx context.Context, student model.Mahasiswa)
	Update(ctx context.Context, student model.Mahasiswa)
	Delete(ctx context.Context, studentId int)
	FindById(ctx context.Context, studentId int) (model.Mahasiswa, error)
	FindAll(ctx context.Context) []model.Mahasiswa
}