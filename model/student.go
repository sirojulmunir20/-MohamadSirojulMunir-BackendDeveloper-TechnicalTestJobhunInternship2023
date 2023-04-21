package model

import "time"

type Mahasiswa struct {
	ID         int       `json:"id"`
	Nama       string    `json:"nama"`
	Usia       int       `json:"usia"`
	Gender     int       `json:"gender"`
	TanggalReg time.Time `json:"tanggal_reg"`
	Jurusan    Jurusan   `json:"jurusan"`
	Hobi       []Hobi    `json:"hobi"`
}

type Jurusan struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type Hobi struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type MahasiswaHobi struct {
	ID          int `json:"id"`
	IDMahasiswa int `json:"id_mahasiswa"`
	IDHobi      int `json:"id_hobi"`
}
