package controllers

import (
    "net/http"
    "order_management/models"
    "order_management/utils"
    "order_management/middlewares"

    "github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
    var input models.RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Input tidak valid"))
        return
    }

    // Cek duplikasi email
    var existing models.User
    if err := models.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Email sudah terdaftar"))
        return
    }

    hashedPassword, _ := models.HashPassword(input.Password)
    user := models.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: hashedPassword,
        Role:     "customer",
    }

    if err := models.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal membuat akun"))
        return
    }

    c.JSON(http.StatusCreated, utils.SuccessResponse(gin.H{
        "message": "Pendaftaran berhasil",
        "user":    gin.H{"name": user.Name, "email": user.Email, "role": user.Role},
    }))
}

func Login(c *gin.Context) {
    var input models.LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Input tidak valid"))
        return
    }

    var user models.User
    if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Email tidak ditemukan"))
        return
    }

    if !models.CheckPasswordHash(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Password salah"))
        return
    }

    token, err := middlewares.GenerateJWT(user.ID, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal membuat token"))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(gin.H{
        "token": token,
        "user":  gin.H{"name": user.Name, "email": user.Email, "role": user.Role},
    }))
}
