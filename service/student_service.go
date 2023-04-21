package service

import (
	"context"
	"jobhun-api/data/request"
	"jobhun-api/data/response"
)

type StudentService interface {
	Create(ctx context.Context, request request.StudentCreateRequest)
	Update(ctx context.Context, request request.StudentUpdateRequest)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) response.StudentResponse
	FindAll(ctx context.Context) []response.StudentResponse
}