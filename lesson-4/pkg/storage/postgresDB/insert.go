package postgresDB

import (
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"postgre-dev-go/internal/models"
	"time"
)

type PostgresDB struct {
	dbpool *pgxpool.Pool
}
func NewPostgresDB(dbpool *pgxpool.Pool) *PostgresDB {
	return &PostgresDB{dbpool}
}

// InsertClient добавляент нового клиента в БД. Ей на вход мы передаём заполненую данными структуру Client.
// Поле MiddleName (Отчество) является не обязательным, его можно не заполнять.
func (pg *PostgresDB)InsertClient(ctx context.Context, client models.Client) (models.ClientID, error) {
	var (
		id models.ClientID
		sql string
		err error
	)
	if client.MiddleName != "" {
		sql = `insert into client (first_name, middle_name, last_name, phone) values ($1, $2, $3, $4)
		 on conflict (phone) do update set phone=excluded.phone returning id;`
		err = pg.dbpool.QueryRow(ctx, sql, client.FirstName, client.MiddleName, client.LastName, client.Phone).Scan(&id)
	}
	if client.MiddleName == "" {
		//200013 DELETE FROM client WHERE id = 200012; 200025
		sql = `insert into client (first_name, last_name, phone) values ($1, $2, $3)
		on conflict (phone) do update set phone=excluded.phone returning id;`
		err = pg.dbpool.QueryRow(ctx, sql, client.FirstName, client.LastName, client.Phone).Scan(&id)
	}
	if err != nil {
		return id, fmt.Errorf("failed to insert client: %w", err)
	}

	return id, nil
}


// InsertItem добавляет запись в БД о новой вещи которую могут оформить в аренду.
// Связанную таблицу (rent_price) о стоимости и условиях аренды можно заполнить позже.
func (pg *PostgresDB)InsertItem(ctx context.Context, item models.Item) (*models.ItemID, error) {
	const sql = `insert into item (name, description) values ($1, $2) returning id;`

	var id models.ItemID
	err := pg.dbpool.QueryRow(ctx, sql,
		item.Name,
		item.Description,
	).Scan(&id)
	if err != nil {
		return &id, fmt.Errorf("failed to insert item: %w", err)
	}
	return &id, nil
}

func (pg *PostgresDB)updateItemExpiresAt(ctx context.Context, id models.ItemID, expiresAT pgtype.Date) error {
	const sql = `UPDATE item SET expires_at = $2 WHERE id = $1;`
	_, err := pg.dbpool.Exec(ctx, sql,
		id,
		expiresAT.Time,
	)
	if err != nil {
		return fmt.Errorf("failed to update expires_at in item id = %d: %w", id, err)
	}
	return nil
}

// InsertRentPrice добавляет запись в таблицу (rent_price) связанную с таблицей (item) о стоимости и условиях аренды.
func (pg *PostgresDB)InsertRentPrice(ctx context.Context, rentPrice models.RentPrice, itemID models.ItemID) (*models.RentPriceID, error) {
	const sql = `insert into rent_price (item_id, name, price) values ($1, $2, $3) returning id;`

	var id models.RentPriceID
	err := pg.dbpool.QueryRow(ctx, sql,
		itemID,
		rentPrice.Name,
		rentPrice.Price,
	).Scan(&id)
	if err != nil {
		return &id, fmt.Errorf("failed to insert rent_price: %w", err)
	}
	return &id, nil
}

// InsertRentlist при оформлени аренды делает запись в БД информаци о клиенте, о вещи о стоимости, о сроке аренды.
// Информация о сроке завершения аренды вноситься в таблицу (item) в поле expires_at функцией updateItemExpiresAt.
func (pg *PostgresDB)InsertRentlist(ctx context.Context, rentList models.RentList) (*models.RentListID, error) {
	const sql = `insert into rent_list (client_id, item_id, rent_price_id, duration, rental_amount) values ($1, $2, $3, $4, $5) returning id;`

	var id models.RentListID
	err := pg.dbpool.QueryRow(ctx, sql,
		rentList.ClientID,
		rentList.ItemID,
		rentList.RentPriceID,
		rentList.Duration,
		rentList.RentalAmount,
	).Scan(&id)
	if err != nil {
		return &id, fmt.Errorf("failed to insert rent_list: %w", err)
	}

	expiresAT := pgtype.Date{
		Time:   time.Now().Add(time.Duration(rentList.Duration) * time.Hour),
		Status: pgtype.Present,
	}
	err = pg.updateItemExpiresAt(ctx, rentList.ItemID, expiresAT)
	if err != nil {
		return &id, err
	}
	return &id, nil
}
