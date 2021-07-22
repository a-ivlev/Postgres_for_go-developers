// +build integration

package postgresDB_test

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
	"os"
	"postgre-dev-go/internal/models"
	"postgre-dev-go/pkg/storage/postgresDB"
	"testing"
)

// Соединение с экземпляром Postgres
func connect(ctx context.Context) *pgxpool.Pool {
	dbpool, err := pgxpool.Connect(ctx, os.Getenv("POSTGRES_DSN"))
	if err != nil {
		panic(err)
	}
	return dbpool
}

func TestPostgresDB_InsertClient(t *testing.T) {
	ctx := context.Background()
	dbpool := connect(ctx)
	defer dbpool.Close()

	tests := []struct {
		name string
		store *postgresDB.PostgresDB
		ctx	context.Context
		client models.Client
		check func(*testing.T, models.ClientID, error)
	}{
		{
			name: "success",
			store: postgresDB.NewPostgresDB(dbpool),
			ctx: context.Background(),
			client: models.Client{
				FirstName: "Егор",
				MiddleName: "Алексевич",
				LastName: "Саблин",
				Phone: "+79225555353",
			},
			check: func(t *testing.T, hint models.ClientID, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, hint)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			hint, err := tt.store.InsertClient(tt.ctx, tt.client)
			tt.check(t, hint, err)
		})
	}
}

func TestPostgresDB_InsertItem(t *testing.T) {
	ctx := context.Background()
	dbpool := connect(ctx)
	defer dbpool.Close()

	tests := []struct {
		name string
		store *postgresDB.PostgresDB
		ctx	context.Context
		item models.Item
		check func(*testing.T, *models.ItemID, error)
	}{
		{
			name: "success",
			store: postgresDB.NewPostgresDB(dbpool),
			ctx: context.Background(),
			item: models.Item{
				Name: "Вентилятор напольный ZANUSSI ZFF-705",
				Description: `Описание
					Вентилятор Zanussi ZFF-705 напольный`,
			},
			check: func(t *testing.T, hint *models.ItemID, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, hint)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			hint, err := tt.store.InsertItem(tt.ctx, tt.item)
			tt.check(t, hint, err)
		})
	}
}

func TestPostgresDB_InsertRentPrice(t *testing.T) {
	ctx := context.Background()
	dbpool := connect(ctx)
	defer dbpool.Close()

	tests := []struct {
		name string
		store *postgresDB.PostgresDB
		ctx	context.Context
		rentPrice models.RentPrice
		itemID models.ItemID
		check func(*testing.T, *models.RentPriceID, error)
	}{
		{
			name: "success",
			store: postgresDB.NewPostgresDB(dbpool),
			ctx: context.Background(),
			rentPrice: models.RentPrice{
				Name: "Тест - Аренда с оплатой за час.",
				Price: 10,
			},
			itemID: 1,
			check: func(t *testing.T, hint *models.RentPriceID, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, hint)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			hint, err := tt.store.InsertRentPrice(tt.ctx, tt.rentPrice, tt.itemID)
			tt.check(t, hint, err)
		})
	}
}

func TestPostgresDB_InsertRentlist(t *testing.T) {
	ctx := context.Background()
	dbpool := connect(ctx)
	defer dbpool.Close()

	tests := []struct {
		name string
		store *postgresDB.PostgresDB
		ctx	context.Context
		rentList models.RentList
		check func(*testing.T, *models.RentListID, error)
	}{
		{
			name: "success",
			store: postgresDB.NewPostgresDB(dbpool),
			ctx: context.Background(),
			rentList: models.RentList{
				ClientID: 1,
				ItemID: 1,
				RentPriceID: 1,
				Duration: 2,
				RentalAmount: 10,
			},
			check: func(t *testing.T, hint *models.RentListID, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, hint)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			hint, err := tt.store.InsertRentlist(tt.ctx, tt.rentList)
			tt.check(t, hint, err)
		})
	}
}


func TestPostgresDB_SearchClientByPhone(t *testing.T) {
	ctx := context.Background()
	dbpool := connect(ctx)
	defer dbpool.Close()

	tests := []struct {
		name string
		store *postgresDB.PostgresDB
		ctx	context.Context
		phone string
		prepare func(*pgxpool.Pool)
		check func(*testing.T, *models.Client, error)
	}{
		{
			name: "success",
			store: postgresDB.NewPostgresDB(dbpool),
			ctx: context.Background(),
			phone: "+79991112233",
			prepare: func(dbpool *pgxpool.Pool) {
				// Подготовка тестовых данных.
				dbpool.Exec(context.Background(), `INSERT INTO client (first_name, last_name, phone)
				VALUES ('Тест', 'Тест', '+79991112233') on conflict DO NOTHING;`)
			},
			check: func(t *testing.T, hint *models.Client, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, hint)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(dbpool)
			hint, err := tt.store.SearchClientByPhone(tt.ctx, tt.phone)
			tt.check(t, hint, err)
		})
	}
}

func TestPostgresDB_SearchClientByLastName(t *testing.T) {
	ctx := context.Background()
	dbpool := connect(ctx)
	defer dbpool.Close()

	tests := []struct {
		name string
		store *postgresDB.PostgresDB
		ctx	context.Context
		lastName string
		prepare func(*pgxpool.Pool)
		check func(*testing.T, []models.Client, error)
	}{
		{
			name: "success",
			store: postgresDB.NewPostgresDB(dbpool),
			ctx: context.Background(),
			lastName: "Тест",
			prepare: func(dbpool *pgxpool.Pool) {
				// Подготовка тестовых данных.
				dbpool.Exec(context.Background(), `INSERT INTO client (first_name, last_name, phone)
				VALUES ('Тест', 'Тест', '+79991112233') on conflict DO NOTHING;`)
			},
			check: func(t *testing.T, hints []models.Client, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, hints)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(dbpool)
			hints, err := tt.store.SearchClientByLastName(tt.ctx, tt.lastName)
			tt.check(t, hints, err)
		})
	}
}