package repository

import (
	"database/sql"

	"github.com/edupooter/intensivo-golang/entity"
)

type CourseMySQLRepository struct {
	Db *sql.DB
}

func (c CourseMySQLRepository) Insert(course entity.Course) error {
	stmt, err := c.Db.Prepare(`INSERT INTO courses (id, name, description, status) values (?,?,?,?)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		course.ID,
		course.Name,
		course.Description,
		course.Status,
	)

	if err != nil {
		return err
	}

	return nil
}
