package models

import "github.com/jackc/pgtype"

type RentList struct {
	ClientID     ClientID
	ItemID       ItemID
	RentPriceID  RentPriceID
	Duration     int
	RentalAmount int
	StartRentAt  pgtype.Date
}

type RentListID int
