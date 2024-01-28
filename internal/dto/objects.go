package dto

type Info struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Sex        string `json:"sex"`
	Country    string `json:"country"`
}
