package main

import (
	"fmt"
	"graphql-boilerplate/configs"
	"graphql-boilerplate/db"
	_ "graphql-boilerplate/routes"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Gracefully stop application
	var gracefulStop = make(chan os.Signal)

	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		fmt.Printf("Caught sig: %+v\n", sig)
		if err := db.Conn.Close(); err != nil {
			fmt.Println("cannot close db connection")
		}
		fmt.Printf(configs.NoticeColor, "To be or to have? Make your choice please... I die but not give up.")
		os.Exit(0)
	}()

	// поднимаем и тестируем коннект с базой
	err := db.Conn.Ping()
	if err != nil {
		fmt.Printf(configs.ErrorColor, "Нет соединения с базой данных.")
	}

	fmt.Printf(configs.NoticeColor, "Всё прекрасно, дорогая маркиза. Работаем.")
}
