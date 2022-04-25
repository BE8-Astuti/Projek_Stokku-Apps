package user

type InsertUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `gorm:"unique" json:"password" form:"password"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type UpdateUserRequest struct {
	Email string `json:"email" validate:"required"`
}
