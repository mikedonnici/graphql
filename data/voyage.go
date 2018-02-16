package data

import (
	"sort"
)

type Voyage struct {
	ID          string     `json:"id"`
	VesselID    string     `json:"vesselId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Waypoints   []Waypoint `json:"waypoints"`
	Positions   []Position `json:"positions"`
}

// unpackVoyage takes the map[string]interface{} that represents a 'Voyage' object passed in
// a graphql query, and attempts to "unpack" it and assign appropriate values to the corresponding Go type.
// If a non null type is missing, the graphql server will throw an error with a useful message.
// However, for optional types the error is: "interface conversion: interface {} is nil, not string"
// because it cannot assert a missing value. In these cases set to a nil value for the appropriate type.
// Don't really need to check for NonNull values here, but easier to assume they are ALL optional and
// let graphql catch the ones that are not.
func (v *Voyage) Unpack(obj map[string]interface{}) {
	if val, ok := obj["vesselId"].(string); ok {
		v.VesselID = val
	}
	if val, ok := obj["name"].(string); ok {
		v.Name = val
	}
	if val, ok := obj["description"].(string); ok {
		v.Description = val
	}
}

// Sequence attempts to give the Waypoint values in a Voyage, a sequential number based on ETA.
func (v *Voyage) Sequence() {
	var xw waypointSort
	xw = v.Waypoints
	sort.Sort(xw)
	v.Waypoints = xw

	// Once ordered, given them a sequence number to make it easier for clients to sort them.
	for i := range v.Waypoints {
		v.Waypoints[i].Sequence = i
	}
}

// Implement the Sort interface for Waypoints
type waypointSort []Waypoint

func (p waypointSort) Len() int {
	return len(p)
}
func (p waypointSort) Less(i, j int) bool {
	return p[i].ETA.Before(p[j].ETA)
}
func (p waypointSort) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
