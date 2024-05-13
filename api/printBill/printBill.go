package printBill

import (
	"busbooking/db"
	"busbooking/logger"
	"encoding/json"
	"net/http"
	"strconv"
)

func Print(w http.ResponseWriter, r *http.Request) {
	bookingIdString := r.FormValue("bookingId")
	bookingIdInt, err := strconv.ParseInt(bookingIdString, 10, 64)
	if err != nil {
		logger.Logs.Error().Err(err)
		http.Error(w, "Error Parsing the string to Int", http.StatusBadRequest)
	}

	data, exists := db.BillIdToBill[int(bookingIdInt)]
	var msgjson []byte
	if exists {
		msgjson, _ = json.Marshal(data)
	} else {
		msgjson, _ = json.Marshal("No booking Found")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(msgjson)
}
