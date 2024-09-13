package storage

type Step struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

type Sprint struct {
	ID    int64  `json:"id"`
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

func NewSprint(id int64, title string) *Sprint {
	return &Sprint{
		ID:    id,
		Title: title,
	}
}

func (s *Sprint) AddStep(step Step) {
	s.Steps = append(s.Steps, step)
}

func (s *Sprint) GetSteps() []Step {
	return s.Steps
}
