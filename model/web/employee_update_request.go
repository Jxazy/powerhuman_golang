package web

type EmployeeUpdateRequest struct {
	Id         int    `json:"id"`
	Name       string `validate:"required,min=1,max=255" json:"name"`
	Email      string `validate:"required" json:"email"`
	Gender     string `json:"gender"`
	Age        int    `validate:"required" json:"age"`
	Phone      string `validate:"required" json:"phone"`
	TeamId     int    `validate:"required" json:"team_id"`
	RoleId     int    `validate:"required" json:"role_id"`
	IsVerified bool   `json:"is_verified"`
}
