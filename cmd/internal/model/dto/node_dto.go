package dto

import (
	"github.com/FatiaGlacier/navigation-service/cmd/internal/model/entity"
)

type NodeDTO struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"` // ?
	ZoneId   int64   `json:"zone_id"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	NodeType string  `json:"node_type"`
}

func ToDTO(e *entity.NodeEntity) *NodeDTO {
	return &NodeDTO{
		ID:       e.ID,
		Name:     e.Name,
		ZoneId:   e.ZoneId,
		X:        e.X,
		Y:        e.Y,
		NodeType: e.NodeType,
	}
}
