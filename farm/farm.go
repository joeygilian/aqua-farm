package farm

type Farm struct {
	ID   int64  `json:"id"`
	Name string `json:"name" validate:"required"`
}
