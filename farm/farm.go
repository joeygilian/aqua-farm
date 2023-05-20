package farm

import pond "github.com/aqua-farm/pond"

type Farm struct {
	ID   int64  `json:"id"`
	Name string `json:"name" validate:"required"`
	Pond []*pond.Pond
}
