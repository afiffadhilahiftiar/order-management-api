package models

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"

    "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
    // Muat file .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Gagal memuat file .env")
    }

    // Ambil variabel dari .env
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    name := os.Getenv("DB_NAME")

    // Format DSN
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, pass, host, port, name,
    )

    // Koneksi ke database
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal koneksi database: ", err)
    }

    fmt.Println("Database terkoneksi...")

    DB = database

    // Auto-migrate semua model
    DB.AutoMigrate(&User{}, &Product{}, &Order{}, &OrderItem{})
}
