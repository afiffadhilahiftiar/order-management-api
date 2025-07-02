package models

import (
    "time"
    "gorm.io/gorm"
)

type Product struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Name      string         `json:"name" binding:"required"`
    Stock     int            `json:"stock" binding:"required,gte=0"`
    Price     float64        `json:"price" binding:"required,gt=0"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Digunakan saat input produk dari admin (POST/PUT)
type ProductInput struct {
    Name  string  `json:"name" binding:"required"`
    Stock int     `json:"stock" binding:"required,gte=0"`
    Price float64 `json:"price" binding:"required,gt=0"`
}
