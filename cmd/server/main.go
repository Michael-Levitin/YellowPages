package main

import (
	"encoding/json"
	"fmt"
	"github.com/Michael-Levitin/YellowPages/internal/dto"
	"log"
	"net/http"
)

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

//type Info struct {
//	Name       string `json:"name"`
//	Surname    string `json:"surname"`
//	Patronymic string `json:"patronymic"`
//	Age        int    `json:"age"`
//	Sex        string `json:"sex"`
//	Country    string `json:"country"`
//}

func getAge(name string) (int, error) {
	resp, err := http.Get("https://api.agify.io/?name=" + name)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var answer Age
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

	var answer Gender
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

	var answer Origin
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

func updateInfo(w http.ResponseWriter, r *http.Request) {
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

func main() {
	http.HandleFunc("/updateInfo", updateInfo)
	fmt.Println("server is running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
