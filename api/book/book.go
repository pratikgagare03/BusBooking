package book

import (
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
			panic(err)
		}

		bill := types.Bill{}
		bookingDone := false
		for _, val := range *buses {
			if !val.BookedStatus[booking.Date] && val.Source == booking.Source {
				val.BookedStatus[booking.Date] = true
				bill.Bus_Number = val.No
				bill.BookingDate = time.Now().Local().Weekday().String() + " " + time.Now().Format(("2006-01-02 15:04:05"))
				bill.JounrneyDate = booking.Date
				bill.CustomerName = booking.Name
				bill.Contact = booking.Contact
				bill.TotalFare = val.Fare * float64(val.Seats)
				bill.Source = booking.Source
				bill.Dest = booking.Dest
				bill.Booked_Seats = val.Seats
				bookingDone = true

				break
			}
		}
		if bookingDone {
			billjson, _ := json.Marshal(bill)
			w.Header().Set("Content-Type", "application/json")
			w.Write(billjson)
		} else {
			w.Header().Set("Content-Type", "application/json")
			op, _ := json.Marshal("No Buses Available Sorry for inconvenience occured")
			w.Write(op)
		}
	}
}
