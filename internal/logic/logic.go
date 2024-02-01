package logic

import (
	"context"
	"encoding/json"
	"github.com/Michael-Levitin/YellowPages/internal/database"
	"github.com/Michael-Levitin/YellowPages/internal/dto"
	"net/http"
)

type PagesLogic struct {
	PagesDB database.PagesDbI
}

// NewPagesLogic подключаем интерфейс БД в новую логику
func NewPagesLogic(PagesDb database.PagesDbI) *PagesLogic {
	return &PagesLogic{PagesDB: PagesDb}
}

func (p PagesLogic) GetInfo(ctx context.Context, info *dto.Info) (*dto.Info, error) {
	p.PagesDB.GetInfo(ctx, info)
	return &dto.Info{}, nil
}

func (p PagesLogic) SetInfo(ctx context.Context, info *dto.Info) (*dto.Info, error) {
	info, err := getInfoApi(info)
	if err != nil {
		return &dto.Info{}, err
	}
	return p.PagesDB.SetInfo(ctx, info)
}

func getAge(name string) (int, error) {
	resp, err := http.Get("https://api.agify.io/?name=" + name)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var answer dto.Age
	if err = json.NewDecoder(resp.Body).Decode(&answer); err != nil {
		return 0, err
	}

	return answer.Age, nil
}

func getGender(name string) (string, error) {
	resp, err := http.Get("https://api.genderize.io/?name=" + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var answer dto.Gender
	if err = json.NewDecoder(resp.Body).Decode(&answer); err != nil {
		return "", err
	}

	return answer.Gender, nil
}

func getNationality(name string) (string, error) {
	resp, err := http.Get("https://api.nationalize.io/?name=" + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var answer dto.Origin
	if err = json.NewDecoder(resp.Body).Decode(&answer); err != nil {
		return "", err
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

// http://localhost:8080/setInfo?name=Andrej&surname=Sedov&patronymic=Aleksandorvich

func getInfoApi(info *dto.Info) (*dto.Info, error) {
	age, err := getAge(info.Name)
	if err != nil {
		return &dto.Info{}, err
	}
	info.Age = age

	sex, err := getGender(info.Name)
	if err != nil {
		return &dto.Info{}, err
	}
	info.Sex = sex

	country, err := getNationality(info.Name)
	if err != nil {
		return &dto.Info{}, err
	}
	info.Country = country
	return info, nil
}
