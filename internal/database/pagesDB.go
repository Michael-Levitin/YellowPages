package database

import (
	"context"
	"fmt"
	"github.com/Michael-Levitin/YellowPages/internal/dto"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

const (
	_setInfoQuery = `
INSERT INTO people_data  (name, surname, patronymic, age, sex, country)
VALUES (@name, @surname, @patronymic, @age, @sex, @country)
RETURNING id;`
	_getInfoQuery = `
SELECT id, name, surname, patronymic, age, sex, country
FROM people_data
WHERE`
	_delInfoQuery = `
DELETE from people_data
WHERE`
	_retId = `
RETURNING id;
`
	_retAll = `
RETURNING id, name, surname, patronymic, age, sex, country;`
)

type PagesDB struct {
	db *pgxpool.Pool
}

func NewPagesDB(db *pgxpool.Pool) *PagesDB {
	return &PagesDB{db: db}
}

func (p PagesDB) GetInfo(ctx context.Context, info *dto.Info) (*[]dto.Info, error) {
	log.Trace().Msg(fmt.Sprintf("DB recieve %+v\n", info))
	query := _getInfoQuery + dto.Info2String(info) + " Order by id;"
	log.Trace().Msg(fmt.Sprintf("query: ", query))

	rows, err := p.db.Query(context.TODO(),
		query,
		pgx.NamedArgs(dto.Info2map(info)))
	if err != nil {
		log.Debug().Msg(fmt.Sprintf("Could not query %+v, %s", info, err))
		return &[]dto.Info{}, err
	}

	people, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.Info])
	if err != nil {
		log.Debug().Msg(fmt.Sprintf("CollectRows error: %v", err))
		return &[]dto.Info{}, err
	}
	return &people, nil
}

func (p PagesDB) SetInfo(ctx context.Context, info *dto.Info) (*dto.Info, error) {
	var id int
	err := p.db.QueryRow(context.TODO(), _setInfoQuery, pgx.NamedArgs(dto.Info2map(info))).Scan(&id)
	info.Id = id
	if err != nil {
		log.Debug().Msg(fmt.Sprint("Could not add ", info, err))
		return &dto.Info{}, err
	}
	return info, nil
}

func (p PagesDB) DeleteInfo(ctx context.Context, info *dto.Info) (*[]dto.Info, error) {
	where := dto.Info2String(info)
	if where == " true " {
		log.Info().Msg("empty clause atempt")
		return &[]dto.Info{}, fmt.Errorf("clause cannot be empty")
	}
	rows, err := p.db.Query(context.TODO(),
		_delInfoQuery+where+_retAll,
		pgx.NamedArgs(dto.Info2map(info)))
	if err != nil {
		log.Debug().Msg(fmt.Sprint("Could not query %+v, %s", info, err))
		return &[]dto.Info{}, err
	}

	people, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.Info])
	if err != nil {
		log.Debug().Msg(fmt.Sprint("CollectRows error: %v", err))
		return &[]dto.Info{}, err
	}
	return &people, nil
}
