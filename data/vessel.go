package data

import "sort"

// Vessel is vehicle capable of undertaking a Voyage
type Vessel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CallSign    string `json:"callSign"`
	Type        string `json:"type""`
	Description string `json:"description"`
}

// VesselPosition struct holds the vessel and a list of Positions
type VesselPositions struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	CallSign    string     `json:"callSign"`
	Type        string     `json:"type""`
	Description string     `json:"description"`
	Positions   []Position `json:"positions"`
}

// unpack a waypoint 'object' into a Waypoint value
func (v *Vessel) Unpack(obj map[string]interface{}) {
	if val, ok := obj["callSign"].(string); ok {
		v.CallSign = val
	}
	if val, ok := obj["name"].(string); ok {
		v.Name = val
	}
	if val, ok := obj["type"].(string); ok {
		v.Type = val
	}
	if val, ok := obj["description"].(string); ok {
		v.Description = val
	}
}

// Sequence orders and gives a sequence number to the positions in a VesselPositions value
func (vp *VesselPositions) Sequence() {
	var xp positionSort
	xp = vp.Positions
	sort.Sort(xp)
	vp.Positions = xp
	// Once ordered, given them a sequence number
	for i := range vp.Positions {
		vp.Positions[i].Sequence = i
	}

}

// Implement the Sort interface for Positions
type positionSort []Position

func (p positionSort) Len() int {
	return len(p)
}
func (p positionSort) Less(i, j int) bool {
	return p[i].ATA.Before(p[j].ATA)
}
func (p positionSort) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
