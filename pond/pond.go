package pond

type Pond struct {
	ID     int64  `json:"id"`
	FarmID int64  `json:"farmid"`
	Name   string `json:"name" validate:"required"`
}
