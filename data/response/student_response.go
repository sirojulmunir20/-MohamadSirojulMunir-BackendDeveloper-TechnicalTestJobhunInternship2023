package response

import "jobhun-api/model"

type StudentResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

func FromStudent(mahasiswa model.Mahasiswa) StudentResponse {
	return StudentResponse{
		ID:   mahasiswa.ID,
		Nama: mahasiswa.Nama,

		// tambahkan properti lainnya sesuai kebutuhan
	}
}