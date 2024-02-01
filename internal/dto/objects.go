package dto

type Info struct {
	Id         int    `json:"id"`
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

func Info2map(info *Info) map[string]any {
	infoMap := make(map[string]any, 6)
	infoMap["name"] = info.Name
	infoMap["surname"] = info.Surname
	infoMap["patronymic"] = info.Patronymic
	infoMap["age"] = info.Age
	infoMap["sex"] = info.Sex
	infoMap["country"] = info.Country
	return infoMap
}
