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

	info, err := p.logic.SetInfo(context.TODO(), info)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
