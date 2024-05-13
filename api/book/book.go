package book

import (
	"busbooking/db"
	"busbooking/logger"
	"busbooking/types"
	"busbooking/utils"
	"encoding/json"
	"net/http"
	"time"
)

func BookingHandler(buses *[]types.Bus) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		booking := types.Booking{}
		err := json.NewDecoder(r.Body).Decode(&booking)
		if err != nil {
			logger.Logs.Error().Err(err)
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		parsedDate, err := utils.ValidDate(booking.JourneyDate)
		if err!=nil{
			logger.Logs.Error().Err(err)
			http.Error(w, "Recieved Invalid date", http.StatusBadRequest)
			return
		}

		validFutureDate := parsedDate.Compare(time.Now())
		if validFutureDate == -1 {
			logger.Logs.Error().Msg("Dates before today's date not allowed")
			http.Error(w, "Past date not allowed", http.StatusBadRequest)
			return
		}
		bill := types.Bill{}
		bookingDone := false
		for _, bus := range *buses {
			if bus.Source == booking.Source && bus.Dest == booking.Dest {
				availableSeatsOnDate, exists := bus.AvailableSeatsOnDate[booking.JourneyDate]
				if exists {
					if availableSeatsOnDate < booking.NoOfSeats {
						continue
					} else {
						bus.AvailableSeatsOnDate[booking.JourneyDate] -= booking.NoOfSeats
					}

				} else {
					bus.AvailableSeatsOnDate[booking.JourneyDate] = bus.TotalSeats - booking.NoOfSeats
				}
				bill.Bus_Number = bus.No
				bill.BookingDate = time.Now().Local().Weekday().String() + " " + time.Now().Format(("2006-01-02 15:04:05"))
				bill.JounrneyDate = booking.JourneyDate
				bill.CustomerName = booking.Name
				bill.Contact = booking.Contact
				bill.TotalFare = bus.Fare * float64(booking.NoOfSeats)
				bill.Source = booking.Source
				bill.Dest = booking.Dest
				bill.Booked_Seats = booking.NoOfSeats
				bookingDone = true
				break
			}
		}
		if bookingDone {
			BillId := db.GetBillId()
			bill.BookingId = BillId
			db.BillIdToBill[BillId] = bill
			w.Header().Set("Content-Type", "application/json")
			billjson, _ := json.Marshal(bill)
			w.Write(billjson)
		} else {
			w.Header().Set("Content-Type", "application/json")
			op, _ := json.Marshal("No Buses Available Sorry for inconvenience occured")
			w.Write(op)
		}
	}
}
