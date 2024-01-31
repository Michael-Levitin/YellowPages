package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Michael-Levitin/YellowPages/internal/dto"
	"github.com/Michael-Levitin/YellowPages/internal/logic"
	"net/http"
)

type PagesServer struct {
	logic logic.PagesLogicI
}

func NewPagesServer(logic logic.PagesLogicI) *PagesServer {
	return &PagesServer{logic: logic}
}

func (p PagesServer) SetInfo(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query() // Получаем все query параметры из URL запроса

	var info dto.Info
	info.Name = queryParams.Get("name")
	info.Surname = queryParams.Get("surname")
	info.Patronymic = queryParams.Get("patronymic")
	if info.Name == "" || info.Surname == "" {
		fmt.Fprintln(w, "both name and surname are required")
		return
	}

	age, err := logic.GetAge(info.Name)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	info.Age = age

	sex, err := logic.GetGender(info.Name)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	info.Sex = sex

	country, err := logic.GetNationality(info.Name)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	info.Country = country

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)

	p.logic.SetInfo(context.TODO(), info)
}
