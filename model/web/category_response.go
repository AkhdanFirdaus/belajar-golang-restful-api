package web

type CategoryResponse struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required,max=255,min=1" json:"name"`
}
