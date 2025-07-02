package main

import (
    "order_management/controllers"
    "order_management/middlewares"
    "order_management/models"

    "github.com/gin-gonic/gin"
)

func main() {
    models.ConnectDatabase()

    r := gin.Default()
    r.SetTrustedProxies(nil)

    // Serve HTML static files (frontend)
    r.Static("/static", "./static")

    // ✅ Public routes (no JWT)
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.GET("/products", controllers.GetProducts)

    // 🔒 Routes with JWT
    auth := r.Group("/")
    auth.Use(middlewares.JWTAuth())
    {
        // 🛒 Customer order
        auth.POST("/orders", controllers.CreateOrder)
        auth.GET("/orders", controllers.GetUserOrders)

        // 🔐 Admin-only routes
        admin := auth.Group("/admin")
        admin.Use(middlewares.AdminOnly())
        {
            admin.POST("/products", controllers.CreateProduct)
            admin.PUT("/products/:id", controllers.UpdateProduct)
            admin.DELETE("/products/:id", controllers.DeleteProduct)
        }
    }

    r.Run(":8000") // Jalankan server di port 8888
}
