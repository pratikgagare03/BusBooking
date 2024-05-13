package db

import "busbooking/types"

var BillId int = 0
var BillIdToBill = make(map[int]types.Bill)

func GetBillId() int{
	BillId += 1
	return BillId
}