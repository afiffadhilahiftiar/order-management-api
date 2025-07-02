Order Management API (Golang + HTML)

Sistem manajemen pemesanan produk menggunakan Golang (Gin).

✨ Fitur
- Login & Register (JWT-based)
- Role: Admin & Customer
- Produk: CRUD berlaku untuk Admin
- Pesanan: oleh Customer
- Riwayat pesanan
- GORM (ORM untuk MySQL)
- dotenv (`.env` file config)

🚀 Cara Menjalankan
1. Buat folder order_management, lalu buka
2. Inisialisasi Module: go mod init order_management
3. Install Dependency:
   ```bash
   go get github.com/gin-gonic/gin
   go get gorm.io/gorm
   go get gorm.io/driver/mysql
   go get github.com/golang-jwt/jwt/v5
   go get github.com/joho/godot

5. CREATE DATABASE orderdb;
6. Import file Database orderdb ke mysql anda
7. Import file yang ada di GitHub ke dalam folder order_management
8. Jalankan:
   ```bash
   go run main.go
9. Akses di browser: http://localhost:8000/static/

📝 Asumsi

- Hanya admin bisa tambah/edit produk
- Customer hanya bisa order
- Stok dikurangi saat order dibuat
- Tidak menggunakan database online, default: MySQL lokal
- Semua akun baru (via register) memiliki role default: customer
- Login menyimpan token & role di localStorage
- Role admin harus ditambahkan manual via database
- Tidak ada reset password; login berbasis email-password
