package usecase

type CreateCourseInputDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateCourseOutputDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
