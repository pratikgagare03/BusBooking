package types

type Bus struct {
	No             int             `json:"busno"`
	TotalSeats     int             `json:"totalseats"`
	AvailableSeats int             `json:"availableseats"`
	Source         string          `json:"source"`
	Dest           string          `json:"dest"`
	Fare           float64         `json:"fare"`
	BoardingTime   string          `json:"boardingTime"`
	DroppingTime   string          `json:"droppingTime"`
	BookedStatus   map[string]bool `json:"-"`
}

type Customer struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

type Booking struct {
	Name        string `json:"name"`
	Contact     string `json:"contact"`
	JourneyDate string `json:"date"`
	Source      string `json:"source"`
	Dest        string `json:"dest"`
	NoOfSeats   int    `json:"seats"`
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
