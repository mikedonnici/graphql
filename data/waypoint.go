package data

import (
	"github.com/satori/uuid"
	"time"
)

// Waypoint is an intended Position
type Waypoint struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Sequence int       `json:"sequence"`
	ETA      time.Time `json:"eta"`
	ATA      time.Time `json:"ata"`
	Lat      float64   `json:"lat"`
	Lon      float64   `json:"lon"`
}

// Unpack a waypoint 'object'
func (w *Waypoint) Unpack(obj map[string]interface{}) {
	if val, ok := obj["id"].(string); ok {
		w.ID = val
	} else {
		w.ID = uuid.NewV4().String()
	}
	if val, ok := obj["name"].(string); ok {
		w.Name = val
	}
	if val, ok := obj["lat"].(float64); ok {
		w.Lat = val
	}
	if val, ok := obj["lon"].(float64); ok {
		w.Lon = val
	}
	if val, ok := obj["sequence"].(int); ok {
		w.Sequence = val
	}
	if val, ok := obj["eta"].(string); ok {
		eta, _ := time.Parse(time.RFC3339, val)
		w.ETA = eta
	}
}
