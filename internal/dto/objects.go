package dto

import (
	"fmt"
	"strings"
)

type Info struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
	Age        int    `json:"age" db:"age"`
	Sex        string `json:"sex" db:"sex"`
	Country    string `json:"country" db:"country"`
}

type Age struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type Gender struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	CountryID   string  `json:"country_id,omitempty"`
	Probability float64 `json:"probability"`
}

type Origin struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

func Info2map(info *Info) map[string]any {
	infoMap := make(map[string]any, 7)
	infoMap["id"] = info.Id
	infoMap["name"] = info.Name
	infoMap["surname"] = info.Surname
	infoMap["patronymic"] = info.Patronymic
	infoMap["age"] = info.Age
	infoMap["sex"] = info.Sex
	infoMap["country"] = info.Country
	return infoMap
}

func Info2String(info *Info) string {
	fmt.Printf("Objects recieve %+v\n", info)
	infoString := make([]string, 0, 8)
	var s string
	infoString = append(infoString, " true ")
	if info.Id != 0 {
		s = " id = @id "
		infoString = append(infoString, s)
	}
	if info.Name != "" {
		info.Name = "%" + info.Name + "%"
		s = " name ILIKE @name "
		infoString = append(infoString, s)
	}
	if info.Surname != "" {
		info.Surname = "%" + info.Surname + "%"
		s = " surname ILIKE @surname "
		infoString = append(infoString, s)
	}
	if info.Patronymic != "" {
		info.Patronymic = "%" + info.Patronymic + "%"
		s = " patronymic ILIKE @patronymic "
		infoString = append(infoString, s)
	}
	if info.Age != 0 {
		s = " age = @age "
		infoString = append(infoString, s)
	}
	if info.Sex != "" {
		info.Sex = "%" + info.Sex + "%"
		s = " sex ILIKE @sex "
		infoString = append(infoString, s)
	}
	if info.Country != "" {
		info.Country = "%" + info.Country + "%"
		s = " country ILIKE @country "
		infoString = append(infoString, s)
	}

	return strings.Join(infoString, "AND") + " Order by id;"
}
