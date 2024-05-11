package model

import "time"

type Reservation struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	RoomID    uint      `json:"room_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	UserID    uint      `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
