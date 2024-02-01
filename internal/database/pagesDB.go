package database

import (
	"context"
	"fmt"
	"github.com/Michael-Levitin/YellowPages/internal/dto"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const (
	_setInfoQuery = `INSERT INTO people_data  (name, surname, patronymic, age, sex, nationality)
VALUES (@name, @surname, @patronymic, @age, @sex, @country)
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
	fmt.Println(dto.Info2map(info))
	_, err := p.db.Exec(context.TODO(), _setInfoQuery, pgx.NamedArgs(dto.Info2map(info)))
	if err != nil {
		log.Println("Could not add ", info, err)
		return dto.Info{}, err
	}
	return info, nil
}
