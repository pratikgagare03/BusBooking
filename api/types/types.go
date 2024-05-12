package types

type Bus struct {
	No           int
	Seats        int
	Source       string
	Dest         string
	Fare         float64
	BoardingTime string
	DroppingTime string
	BookedStatus map[string]bool
}

type Customer struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

type Booking struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Date    string `json:"date"`
	Source  string `json:"source"`
	Dest    string `json:"dest"`
}

// Agency name: xyz travels, Date:
// Bus number:
// Booked seats:
// Source stop:pune destination:mumbai
// ------------------------
// Customer name: Contact:
// ------------------------
// Total Fare:
type Bill struct {
	BookingId    int     `json:"id"`
	Bus_Number   int     `json:"busno"`
	Booked_Seats int     `json:"bookedseats"`
	Source       string  `json:"source"`
	Dest         string  `json:"dest"`
	BookingDate  string  `json:"bookingdate"`
	JounrneyDate string  `json:"journeydate"`
	CustomerName string  `json:"customername"`
	Contact      string  `json:"contact"`
	TotalFare    float64 `json:"totalfare"`
}
