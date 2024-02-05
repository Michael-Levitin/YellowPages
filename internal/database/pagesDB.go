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
INSERT INTO people_data  (name, surname, patronymic, age, gender, country)
VALUES (@name, @surname, @patronymic, @age, @gender, @country)
RETURNING id;`
	_getInfoQuery = `
SELECT id, name, surname, patronymic, age, gender, country
FROM people_data
WHERE`
	_getEnd = `
ORDER BY id
LIMIT @limit OFFSET @offset;`
	_delInfoQuery = `
with deleted as (
    DELETE from people_data
        WHERE`
	_delEnd = `
RETURNING id, name, surname, patronymic, age, gender, country)
select *
from deleted
ORDER BY id
LIMIT @limit OFFSET @offset;`
	_updateInfoQuery = `
UPDATE people_data
SET name = @name, surname = @surname, patronymic = @patronymic, age = @age, gender = @gender, country = @country
WHERE id = @id`
)

type PagesDB struct {
	db *pgxpool.Pool
}

func NewPagesDB(db *pgxpool.Pool) *PagesDB {
	return &PagesDB{db: db}
}

func (p PagesDB) GetInfo(ctx context.Context, info *dto.Info, page *dto.Page) (*[]dto.Info, error) {
	log.Trace().Msg(fmt.Sprintf("DB recieve %+v\n", info))
	query := _getInfoQuery + dto.Info2String(info) + _getEnd
	log.Trace().Msg(fmt.Sprintf("query: ", query))

	args := dto.Info2map(info)
	args["limit"] = page.Limit
	args["offset"] = page.Offset

	rows, err := p.db.Query(context.TODO(), query, pgx.NamedArgs(args))
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprintf("GetInfo could not get %+v", info))
		return &[]dto.Info{}, dto.QueryExecuteErorr
	}

	people, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.Info])
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprintf("CollectRows error"))
		return &[]dto.Info{}, dto.QueryExecuteErorr
	}
	return &people, nil
}

func (p PagesDB) SetInfo(ctx context.Context, info *dto.Info) (*dto.Info, error) {
	var id int
	err := p.db.QueryRow(context.TODO(), _setInfoQuery, pgx.NamedArgs(dto.Info2map(info))).Scan(&id)
	info.Id = id
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprintf("SetInfo could not set %+v", info))
		return &dto.Info{}, dto.QueryExecuteErorr
	}
	return info, nil
}

func (p PagesDB) DeleteInfo(ctx context.Context, info *dto.Info, page *dto.Page) (*[]dto.Info, error) {
	where := dto.Info2String(info)
	if where == " true " {
		log.Warn().Msg("empty clause atempt")
		return &[]dto.Info{}, fmt.Errorf("clause cannot be empty")
	}

	args := dto.Info2map(info)
	args["limit"] = page.Limit
	args["offset"] = page.Offset

	rows, err := p.db.Query(context.TODO(),
		_delInfoQuery+where+_delEnd,
		pgx.NamedArgs(args))
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprintf("DeleteInfo could not delete %+v", info))
		return &[]dto.Info{}, dto.QueryExecuteErorr
	}

	people, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.Info])
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprint("CollectRows error"))
		return &[]dto.Info{}, dto.QueryExecuteErorr
	}
	return &people, nil
}

func (p PagesDB) UpdateInfo(ctx context.Context, info *dto.Info) (*dto.Info, error) {
	res, err := p.db.Exec(context.TODO(), _updateInfoQuery, pgx.NamedArgs(dto.Info2map(info)))
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprintf("UpdateInfo Could not update %+v", info))
		return &dto.Info{}, dto.QueryExecuteErorr
	}
	n := res.RowsAffected()
	if n == 0 {
		return info, fmt.Errorf("query found nothing to update")
	}

	return info, nil
}
