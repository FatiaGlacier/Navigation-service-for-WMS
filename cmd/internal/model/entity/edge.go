package entity

type EdgeEntity struct {
	ID        int64   `json:"id" gorm:"primary_key"`
	FromID    int64   `json:"from_id"`
	ToID      int64   `json:"to_id"`
	Distance  float64 `json:"distance"`
	Angle     float64 `json:"angle"`
	IsReverse bool    `json:"is_reverse"`
}
