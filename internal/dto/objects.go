package dto

type Info struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Sex        string `json:"sex"`
	Country    string `json:"country"`
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
