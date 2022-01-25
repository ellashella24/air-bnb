package city

type CreateCityRequestFormat struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateCityRequestFormat struct {
	Name string `json:"name" form:"name" validate:"required"`
}
