package cli

import (
	"fmt"
)

type Step struct {
	ID    int         `json:"id"`
	Body  string      `json:"body"`
	Check func() bool `json:"-"`
}

type Sprint struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Steps []Step `json:"steps"`
	data  map[string]string
}

func True() bool {
	return true
}

func prettyParam(param string) string {
	return ColorCyan + param + ColorReset
}

func NewSprint(id int64, title string) *Sprint {
	return &Sprint{
		ID:    id,
		Title: title,
		data:  make(map[string]string),
	}
}

func (s *Sprint) AddStep(step Step) {
	s.Steps = append(s.Steps, step)
}

func (s *Sprint) GetSteps() []Step {
	return s.Steps
}

func printSprint(s *Sprint) {
	printAgent(
		fmt.Sprintf("Sprint %d. %s", s.ID, prettyTitle(s.Title)))
	for _, step := range s.GetSteps() {
		if step.Body == "" {
			continue
		}
		ok := NOTOK
		if step.Check() {
			ok = OK
		}
		printAgent(
			fmt.Sprintf("%d. %s... %s", step.ID, step.Body, ok))
	}
}

// Sprint first (1)
var SprintVPN = NewSprint(1, "Настрока VPN")

// Sprint second (2)
var SprintFS = NewSprint(2, "Базовое окружение и структура файловой системы Linux")

func SprintsInit() {
	SprintVPN.AddStep(Step{
		ID:    1,
		Body:  "Если вы видите это сообщение, значит вы все уже настроили",
		Check: True,
	})
	SprintVPN.AddStep(Step{
		ID:    2,
		Body:  fmt.Sprintf("Запустите команду '%s' для проверки этого спринта", prettyParam("cliga check --user [myname] sprint 1")),
		Check: True,
	})

	// 2
	SprintFS.AddStep(Step{
		ID:    1,
		Body:  fmt.Sprintf("Создан скрипт '%s' и находится в директории '%s'", prettyParam("sprint1.sh"), prettyParam("/opt/sprint1/")),
		Check: sprint1Step1,
	})
	SprintFS.AddStep(Step{
		ID: 2,
		Body: fmt.Sprintf(`Скрипт умеет создавать следующую структуру и файл: '%s' 
	- Проверка на наличие директорий и файлов`, prettyParam("/tmp/dir/subdir/file.txt")),
		Check: sprint1Step2,
	})

	idx, sprint1SecretPart := getSecretSprint1()
	SprintFS.AddStep(Step{
		ID: 3,
		Body: fmt.Sprintf(`[%s %d/%d] Вы получили фрагмент сообщения: '%s'
	- Найдены недостающие части у ваших коллег и выполнено условие сообщения`,
			prettyTitle("Group"), idx, len(sprint1Secret1Parts),
			prettyParam(sprint1SecretPart)),
		Check: sprint1Step3,
	})
}
