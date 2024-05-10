package availability

import (
	"busbooking/types"
	"encoding/json"
	"net/http"
)

func Check_availability(buses *[]types.Bus) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.FormValue("date")
		var op []types.Bus

		for _, val := range *buses {
			if !val.BookedStatus[date] {
				op = append(op, val)
			}
		}
		var opjson []byte
		if len(op) != 0 {
			opjson, _ = json.Marshal(op)
		} else {
			opjson, _ = json.Marshal("No buses available for selected date")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(opjson)
	}
}
