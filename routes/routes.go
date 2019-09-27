package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"graphql-boilerplate/configs"
	"graphql-boilerplate/controllers"
	"net/http"
	"os"
	"time"
)

func init() {
	r := gin.New()
	r.Use(controllers.SetHeaders)
	r.OPTIONS("/schema", controllers.OptionsHandler)
	r.POST("/schema", controllers.GQLHandler)

	s := &http.Server{
		Addr:         ":" + os.Getenv("APP_PORT"),
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	fmt.Printf(configs.NoticeColor, "поднимаю на порту "+os.Getenv("APP_PORT"))
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf(configs.ErrorColor+"%s", "Не могу сервак поднять:", err)
	}
}
