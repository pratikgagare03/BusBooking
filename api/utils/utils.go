package utils

import (
	"busbooking/types"
	"time"
)

func CreateBuses() *[]types.Bus {
	buses := make([]types.Bus, 0, 10)
	Currtime := time.Now()
	for i := 0; i < 10; i++ {
		Currtime = Currtime.Add(time.Hour * 2)
		var src, dest string
		if i%2 == 0 {
			src = "Pune"
			dest = "Mumbai"
		} else {
			src = "Mumbai"
			dest = "Pune"
		}
		buses = append(buses, types.Bus{
			No:             i + 1,
			TotalSeats:     40,
			AvailableSeats: 40,
			Source:         src,
			Dest:           dest,
			Fare:           500,
			BoardingTime:   Currtime.Format("15:04"),
			DroppingTime:   Currtime.Add(time.Hour * 5).Format("15:04"),
			BookedStatus:   make(map[string]bool),
		})
	}

	return &buses

}
