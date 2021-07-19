package postgresDB

import (
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"postgre-dev-go/internal/models"
	"time"
)

func InsertClient(ctx context.Context, dbpool *pgxpool.Pool, client models.Client) (models.ClientID, error) {
	var (
		id models.ClientID
		sql string
		err error
	)
	if client.MiddleName != "" {
		sql = `insert into client (first_name, middle_name, last_name, phone) values ($1, $2, $3, $4) returning id;`
		err = dbpool.QueryRow(ctx, sql, client.FirstName, client.MiddleName, client.LastName, client.Phone).Scan(&id)
	}
	if client.MiddleName == "" {
		//200013 DELETE FROM client WHERE id = 200012;
		sql = `insert into client (first_name, last_name, phone) values ($1, $2, $3) returning id;`
		err = dbpool.QueryRow(ctx, sql, client.FirstName, client.LastName, client.Phone).Scan(&id)
	}
	if err != nil {
		return id, fmt.Errorf("failed to insert client: %w", err)
	}

	return id, nil
}

func InsertItem(ctx context.Context, dbpool *pgxpool.Pool, item models.Item) (models.ItemID, error) {
	const sql = `insert into item (name, description) values ($1, $2) returning id;`

	var id models.ItemID
	err := dbpool.QueryRow(ctx, sql,
		item.Name,
		item.Description,
	).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("failed to insert item: %w", err)
	}
	return id, nil
}

func updateItemExpiresAt(ctx context.Context, dbpool *pgxpool.Pool, id models.ItemID, expiresAT pgtype.Date) error {
	const sql = `UPDATE item SET expires_at = $2 WHERE id = $1;`
	_, err := dbpool.Exec(ctx, sql,
		id,
		expiresAT.Time,
	)
	if err != nil {
		return fmt.Errorf("failed to update expires_at in item id = %d: %w", id, err)
	}
	return nil
}

func InsertRentPrice(ctx context.Context, dbpool *pgxpool.Pool, rentPrice models.RentPrice, itemID models.ItemID) (models.RentPriceID, error) {
	const sql = `insert into rent_price (item_id, name, price) values ($1, $2, $3) returning id;`

	var id models.RentPriceID
	err := dbpool.QueryRow(ctx, sql,
		itemID,
		rentPrice.Name,
		rentPrice.Price,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert rent_price: %w", err)
	}
	return id, nil
}

func InsertRentlist(ctx context.Context, dbpool *pgxpool.Pool, rentList models.RentList) (models.RentListID, error) {
	const sql = `insert into rent_list (client_id, item_id, rent_price_id, duration, rental_amount) values ($1, $2, $3, $4, $5) returning id;`

	var id models.RentListID
	err := dbpool.QueryRow(ctx, sql,
		rentList.ClientID,
		rentList.ItemID,
		rentList.RentPriceID,
		rentList.Duration,
		rentList.RentalAmount,
	).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("failed to insert rent_list: %w", err)
	}

	expiresAT := pgtype.Date{
		Time:   time.Now().Add(time.Duration(rentList.Duration) * time.Hour),
		Status: pgtype.Present,
	}
	err = updateItemExpiresAt(ctx, dbpool, rentList.ItemID, expiresAT)
	if err != nil {
		return id, err
	}
	return id, nil
}
