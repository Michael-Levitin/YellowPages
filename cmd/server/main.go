package main

import (
	"context"
	"fmt"
	"github.com/Michael-Levitin/YellowPages/config"
	"github.com/Michael-Levitin/YellowPages/internal/database"
	"github.com/Michael-Levitin/YellowPages/internal/logic"
	"github.com/jackc/pgx/v5/pgxpool"
	//_ "github.com/lib/pq"
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

	database.SetInfo(db)

	http.HandleFunc("/updateInfo", logic.UpdateInfo)
	log.Println("server is running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
