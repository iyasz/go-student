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