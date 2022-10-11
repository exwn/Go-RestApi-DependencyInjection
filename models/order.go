package models

import "time"

type Order struct {
	Order_id      uint      `gorm:"primaryKey" json:"id"`
	Customer_name string    `gorm:"not null; type:varchar(25)" json:"customerName"`
	Ordered_at    time.Time `json:"orderedAt"`
	Items         []Item    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items"`
}
