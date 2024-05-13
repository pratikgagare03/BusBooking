package availability

import (
	"busbooking/logger"
	"busbooking/types"
	"busbooking/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func Check_availability(buses *[]types.Bus) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.FormValue("date")
		seats := r.FormValue("seats")
		_, err := utils.ValidDate(date)
		if err != nil {
			logger.Logs.Error().Err(err)
			http.Error(w, "Recieved Invalid date", http.StatusBadRequest)
			return
		}

		var seatsInt int64
		if seats == "" {
			seatsInt = 0
		} else {
			seatsInt, _ = strconv.ParseInt(seats, 10, 64)
		}

		var op []types.CheckAvailabilityOp
		insuffSeats := false
		for _, bus := range *buses {
			var op1 types.CheckAvailabilityOp
			availableSeatsOnDate, exists := bus.AvailableSeatsOnDate[date]
			if exists {
				if availableSeatsOnDate < int(seatsInt) {
					insuffSeats = true
					continue
				}
				op1.AvailableSeats = bus.AvailableSeatsOnDate[date]
			} else if seatsInt > int64(bus.TotalSeats) {
				insuffSeats = true
				continue
			} else {
				op1.AvailableSeats = bus.TotalSeats
			}

			op1.No = bus.No
			op1.TotalSeats = bus.TotalSeats
			op1.Fare = bus.Fare
			op1.Source = bus.Source
			op1.Dest = bus.Dest
			op1.BoardingTime = bus.BoardingTime
			op1.DroppingTime = bus.DroppingTime
			op = append(op, op1)
		}
		var opjson []byte
		if len(op) != 0 {
			opjson, _ = json.Marshal(op)
		} else {
			if insuffSeats {
				opjson, _ = json.Marshal("This many seats are not available in single bus, you can try booking in separate Buses")
			} else {
				opjson, _ = json.Marshal("No buses available for selected date")
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(opjson)
	}
}
