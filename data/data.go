package data

import (
	"time"
)

// dummy data
var DataVessels = []Vessel{
	{
		ID:          "a7f4e2fd-6747-4180-a0bb-9db14a517c7c",
		Name:        "Icarus",
		CallSign:    "JB-101",
		Type:        "Traditional 30",
		Description: "Sloop, white hull, white topsides",
	},
	{
		ID:          "af44fdfd-4184-4b47-8296-b90eef262309",
		Name:        "Daedalus",
		CallSign:    "JB-222",
		Type:        "Mason 55",
		Description: "Ketch, white hull, white topsides",
	},
	{
		ID:          "23989670-8da5-4f3a-bdca-da7143cf764d",
		Name:        "IgglePiggle",
		CallSign:    "JB-145",
		Type:        "Redjacket",
		Description: "Sloop, red hull, white topsides",
	},
}

var t, _ = time.Parse(time.RFC3339, "2017-11-02T06:04:05+10:00")

var DataVoyages = []Voyage{
	{
		ID:          "21af638c-de93-4f05-8000-a0b9f18a024b",
		VesselID:    "a7f4e2fd-6747-4180-a0bb-9db14a517c7c",
		Name:        "JB to Port Mac",
		Description: "JB -> Cronulla -> Port Stephens -> Port Macquarie",
		Waypoints: []Waypoint{
			{
				ID:       "c64949f7-6faa-48b8-ae73-796876eb6964",
				Name:     "Point Perp.",
				Sequence: 1,
				ETA:      t,
				ATA:      t,
				Lat:      -35.099220,
				Lon:      150.803005,
			},
			{
				ID:       "a08146f8-826b-46db-b62a-5c849115d361",
				Name:     "Lamond Head",
				Sequence: 2,
				ETA:      t,
				Lat:      -35.050142,
				Lon:      150.850169,
			},
			{
				ID:       "921a0d91-4e72-4ad0-9ea8-3cb31815462b",
				Name:     "Beecroft Head",
				Sequence: 3,
				ETA:      t,
				Lat:      -35.013395,
				Lon:      150.853153,
			},
		},
	},
}

var DataPositions = []Position{}
