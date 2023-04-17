package studentmodel

import (
	"go-student/config"
	"go-student/entities"
)

func GetAll() []entities.Student {
	rows, err := config.DB.Query(`SELECT * FROM siswa`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var students []entities.Student

	for rows.Next() {
		var student entities.Student
		if err := rows.Scan(&student.Id, &student.Nama, &student.Nis, &student.Telp, &student.CreatedAt); err != nil {
			panic(err)
		}

		students = append(students, student)
	}

	return students
}

func Store(student entities.Student) bool {
	result, err := config.DB.Exec(`
		INSERT INTO siswa (nama, nis, telp, created_at)
		VALUE (? , ?, ?, ?)`,
		student.Nama, student.Nis, student.Telp, student.CreatedAt,
	)

	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return LastInsertId > 0
}

func Detail(id int) entities.Student {
	row := config.DB.QueryRow(`SELECT id, nama, nis, telp FROM siswa WHERE id = ?`, id)

	var student entities.Student
	if err := row.Scan(&student.Id, &student.Nama, &student.Nis, &student.Telp); err != nil {
		panic(err.Error())
	}

	return student
}

func Update(id int, student entities.Student) bool {
	query, err := config.DB.Exec(`UPDATE siswa SET nama = ?, nis = ?, telp = ? WHERE id = ?`, student.Nama, student.Nis, student.Telp, id)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()

	if err != nil {
		panic(err)
	}
	
	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM siswa WHERE id = ?`, id)
	return err
}