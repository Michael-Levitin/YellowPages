package main

import (
	"context"
	"fmt"
	"github.com/Michael-Levitin/YellowPages/config"
	"github.com/Michael-Levitin/YellowPages/internal/database"
	"github.com/Michael-Levitin/YellowPages/internal/delivery"
	"github.com/Michael-Levitin/YellowPages/internal/logic"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	// загружаем конфиг
	config.Init()
	sc := config.New()
	//logger := zerolog.New(os.Stdout)
	zerolog.SetGlobalLevel(sc.LogLevel)

	// подключаемся к базе данных
	dbAdrr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.DbUsername, sc.DbPassword, sc.DbHost, sc.DbPort, sc.DbName)
	db, err := pgxpool.New(context.TODO(), dbAdrr)
	if err != nil {
		log.Fatal().Err(err).Msg("error connecting to database")
	}
	log.Info().Msg("connected to database")
	defer db.Close()

	pagesDB := database.NewPagesDB(db)                 // подключаем бд
	pagesLogic := logic.NewPagesLogic(pagesDB)         // подключаем бд к логике...
	pagesServer := delivery.NewPagesServer(pagesLogic) // ... а логику в библиотеку

	http.HandleFunc("/setInfo", pagesServer.SetInfo)
	http.HandleFunc("/getInfo", pagesServer.GetInfo)
	http.HandleFunc("/deleteInfo", pagesServer.DeleteInfo)
	http.HandleFunc("/updateInfo", pagesServer.UpdateInfo)
	log.Info().Msg("server is running...")
	err = http.ListenAndServe(":8080", nil)
	log.Fatal().Err(err).Msg("http server crashed")
}
