package controllers

import (
    "log"
    "net/http"
    "order_management/models"
    "order_management/utils"

    "github.com/gin-gonic/gin"
)

// GET /products — publik
func GetProducts(c *gin.Context) {
    var products []models.Product
    if err := models.DB.Find(&products).Error; err != nil {
        log.Println("[GET /products] Gagal mengambil produk:", err)
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal mengambil produk"))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(products))
}

// POST /admin/products — hanya admin
func CreateProduct(c *gin.Context) {
    var input models.ProductInput
    if err := c.ShouldBindJSON(&input); err != nil {
        log.Println("[POST /admin/products] Input tidak valid:", err)
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Input tidak valid"))
        return
    }

    product := models.Product{
        Name:  input.Name,
        Stock: input.Stock,
        Price: input.Price,
    }

    if err := models.DB.Create(&product).Error; err != nil {
        log.Println("[POST /admin/products] Gagal menyimpan produk:", err)
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal menyimpan produk"))
        return
    }

    log.Printf("[POST /admin/products] Produk ditambahkan: %s", product.Name)
    c.JSON(http.StatusCreated, utils.SuccessResponse(product))
}

// PUT /admin/products/:id — update produk
func UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product

    if err := models.DB.First(&product, id).Error; err != nil {
        log.Printf("[PUT /admin/products/%s] Produk tidak ditemukan: %v", id, err)
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Produk tidak ditemukan"))
        return
    }

    var input models.ProductInput
    if err := c.ShouldBindJSON(&input); err != nil {
        log.Printf("[PUT /admin/products/%s] Input tidak valid: %v", id, err)
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Input tidak valid"))
        return
    }

    product.Name = input.Name
    product.Stock = input.Stock
    product.Price = input.Price

    if err := models.DB.Save(&product).Error; err != nil {
        log.Printf("[PUT /admin/products/%s] Gagal memperbarui produk: %v", id, err)
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal memperbarui produk"))
        return
    }

    log.Printf("[PUT /admin/products/%s] Produk diperbarui: %s", id, product.Name)
    c.JSON(http.StatusOK, utils.SuccessResponse(product))
}

// DELETE /admin/products/:id — hapus produk
func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product

    if err := models.DB.First(&product, id).Error; err != nil {
        log.Printf("[DELETE /admin/products/%s] Produk tidak ditemukan: %v", id, err)
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Produk tidak ditemukan"))
        return
    }

    if err := models.DB.Delete(&product).Error; err != nil {
        log.Printf("[DELETE /admin/products/%s] Gagal menghapus produk: %v", id, err)
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal menghapus produk"))
        return
    }

    log.Printf("[DELETE /admin/products/%s] Produk berhasil dihapus: %s", id, product.Name)
    c.JSON(http.StatusOK, utils.SuccessResponse("Produk berhasil dihapus"))
}
