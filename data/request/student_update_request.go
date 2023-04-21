package request

type StudentUpdateRequest struct {
	Id   int
	Nama string `validate:"required min=1,max=100" json:"nama"`
}