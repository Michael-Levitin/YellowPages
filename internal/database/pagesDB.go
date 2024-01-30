package database

import (
	"context"
	"github.com/Michael-Levitin/YellowPages/internal/dto"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const (
	_setInfoQuery = `INSERT INTO people_data  (name, surname, patronymic, age, sex, nationality)
VALUES (@name, @surname, @patronymic, @age, @sex, @nationality)
RETURNING id`
)

type PagesDB struct {
	db *pgxpool.Pool
}

func NewPagesDB(db *pgxpool.Pool) *PagesDB {
	return &PagesDB{db: db}
}

func (p PagesDB) GetInfo(ctx context.Context, info dto.Info) (dto.Info, error) {
	return dto.Info{}, nil
}

func (p PagesDB) SetInfo(ctx context.Context, info dto.Info) (dto.Info, error) {
	return dto.Info{}, nil
}

func SetInfo(db *pgxpool.Pool) {
	args := pgx.NamedArgs{
		"name":        "Anton",
		"surname":     "Sidorov",
		"patronymic":  "Olegovich",
		"age":         "26",
		"sex":         "male",
		"nationality": "RU"}

	_, err := db.Exec(context.TODO(), _setInfoQuery, args)
	if err != nil {
		log.Println("Could not add ", args, err)
	}
}
