package entity

type NodeEntity struct {
	ID       int64   `json:"id" gorm:"primary_key"`
	Name     string  `json:"name"` // ?
	ZoneId   int64   `json:"zone_id"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	NodeType string  `json:"node_type"`
}
