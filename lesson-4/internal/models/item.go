package models

import "github.com/jackc/pgtype"

type Item struct {
	Name        string
	Description string
	ExpiresAt pgtype.Date
}

type ItemID int
