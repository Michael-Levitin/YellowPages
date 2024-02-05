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
	info, err := getParam(r)
	if err != nil {
		log.Warn().Err(err).Msg("error reading parameters")
		fmt.Fprintln(w, "error reading parameters")
		return
	}
	err = checkFIO(info)
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
		fmt.Fprintln(w, "error reading parameters: ", err)
		return
	}

	page, err := getPage(r)
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, "showing results from page", page.Page)
	}

	people, err := p.logic.GetInfo(context.TODO(), info, page)
	if err != nil {
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
		fmt.Fprintln(w, "error reading parameters: ", err)
		return
	}

	page, err := getPage(r)
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, "showing result from page", page.Page)
	}

	people, err := p.logic.DeleteInfo(context.TODO(), info, page)
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

func (p PagesServer) UpdateInfo(w http.ResponseWriter, r *http.Request) {
	info, err := getParam(r)
	if err != nil {
		log.Warn().Err(err).Msg("error reading parameters")
		fmt.Fprintln(w, "error reading parameters", err)
		return
	}
	if info.Id == 0 {
		fmt.Fprintln(w, "you must specify Id")
		return
	}
	err = checkFIO(info)
	if err != nil {
		log.Warn().Err(err).Msg("error checking full name")
		fmt.Fprintln(w, "error checking full name: ", err)
		return
	}

	info, err = p.logic.UpdateInfo(context.TODO(), info)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "update info:\n")
	json.NewEncoder(w).Encode(info)

}

func checkFIO(info *dto.Info) error {
	if info.Name == "" || info.Surname == "" {
		return fmt.Errorf("both name and surname are required")
	}

	if !isLetter(info.Name) || !isLetter(info.Surname) || !isLetter(info.Patronymic) {
		return fmt.Errorf("full name must contain letters only")
	}

	if !isCapital(info.Name) || !isCapital(info.Surname) || !isCapital(info.Patronymic) {
		return fmt.Errorf("full name must be capitalized")
	}
	return nil
}

func getParam(r *http.Request) (*dto.Info, error) {
	var info dto.Info
	var err error
	queryParams := r.URL.Query()
	idS := queryParams.Get("id")
	if idS != "" {
		info.Id, err = strconv.Atoi(idS)
		if err != nil {
			log.Info().Err(err).Msg("couldn't get ID")
			return &dto.Info{}, fmt.Errorf("couldn't get Id")
		}
	}
	info.Name = queryParams.Get("name")
	info.Surname = queryParams.Get("surname")
	info.Patronymic = queryParams.Get("patronymic")
	ageS := queryParams.Get("age")
	if ageS != "" {
		info.Age, err = strconv.Atoi(ageS)
		if err != nil {
			log.Info().Err(err).Msg("couldn't get Age")
			return &dto.Info{}, fmt.Errorf("couldn't get Age")
		}
	}
	info.Gender = queryParams.Get("gender")
	info.Country = queryParams.Get("country")
	return &info, nil
}

func getPage(r *http.Request) (*dto.Page, error) {
	var page dto.Page
	var err error
	queryParams := r.URL.Query()
	pageS := queryParams.Get("page")
	if pageS != "" {
		page.Page, err = strconv.Atoi(pageS)
		if err != nil {
			log.Info().Err(err).Msg("couldn't get Page")
		}
	}
	if page.Page == 0 {
		return &dto.Page{
			Page: 1,
		}, fmt.Errorf("couldn't get page, setting page = 1")
	}
	return &page, nil
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isCapital(s string) bool {
	var fl rune
	for _, r := range s {
		fl = r
		break
	}
	if !unicode.IsUpper(fl) {
		return false
	}
	return true
}
