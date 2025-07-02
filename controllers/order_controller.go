package controllers

import (
    "net/http"
    "order_management/models"
    "order_management/utils"

    "github.com/gin-gonic/gin"
)

type OrderRequest struct {
    Items []struct {
        ProductID uint `json:"product_id"`
        Quantity  int  `json:"quantity"`
    } `json:"items" binding:"required,dive"`
}

// POST /orders
func CreateOrder(c *gin.Context) {
    var req OrderRequest
    if err := c.ShouldBindJSON(&req); err != nil || len(req.Items) == 0 {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Input pesanan tidak valid"))
        return
    }

    userID := c.GetUint("userID")

    var orderItems []models.OrderItem
    var total float64
    tx := models.DB.Begin()

    for _, item := range req.Items {
        var product models.Product
        if err := tx.First(&product, item.ProductID).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusNotFound, utils.ErrorResponse("Produk tidak ditemukan"))
            return
        }

        if product.Stock < item.Quantity {
            tx.Rollback()
            c.JSON(http.StatusBadRequest, utils.ErrorResponse("Stok tidak mencukupi: "+product.Name))
            return
        }

        // Kurangi stok
        product.Stock -= item.Quantity
        if err := tx.Save(&product).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal mengurangi stok"))
            return
        }

        subtotal := float64(item.Quantity) * product.Price
        total += subtotal

        orderItems = append(orderItems, models.OrderItem{
            ProductID: item.ProductID,
            Quantity:  item.Quantity,
        })
    }

    order := models.Order{
        UserID:     userID,
        OrderItems: orderItems,
        Total:      total,
    }

    if err := tx.Create(&order).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal membuat pesanan"))
        return
    }

    tx.Commit()
    c.JSON(http.StatusCreated, utils.SuccessResponse("Pesanan berhasil dibuat"))
}

// GET /orders
func GetUserOrders(c *gin.Context) {
    userID := c.GetUint("userID")
    var orders []models.Order

    if err := models.DB.
        Preload("OrderItems.Product").
        Where("user_id = ?", userID).
        Order("created_at desc").
        Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal mengambil riwayat pesanan"))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(orders))
}
