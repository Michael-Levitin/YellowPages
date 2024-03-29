package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Michael-Levitin/YellowPages/internal/database"
	"github.com/Michael-Levitin/YellowPages/internal/dto"
	"github.com/rs/zerolog/log"
	"net/http"
)

const RowsPerPage = 20

type PagesLogic struct {
	PagesDB database.PagesDbI
}

// NewPagesLogic подключаем интерфейс БД в новую логику
func NewPagesLogic(PagesDb database.PagesDbI) *PagesLogic {
	return &PagesLogic{PagesDB: PagesDb}
}

func (p PagesLogic) GetInfo(ctx context.Context, info *dto.Info, page *dto.Page) (*[]dto.Info, error) {
	log.Trace().Msg(fmt.Sprintf("Logic recieved %+v\n", info))
	page.Limit = RowsPerPage
	page.Offset = RowsPerPage * (page.Page - 1)

	people, err := p.PagesDB.GetInfo(ctx, info, page)
	if err != nil {
		return &[]dto.Info{}, err
	}
	if len(*people) == 0 {
		return &[]dto.Info{}, fmt.Errorf("query found nothing")
	}
	return people, nil
}

func (p PagesLogic) SetInfo(ctx context.Context, info *dto.Info) (*dto.Info, error) {
	log.Trace().Msg(fmt.Sprintf("Logic recieved %+v\n", info))
	info, err := getInfoApi(info)
	if err != nil {
		return &dto.Info{}, err
	}
	return p.PagesDB.SetInfo(ctx, info)
}

// http://localhost:8080/setInfo?name=Andrej&surname=Sedov&patronymic=Aleksandorvich

func (p PagesLogic) DeleteInfo(ctx context.Context, info *dto.Info, page *dto.Page) (*[]dto.Info, error) {
	log.Trace().Msg(fmt.Sprintf("Logic recieved %+v\n", info))
	page.Limit = RowsPerPage
	page.Offset = RowsPerPage * (page.Page - 1)

	people, err := p.PagesDB.DeleteInfo(ctx, info, page)
	if err != nil {
		log.Warn().Err(err).Msg("error executing p.logic.DeleteInfo")
		return &[]dto.Info{}, err
	}
	if len(*people) == 0 {
		return &[]dto.Info{}, fmt.Errorf("query found nothing - no rows were deleted")
	}
	return people, nil
}

func (p PagesLogic) UpdateInfo(ctx context.Context, info *dto.Info) (*dto.Info, error) {
	log.Trace().Msg(fmt.Sprintf("Logic recieved %+v\n", info))
	people, err := p.PagesDB.UpdateInfo(ctx, info)
	if err != nil {
		return &dto.Info{}, err
	}
	return people, nil
}

func getAge(name string) (int, error) {
	resp, err := http.Get("https://api.agify.io/?name=" + name)
	if err != nil {
		log.Info().Err(err).Msg("error getting age")
		return 0, fmt.Errorf("error getting age")
	}
	defer resp.Body.Close()

	var answer dto.Age
	if err = json.NewDecoder(resp.Body).Decode(&answer); err != nil {
		log.Info().Err(err).Msg("error decoding age")
		return 0, fmt.Errorf("error getting age")
	}

	return answer.Age, nil
}

func getGender(name string) (string, error) {
	resp, err := http.Get("https://api.genderize.io/?name=" + name)
	if err != nil {
		log.Info().Err(err).Msg("error getting gender")
		return "", fmt.Errorf("error getting gender")
	}
	defer resp.Body.Close()

	var answer dto.Gender
	if err = json.NewDecoder(resp.Body).Decode(&answer); err != nil {
		log.Info().Err(err).Msg("error decoding gender")
		return "", fmt.Errorf("error getting gender")
	}

	return answer.Gender, nil
}

func getNationality(name string) (string, error) {
	resp, err := http.Get("https://api.nationalize.io/?name=" + name)
	if err != nil {
		log.Info().Err(err).Msg("error getting nationality")
		return "", fmt.Errorf("error getting nationality")
	}
	defer resp.Body.Close()

	var answer dto.Origin
	if err = json.NewDecoder(resp.Body).Decode(&answer); err != nil {
		log.Info().Err(err).Msg("error decoding nationality")
		return "", fmt.Errorf("error getting nationality")
	}

	var country string
	maxP := 0.0
	for _, s := range answer.Country {
		if s.Probability > maxP {
			maxP = s.Probability
			country = s.CountryID
		}
		if s.Probability > 0.5 {
			break
		}
	}
	return country, nil
}

func getInfoApi(info *dto.Info) (*dto.Info, error) {
	age, err := getAge(info.Name)
	if err != nil {
		return &dto.Info{}, err
	}
	info.Age = age

	gender, err := getGender(info.Name)
	if err != nil {
		return &dto.Info{}, err
	}
	info.Gender = gender

	country, err := getNationality(info.Name)
	if err != nil {
		return &dto.Info{}, err
	}
	info.Country = country
	return info, nil
}
