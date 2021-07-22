package models

import "github.com/jackc/pgtype"

type Client struct {
	Id           ClientID
	FirstName    string
	MiddleName   string
	LastName     string
	Phone        string
	RegisteredAt pgtype.Date
}

type ClientID int
