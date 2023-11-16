package web

type EmployeeResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	Phone      string `json:"phone"`
	Photo      string `json:"photo"`
	TeamId     int    `json:"team_id"`
	RoleId     int    `json:"role_id"`
	IsVerified bool   `json:"is_verified"`
}
