package book

import (
	"busbooking/db"
	"busbooking/logger"
	"busbooking/types"
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

		parsedDate, err := time.Parse("02-01-2006", booking.JourneyDate)
		if err != nil {
			logger.Logs.Error().Err(err)
			http.Error(w, "Not a valid date", http.StatusBadRequest)
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
		for _, val := range *buses {
			if !val.BookedStatus[booking.JourneyDate] && val.Source == booking.Source && val.Dest == booking.Dest {
				val.BookedStatus[booking.JourneyDate] = true
				bill.Bus_Number = val.No
				bill.BookingDate = time.Now().Local().Weekday().String() + " " + time.Now().Format(("2006-01-02 15:04:05"))
				bill.JounrneyDate = booking.JourneyDate
				bill.CustomerName = booking.Name
				bill.Contact = booking.Contact
				bill.TotalFare = val.Fare * float64(val.TotalSeats)
				bill.Source = booking.Source
				bill.Dest = booking.Dest
				bill.Booked_Seats = val.TotalSeats
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
