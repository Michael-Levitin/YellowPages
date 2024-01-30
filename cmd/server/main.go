package main

import (
	"context"
	"fmt"
	"github.com/Michael-Levitin/YellowPages/config"
	"github.com/Michael-Levitin/YellowPages/internal/logic"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	// гружем конфиг
	config.Init()
	sc := config.New()

	// подключаемся к базе данных
	db := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.DbUsername, sc.DbPassword, sc.DbHost, sc.DbPort, sc.DbName)
	pool, err := pgxpool.Connect(context.TODO(), db)
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	log.Println("connected to database")
	defer pool.Close()

	http.HandleFunc("/updateInfo", logic.UpdateInfo)
	log.Println("server is running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
