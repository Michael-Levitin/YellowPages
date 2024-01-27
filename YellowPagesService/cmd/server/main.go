package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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

func main() {
	http.HandleFunc("/getAge", getAge)
	http.HandleFunc("/getGender", getGender)
	http.HandleFunc("/getNationality", getNationality)
	fmt.Println("server is running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
