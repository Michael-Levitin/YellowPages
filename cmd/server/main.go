package main

import (
	"context"
	"fmt"
	"github.com/Michael-Levitin/YellowPages/config"
	"github.com/Michael-Levitin/YellowPages/internal/database"
	"github.com/Michael-Levitin/YellowPages/internal/delivery"
	"github.com/Michael-Levitin/YellowPages/internal/logic"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
)

func main() {
	// загружаем конфиг
	config.Init()
	sc := config.New()

	// подключаемся к базе данных
	dbAdrr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.DbUsername, sc.DbPassword, sc.DbHost, sc.DbPort, sc.DbName)
	db, err := pgxpool.New(context.TODO(), dbAdrr)
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	log.Println("connected to database")
	defer db.Close()

	pagesDB := database.NewPagesDB(db)                 // подключаем бд
	pagesLogic := logic.NewPagesLogic(pagesDB)         // подключаем бд к логике...
	pagesServer := delivery.NewPagesServer(pagesLogic) // ... а логику в библиотеку

	http.HandleFunc("/setInfo", pagesServer.SetInfo)
	http.HandleFunc("/getInfo", pagesServer.GetInfo)
	http.HandleFunc("/deleteInfo", pagesServer.DeleteInfo)
	log.Println("server is running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
