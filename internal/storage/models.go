package storage

type Step struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Sprint struct {
	Title string `json:"title"`
	Steps []Step `json:"steps"`
}

type User struct {
	Name string `json:"username"`
	// Якорь, заменяет пароль
	// Будет брать что то из ОС
	// Только для обновления чего либо
	Anchor string `json:"anchor"`
}
