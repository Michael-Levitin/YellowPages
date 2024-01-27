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

type Agify struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
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

	var age Agify
	if err := json.NewDecoder(resp.Body).Decode(&age); err != nil {
		http.Error(w, "Failed to decode name", http.StatusInternalServerError)
		return
	}

	// Отправляем полученное ФИО в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(age)
}

func main() {
	http.HandleFunc("/getage", getAge)
	fmt.Println("server is running")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
