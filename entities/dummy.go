package entities

type Dummy struct {
	Dummy string `json:"dummy" binding:"required" validate:"dummy"`
}
