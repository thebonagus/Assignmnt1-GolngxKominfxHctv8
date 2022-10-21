package model

import "time"

type Order struct {
	OrderID      int       `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `gorm:"not null;type:varchar(50)" json:"customerName"`
	OrderedAt    time.Time `gorm:"not null;type:timestamp" json:"orderedAt"`
	Items        []Item    `gorm:"constraint:onDelete:CASCADE" json:"items"`
}

type Item struct {
	ItemID      int    `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `gorm:"not null;type:varchar" json:"itemCode"`
	Description string `gorm:"type:text" json:"description"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	OrderID     int    `json:"-"`
}
