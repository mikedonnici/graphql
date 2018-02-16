package data

import "time"

// Position is a reported vessel location
type Position struct {
	ID       string    `json:"id"`
	VesselID string    `json:"vesselId"`
	ATA      time.Time `json:"ata"`
	Sequence int       `json:"sequence"`
	Lat      float64   `json:"lat"`
	Lon      float64   `json:"lon"`
}

// Unpack an object into a value of type Position
func (p *Position) Unpack(obj map[string]interface{}) {
	if val, ok := obj["vesselId"].(string); ok {
		p.VesselID = val
	}
	if val, ok := obj["lat"].(float64); ok {
		p.Lat = val
	}
	if val, ok := obj["lon"].(float64); ok {
		p.Lon = val
	}
	if val, ok := obj["ata"].(string); ok {
		ata, _ := time.Parse(time.RFC3339, val)
		p.ATA = ata
	}
}
