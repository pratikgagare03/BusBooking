package cancel

import (
	"busbooking/db"
	"busbooking/logger"
	"busbooking/types"
	"encoding/json"
	"net/http"
	"strconv"
)

func CancelBooking(buses *[]types.Bus) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookingIdString := r.FormValue("bookingId")
		bookingIdInt, err := strconv.ParseInt(bookingIdString, 10, 64)
		if err != nil {
			logger.Logs.Error().Err(err)
			http.Error(w, "Error Parsing the string to Int", http.StatusBadRequest)
		}

		data, exists := db.BillIdToBill[int(bookingIdInt)]
		var msgjson []byte
		if exists {
			delete((*buses)[data.Bus_Number-1].BookedStatus, data.JounrneyDate)
			delete(db.BillIdToBill, data.BookingId)
			msgjson, _ = json.Marshal("Booking Cancelled Successful")
		} else {
			msgjson, _ = json.Marshal("No booking Found")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(msgjson)
	}
}
