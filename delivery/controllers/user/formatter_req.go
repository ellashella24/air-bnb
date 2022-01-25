package user

type LoginFormatRequest struct {
	Email    string `json:"email" form:"email" validate:"email,required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type RegisterFormatRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"email,required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateFormatRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"email,required"`
	Password string `json:"password" form:"password" validate:"required"`
}
