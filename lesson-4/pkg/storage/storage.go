package storage

import (
	"context"
	"postgre-dev-go/internal/models"
	"time"
)

type RentItems struct {
	RentListID    models.RentListID
	Firstname     string
	Lastname      string
	ItemID        models.ItemID
	ItemName      string
	RentPriceName string
	RentDuration  int
	RentAmount    int
	StartRent     time.Time
	ExpiresAt     time.Time
}


type ClientsAndItems struct {
	RentListID    models.RentListID
	Firstname     string
	Lastname      string
	ItemID        models.ItemID
	ItemName      string
	RentPriceName string
	RentDuration  int
	RentAmount    int
	StartRent     time.Time
	ExpiresAt     time.Time
	NowAt         time.Time
}

// Storage скрыл реализацию БД Postgres за интерфейсом. Представление уровня хранилища.
type Storage interface {
	InsertClient(ctx context.Context, client models.Client) (*models.ClientID, error)
	InsertItem(ctx context.Context, item models.Item) (*models.ItemID, error)
	InsertRentPrice(ctx context.Context, rentPrice models.RentPrice, itemID models.ItemID) (*models.RentPriceID, error)
	InsertRentlist(ctx context.Context, rentList models.RentList) (*models.RentListID, error)
	SearchClientByPhone(ctx context.Context, phone string) (*models.Client, error)
	SearchClientByLastName(ctx context.Context, lastName string) ([]models.Client, error)
	SearchRentItemsByPhone(ctx context.Context, phone string) ([]RentItems, error)
	SearchClientsNotReturnItemsOnTime(ctx context.Context) ([]ClientsAndItems, error)
	SearchClientsRentalAndReturnItemsDatePeriod(ctx context.Context, startDate time.Time, endDate time.Time) ([]ClientsAndItems, error)
}

