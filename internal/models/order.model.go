package models

import "time"

type Order struct {
	ID         int        `db:"id" json:"id"`
	QRCode     string     `db:"qr_code" json:"qr_code"`
	UserID     int        `db:"users_id" json:"user_id" example:"1"`
	ScheduleID int        `db:"schedules_id" json:"schedule_id" example:"10"`
	PaymentID  int        `db:"payments_id" json:"payment_id" example:"2"`
	FullName   string     `db:"fullname" json:"fullname" example:"John Doe"`
	Email      string     `db:"email" json:"email" example:"johndoe@mail.com"`
	Phone      string     `db:"phone_number" json:"phone" example:"08123456789"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at"`
	Seats      []Seat     `db:"-" json:"seats"`
}

type OrderSeat struct {
	OrderID int `db:"orders_id" json:"order_id"`
	SeatID  int `db:"seats_id" json:"seat_id"`
}

type Seat struct {
	ID       int    `db:"id" json:"id"`
	SeatCode string `db:"seat_code" json:"seat_code"`
}

type PaymentMethod struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type CreateOrderRequest struct {
	Order   Order `json:"order"`
	SeatIDs []int `json:"seat_ids"`
}
