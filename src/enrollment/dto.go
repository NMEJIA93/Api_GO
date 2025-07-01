package enrollment

import "github.com/NMEJIA93/Api_GO/pkg/meta"

type (
	CreateEnrollmentDTO struct {
		UserID   string `json:"user_id"`
		CourseID string `json:"course_id"`
	}

	CreateRequest struct {
		UserID   string `json:"user_id"`
		CourseID string `json:"course_id"`
	}
	Response struct {
		Status int64       `json:"status"`
		Data   interface{} `json:"data"`
		Err    string      `json:"error,omitempty"`
		Meta   *meta.Meta  `json:"meta,omitempty"`
	}
)
