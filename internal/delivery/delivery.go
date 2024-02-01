package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Michael-Levitin/YellowPages/internal/dto"
	"github.com/Michael-Levitin/YellowPages/internal/logic"
	"net/http"
	"unicode"
)

type PagesServer struct {
	logic logic.PagesLogicI
}

func NewPagesServer(logic logic.PagesLogicI) *PagesServer {
	return &PagesServer{logic: logic}
}

func (p PagesServer) SetInfo(w http.ResponseWriter, r *http.Request) {
	info, err := checkFIO(r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	info, err = p.logic.SetInfo(context.TODO(), info)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func checkFIO(r *http.Request) (*dto.Info, error) {
	var info dto.Info
	queryParams := r.URL.Query() // Получаем все query параметры из URL запроса
	info.Name = queryParams.Get("name")
	info.Surname = queryParams.Get("surname")
	info.Patronymic = queryParams.Get("patronymic")
	if info.Name == "" || info.Surname == "" {
		return &dto.Info{}, fmt.Errorf("both name and surname are required")
	}

	if !isLetter(info.Name) || !isLetter(info.Surname) || !isLetter(info.Patronymic) {
		return &dto.Info{}, fmt.Errorf("full name must contain letters only")
	}

	if !isCapital(info.Name) || !isCapital(info.Surname) || !isCapital(info.Patronymic) {
		return &dto.Info{}, fmt.Errorf("full name must be capitalizedddd")
	}
	return &info, nil
}

func isLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func isCapital(s string) bool {
	var fl rune
	for _, r := range s {
		fl = r
		break
	}
	if unicode.IsUpper(fl) {
		return true
	}
	return false
}
