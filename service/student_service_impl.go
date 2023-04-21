package service

import (
	"context"
	"jobhun-api/data/request"
	"jobhun-api/data/response"
	"jobhun-api/helper"
	"jobhun-api/model"
	"jobhun-api/repository"
)

type StudentServiceImpl struct {
	StudentRepository repository.StudentRepository
}

func NewStudentServiceImpl(studentRepository repository.StudentRepository) StudentService {
	return &StudentServiceImpl{StudentRepository: studentRepository}
}

// Create implements StudentService
func (s *StudentServiceImpl) Create(ctx context.Context, request request.StudentCreateRequest) {
	student := model.Mahasiswa{
		Nama: request.Nama,
	}
	s.StudentRepository.Save(ctx, student)
}

// Delete implements StudentService
func (s *StudentServiceImpl) Delete(ctx context.Context, studentId int) {
	student, err := s.StudentRepository.FindById(ctx, studentId)
	helper.PanicIfError(err)
	s.StudentRepository.Delete(ctx, student.ID)
}

// FindAll implements StudnetService
func (s *StudentServiceImpl) FindAll(ctx context.Context) []response.StudentResponse {
	students := s.StudentRepository.FindAll(ctx)

	var studentResp []response.StudentResponse

	for _, value := range students {
		student := response.StudentResponse{ID: value.ID, Nama: value.Nama}
		studentResp = append(studentResp, student)
	}
	return studentResp
}

// FindById implements StudentService
func (s *StudentServiceImpl) FindById(ctx context.Context, studentId int) response.StudentResponse {
	student, err := s.StudentRepository.FindById(ctx, studentId)
	helper.PanicIfError(err)
	return response.FromStudent(student)
}

// Update implements StudentService
func (s *StudentServiceImpl) Update(ctx context.Context, request request.StudentUpdateRequest) {
	student, err := s.StudentRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	student.Nama = request.Nama
	s.StudentRepository.Update(ctx, student)
}