package usecase

import (
	"github.com/edupooter/intensivo-golang/entity"
	"github.com/google/uuid"
)

type CreateCourse struct {
	Repository entity.CourseRepository
}

func (c CreateCourse) Execute(input CreateCourseInputDto) (CreateCourseOutputDto, error) {
	var course = entity.Course{}
	course.ID = uuid.New().String()
	course.Name = input.Name
	course.Description = input.Description
	course.Status = input.Status

	var err = c.Repository.Insert(course)

	if err != nil {
		return CreateCourseOutputDto{}, err
	}

	output := CreateCourseOutputDto{}
	output.ID = course.ID
	output.Name = input.Name
	output.Description = input.Description
	output.Status = input.Status

	return output, nil
}
