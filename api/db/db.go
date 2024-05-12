package db

import "busbooking/types"

var Id int = 1
var BillIdToBill = make(map[int]types.Bill)

func GetId() *int {
	Id = Id + 1
	return &Id
}

func GetBillMap() *map[int]types.Bill {
	return &BillIdToBill
}
