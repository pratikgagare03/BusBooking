package main

import (
	"busbooking/availability"
	"busbooking/book"
	"busbooking/utils"
	"net/http"
)

func main() {
	buses := utils.CreateBuses()
	// Register handlers

	http.HandleFunc("/check_availability", availability.Check_availability(buses)) //pass date
	http.HandleFunc("/book", book.BookingHandler(buses))                           // pass customer data, bus no, date
	// http.HandleFunc("/print", PrintBooking)                   // pass booking id
	// http.HandleFunc("/cancel", cancelBookingHandler)          // pass booking id

	// Start server
	// fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}
