package dtos

type TerminalRequest struct {
	Name string `json:"name" form:"name" validate:"required,gte=2,lte=128"`
}