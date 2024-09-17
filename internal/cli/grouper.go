package cli

const (
	GROUPID int = 1
)

// Пусть будет регистрация и возврат айди, по нему будет решаться группа
func getSecretSprint1() (int, string) {
	var secretpart string
	switch GROUPID {
	case 1:
		secretpart = sprint1Secret1Parts[0]
	case 2:
		secretpart = sprint1Secret1Parts[1]
	case 3:
		secretpart = sprint1Secret1Parts[2]
	case 4:
		secretpart = sprint1Secret1Parts[3]
	}
	return GROUPID, secretpart
}
