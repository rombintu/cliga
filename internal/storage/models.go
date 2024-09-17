package storage

type SprintsOK []int64

type User struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	// Якорь, заменяет пароль
	// Будет брать что то из ОС
	// Только для обновления чего либо
	Anchor  string    `json:"anchor"`
	Sprints SprintsOK `json:"sprints_ok"`
}
