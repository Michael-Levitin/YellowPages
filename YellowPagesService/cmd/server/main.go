package main

import (
	"encoding/json"
	"fmt"
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

type Info struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Sex        string `json:"sex"`
	Country    string `json:"Country"`
}

func getAge(w http.ResponseWriter, r *http.Request) {
	// Получаем ФИО по API (это просто пример, здесь нужно заменить на реальный API)
	// Пример: https://example.com/api/name
	// Пример ответа: {"first_name": "John", "last_name": "Doe"}
	resp, err := http.Get("https://api.agify.io/?name=Dmitriy")
	if err != nil {
		http.Error(w, "Failed to get name", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var age Age
	if err := json.NewDecoder(resp.Body).Decode(&age); err != nil {
		http.Error(w, "Failed to decode name", http.StatusInternalServerError)
		return
	}

	// Отправляем полученное ФИО в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(age)
}

func getGender(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.genderize.io/?name=Dmitriy")
	if err != nil {
		http.Error(w, "Failed to get name", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var sex Gender
	if err := json.NewDecoder(resp.Body).Decode(&sex); err != nil {
		http.Error(w, "Failed to decode name", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sex)
}

func getNationality(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.nationalize.io/?name=Dmitriy")
	if err != nil {
		http.Error(w, "Failed to get name", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var country Origin
	if err := json.NewDecoder(resp.Body).Decode(&country); err != nil {
		http.Error(w, "Failed to decode name", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(country)
}

func updateInfo(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query() // Получаем все query параметры из URL запроса
	// Если необходимо получить отдельные параметры, можно использовать методы Get, GetArray, GetBool и другие.
	name := queryParams.Get("name")
	surname := queryParams.Get("surname")
	fmt.Fprintf(w, "param1: %s", name)
	fmt.Fprintf(w, "param2: %s", surname)

	resp, err := http.Get("https://api.nationalize.io/?name=Dmitriy")
	if err != nil {
		http.Error(w, "Failed to get name", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var country Info
	if err := json.NewDecoder(resp.Body).Decode(&country); err != nil {
		http.Error(w, "Failed to decode name", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(country)
}

func main() {
	http.HandleFunc("/getAge", getAge)
	http.HandleFunc("/getGender", getGender)
	http.HandleFunc("/getNationality", getNationality)
	http.HandleFunc("/updateInfo", updateInfo)
	fmt.Println("server is running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
