package data

import (
	"github.com/satori/uuid"
)

func AddVessel(v Vessel) Vessel {
	v.ID = uuid.NewV4().String()
	DataVessels = append(DataVessels, v)
	return v
}

func AddVoyage(v Voyage) Voyage {
	v.ID = uuid.NewV4().String()
	DataVoyages = append(DataVoyages, v)
	return v
}

func AddPosition(p Position) Position {
	p.ID = uuid.NewV4().String()
	DataPositions = append(DataPositions, p)
	return p
}

func GetVessels() []Vessel {
	return DataVessels
}

func GetVessel(id string) Vessel {
	for _, v := range DataVessels {
		if v.ID == id {
			return v
		}
	}
	return Vessel{}
}

func GetVesselPositions(vesselID string) VesselPositions {
	vp := VesselPositions{}
	// Populate the Vessel fields
	v := GetVessel(vesselID)
	vp.ID = v.ID
	vp.Name = v.Name
	vp.CallSign = v.CallSign
	vp.Description = v.Description
	vp.Type = v.Type
	// Get all positions for the vessel
	for _, p := range DataPositions {
		if p.VesselID == vesselID {
			vp.Positions = append(vp.Positions, p)
		}
	}
	vp.Sequence()
	return vp
}

func GetVoyages() []Voyage {
	return DataVoyages
}

func GetPositions() []Position {
	return DataPositions
}
