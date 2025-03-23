package course

type (
	CreateCourseDTO struct {
		Name      string `json:"name"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}
	UpdateCourseDTO struct {
		ID        string  `json:"id"`
		Name      *string `json:"name"`
		StartDate *string `json:"start_date"`
		EndDate   *string `json:"end_date"`
	}
)
