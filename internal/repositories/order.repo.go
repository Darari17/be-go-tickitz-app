package repositories

import (
	"context"

	"github.com/Darari17/be-go-tickitz-app/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{db: db}
}

func (or *OrderRepo) CreateOrder(ctx context.Context, order *models.Order, seatIDs []int) (*models.Order, error) {
	tx, err := or.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO orders (qr_code, users_id, schedules_id, payments_id, fullname, email, phone_number, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,NOW())
		RETURNING id, created_at
	`
	err = tx.QueryRow(ctx, query,
		order.QRCode, order.UserID, order.ScheduleID, order.PaymentID,
		order.FullName, order.Email, order.Phone,
	).Scan(&order.ID, &order.CreatedAt)
	if err != nil {
		return nil, err
	}

	for _, seatID := range seatIDs {
		_, err := tx.Exec(ctx, `INSERT INTO orders_seats (orders_id, seats_id) VALUES ($1,$2)`, order.ID, seatID)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return order, nil
}

func (or *OrderRepo) GetOrderByID(ctx context.Context, id int) (*models.Order, error) {
	var order models.Order
	query := `
		SELECT id, qr_code, users_id, schedules_id, payments_id,
		       fullname, email, phone_number, created_at, updated_at
		FROM orders WHERE id=$1
	`
	err := or.db.QueryRow(ctx, query, id).Scan(
		&order.ID, &order.QRCode, &order.UserID, &order.ScheduleID,
		&order.PaymentID, &order.FullName, &order.Email,
		&order.Phone, &order.CreatedAt, &order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	rows, err := or.db.Query(ctx, `
		SELECT s.id, s.seat_code
		FROM seats s
		INNER JOIN orders_seats os ON os.seats_id = s.id
		WHERE os.orders_id = $1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var seat models.Seat
		if err := rows.Scan(&seat.ID, &seat.SeatCode); err != nil {
			return nil, err
		}
		order.Seats = append(order.Seats, seat)
	}

	return &order, nil
}

func (or *OrderRepo) GetOrdersByUserID(ctx context.Context, userID int) ([]models.Order, error) {
	rows, err := or.db.Query(ctx, `
		SELECT id, qr_code, users_id, schedules_id, payments_id,
		       fullname, email, phone_number, created_at, updated_at
		FROM orders WHERE users_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(
			&order.ID, &order.QRCode, &order.UserID, &order.ScheduleID,
			&order.PaymentID, &order.FullName, &order.Email,
			&order.Phone, &order.CreatedAt, &order.UpdatedAt,
		); err != nil {
			return nil, err
		}

		seatRows, err := or.db.Query(ctx, `
			SELECT s.id, s.seat_code
			FROM seats s
			INNER JOIN orders_seats os ON os.seats_id = s.id
			WHERE os.orders_id = $1
		`, order.ID)
		if err != nil {
			return nil, err
		}
		defer seatRows.Close()

		for seatRows.Next() {
			var seat models.Seat
			if err := seatRows.Scan(&seat.ID, &seat.SeatCode); err != nil {
				return nil, err
			}
			order.Seats = append(order.Seats, seat)
		}

		orders = append(orders, order)
	}

	return orders, nil
}
