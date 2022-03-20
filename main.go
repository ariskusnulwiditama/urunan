package main

import (
	"fmt"
	"urunan/auth"
	"urunan/handler"
	"urunan/user"

	// "fmt"
	"log"

	// "net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:localhost@tcp(127.0.0.1:13306)/bwa-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	
	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.JnTpSCBf41jzuELlysZx6pJODT5G4gIK4VNBVqyos54")
	if err != nil {
		fmt.Println("Error")
	}
	if token.Valid {
		fmt.Println("token valid")
	}else{
		fmt.Println("token invalid")
	}

	userHandler := handler.NewUserHandler(userService, authService)


	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run(":8000")
}
