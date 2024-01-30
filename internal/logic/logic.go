package logic

import (
	"context"
	"encoding/json"
	"fmt"
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

func (p PagesLogic) GetInfo(ctx context.Context, info dto.Info) (dto.Info, error) {
	return dto.Info{}, nil
}

func (p PagesLogic) SetInfo(ctx context.Context, info dto.Info) (dto.Info, error) {
	return dto.Info{}, nil
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

func UpdateInfo(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query() // Получаем все query параметры из URL запроса
	// Если необходимо получить отдельные параметры, можно использовать методы Get, GetArray, GetBool и другие.

	var info dto.Info
	info.Name = queryParams.Get("name")
	info.Surname = queryParams.Get("surname")
	info.Patronymic = queryParams.Get("patronymic")
	if info.Name == "" || info.Surname == "" {
		fmt.Fprintln(w, "both name and surname are required")
		return
	}

	age, err := getAge(info.Name)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	info.Age = age

	sex, err := getGender(info.Name)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	info.Sex = sex

	country, err := getNationality(info.Name)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	info.Country = country

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// http://localhost:8080/updateInfo?name=Andrej&surname=Sedov&patronymic=Aleksandorvich
