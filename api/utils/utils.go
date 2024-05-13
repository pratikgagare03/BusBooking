package utils

import (
	"busbooking/types"
	"math/rand"
	"time"
)

func ValidDate(date string) (time.Time, error) {
	parsedDate, err := time.Parse("02-01-2006", date)
	return parsedDate, err
}

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
			No:                   i + 1,
			TotalSeats:           40,
			AvailableSeatsOnDate: make(map[string]int),
			Source:               src,
			Dest:                 dest,
			Fare:                 float64(rand.Intn(500-400) + 400),
			BoardingTime:         Currtime.Format("15:04"),
			DroppingTime:         Currtime.Add(time.Hour * 5).Format("15:04"),
		})

	}

	return &buses

}
