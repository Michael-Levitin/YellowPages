package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Michael-Levitin/YellowPages/internal/dto"
	"github.com/Michael-Levitin/YellowPages/internal/logic"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
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
		log.Warn().Err(err).Msg("error checking full name")
		fmt.Fprintln(w, "error checking full name: ", err)
		return
	}

	info, err = p.logic.SetInfo(context.TODO(), info)
	if err != nil {
		log.Warn().Err(err).Msg("error executing p.logic.SetInfo")
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func (p PagesServer) GetInfo(w http.ResponseWriter, r *http.Request) {
	info, err := getParam(r)
	if err != nil {
		log.Warn().Err(err).Msg("error checking full name")
		fmt.Fprintln(w, "error checking full name: ", err)
		return
	}

	people, err := p.logic.GetInfo(context.TODO(), info)
	if err != nil {
		log.Warn().Err(err).Msg("error executing p.logic.GetInfo")
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	for _, info := range *people {
		json.NewEncoder(w).Encode(info)
	}
}

func (p PagesServer) DeleteInfo(w http.ResponseWriter, r *http.Request) {
	info, err := getParam(r)
	if err != nil {
		log.Warn().Err(err).Msg("error reading parameters")
		fmt.Fprintln(w, "error reading parameters")
		return
	}

	people, err := p.logic.DeleteInfo(context.TODO(), info)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "deleted %d rows, folowing entries were deleted:\n", len(*people))
	for _, info := range *people {
		json.NewEncoder(w).Encode(info)
	}
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
		return &dto.Info{}, fmt.Errorf("full name must be capitalized")
	}
	return &info, nil
}

func getParam(r *http.Request) (*dto.Info, error) {
	var info dto.Info
	var err error
	queryParams := r.URL.Query()
	idS := queryParams.Get("id")
	if idS != "" {
		info.Id, err = strconv.Atoi(idS)
		if err != nil {
			fmt.Println("delivery error: ", err)
			return &dto.Info{}, err
		}
	}
	info.Name = queryParams.Get("name")
	info.Surname = queryParams.Get("surname")
	info.Patronymic = queryParams.Get("patronymic")
	ageS := queryParams.Get("age")
	if ageS != "" {
		info.Age, err = strconv.Atoi(ageS)
		if err != nil {
			fmt.Println("delivery error: ", err)
			return &dto.Info{}, err
		}
	}
	info.Sex = queryParams.Get("sex")
	info.Country = queryParams.Get("country")
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
