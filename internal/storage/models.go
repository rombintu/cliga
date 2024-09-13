package storage

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"username"`
	// Якорь, заменяет пароль
	// Будет брать что то из ОС
	// Только для обновления чего либо
	Anchor string `json:"anchor"`
}
