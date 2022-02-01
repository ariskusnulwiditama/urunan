package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
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

	input := user.LoginInput{
		Email:    "paul@yahudmail.com",
		Password: "foo123",
	}

	user, err := userService.Login(input)
	if err != nil {
		fmt.Println("terjadi kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run(":8000")
}
