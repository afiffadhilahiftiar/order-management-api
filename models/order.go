package models

import (
    "time"
)

type Order struct {
    ID         uint         `json:"id" gorm:"primaryKey"`
    UserID     uint         `json:"user_id"`
    User       User         `json:"user"`
    OrderItems []OrderItem  `json:"items" gorm:"foreignKey:OrderID"`
    Total      float64      `json:"total"`          // Total harga seluruh item
    CreatedAt  time.Time    `json:"created_at"`     // Untuk riwayat pesanan
}

type OrderItem struct {
    ID        uint    `json:"id" gorm:"primaryKey"`
    OrderID   uint    `json:"order_id"`
    ProductID uint    `json:"product_id"`
    Product   Product `json:"product"`              // Preload agar frontend bisa tampilkan nama/harga
    Quantity  int     `json:"quantity"`
}
