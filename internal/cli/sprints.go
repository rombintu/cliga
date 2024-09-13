package cli

import (
	"fmt"

	"github.com/rombintu/checker-sprints/internal/storage"
)

// Sprint first (1)
var sprintVPN = storage.NewSprint(1, "Настрока VPN")

// Sprint second (2)
var sprintFS = storage.NewSprint(1, "Базовое окружение и структура файловой системы Linux")

func SprintsInit() {
	sprintVPN.AddStep(storage.Step{
		ID:   1,
		Body: "Если вы видите это сообщение, значит вы все уже настроили",
	})
	sprintVPN.AddStep(storage.Step{
		ID:   2,
		Body: fmt.Sprintf("Запустите команду '%s' для проверки этого спринта", ColorCyan+"cliga check --user [myname] sprint 1"+ColorReset),
	})

	// 2
	sprintFS.AddStep(storage.Step{
		ID:   1,
		Body: "",
	})
}
