package repository

import (
	"context"
	"database/sql"
	"errors"
	"jobhun-api/helper"
	"jobhun-api/model"
)

type StudentRepositoryImpl struct {
	Db *sql.DB
}

func NewStudentRepository(Db *sql.DB) StudentRepository {
	return &StudentRepositoryImpl{Db: Db}
}

// Delete implements StudentRepository
func (b *StudentRepositoryImpl) Delete(ctx context.Context, studentId int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "DELETE FROM mahasiswa WHERE id = ?"
	_, errExec := tx.ExecContext(ctx, SQL, studentId)
	helper.PanicIfError(errExec)
}

// FindAll implements StudentRepository
func (b *StudentRepositoryImpl) FindAll(ctx context.Context) []model.Mahasiswa {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id,nama FROM mahasiswa"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var students []model.Mahasiswa

	for result.Next() {
		student := model.Mahasiswa{}
		err := result.Scan(&student.ID, &student.Nama)
		helper.PanicIfError(err)

		students = append(students, student)
	}

	return students
}

// FindById implements StudentRepository
func (b *StudentRepositoryImpl) FindById(ctx context.Context, studentId int) (model.Mahasiswa, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id,nama FROM book WHERE = ?"
	result, errQuery := tx.QueryContext(ctx, SQL, studentId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	student := model.Mahasiswa{}

	if result.Next() {
		err := result.Scan(&student.ID, &student.Nama)
		helper.PanicIfError(err)
		return student, nil
	} else {
		return student, errors.New("student id not found")
	}
}

// Save implements StudentRepository
func (b *StudentRepositoryImpl) Save(ctx context.Context, student model.Mahasiswa) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "INSERT INTO mahasiswa (nama, usia, gender, tanggal_reg, id_jurusan) VALUES (?, ?, ?, ?, ?)"
	res, err := tx.ExecContext(ctx, SQL, student.Nama, student.Usia, student.Gender, student.TanggalReg, student.Jurusan.ID)
	helper.PanicIfError(err)

	// Ambil ID mahasiswa yang baru saja di-insert
	_,err = res.LastInsertId()
	helper.PanicIfError(err)
	
	// Insert data hobi mahasiswa ke database
	for _, h := range student.Hobi {
		SQL := "INSERT INTO mahasiswa_hobi (id_mahasiswa, id_hobi) VALUES (?, ?)"
		_,err = tx.ExecContext(ctx, SQL, student.ID, h.ID)
		helper.PanicIfError(err)
	}
}

// Update implements StudentRepository
func (b *StudentRepositoryImpl) Update(ctx context.Context, student model.Mahasiswa) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "UPDATE book SET nama = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, student.Nama, student.ID)
	helper.PanicIfError(err)
}



