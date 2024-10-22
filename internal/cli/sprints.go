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

func constTrue() bool {
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

// Sprints
var SprintVPN = NewSprint(1, "Настрока VPN")
var SprintFS = NewSprint(2, "Базовое окружение и структура файловой системы Linux")
var SprintGrep = NewSprint(3, "Обработка текста и SSH")
var SprintLVM = NewSprint(4, "Работа с LVM, Файловые системы, Python для DevOps")
var SprintDeamon = NewSprint(5, "Пакетные менеджеры и системы инициализации")
var SprintVLAN = NewSprint(6, "Сетевая подсистема Linux")
var SprintOps = NewSprint(7, "Базовая диагностика неисправностей и основы отказоустойчивых систем")

func SprintsInit() {
	SprintVPN.AddStep(Step{
		ID:    1,
		Body:  "Если вы видите это сообщение, значит вы все уже настроили",
		Check: constTrue,
	})
	SprintVPN.AddStep(Step{
		ID:    2,
		Body:  fmt.Sprintf("Запустите команду '%s' для проверки этого спринта", prettyParam("cliga check --user [myname] 1")),
		Check: constTrue,
	})

	// 2
	SprintFS.AddStep(Step{
		ID: 1,
		Body: fmt.Sprintf(`Созданы директории для хранения скриптов по спринтам 
	- Имя: %sX - вместо X число спринта. Начиная с 1 до 7. В директории '%s' 
	- Проверка на наличие директорий и файлов`,
			prettyParam("sprint"),
			prettyParam("/opt"),
		),
		Check: sprint1Step0,
	})
	SprintFS.AddStep(Step{
		ID:    2,
		Body:  fmt.Sprintf("Создан скрипт '%s' и находится в директории '%s'", prettyParam("sprint1.sh"), prettyParam("/opt/sprint1/")),
		Check: sprint1Step1,
	})
	SprintFS.AddStep(Step{
		ID: 3,
		Body: fmt.Sprintf(`Скрипт умеет создавать следующую структуру и файл: '%s' 
	- Проверка на наличие директорий и файлов`, prettyParam("/tmp/dir/subdir/file.txt")),
		Check: sprint1Step2,
	})

	SprintFS.AddStep(Step{
		ID: 4,
		Body: fmt.Sprintf(`Скрипт умеет создавать 99 директорий 
	- Имя: %sX - вместо X число. Начиная с 1. В директории '%s' 
	- Проверка на наличие директорий и файлов`,
			prettyParam("gendir"),
			prettyParam("/tmp"),
		),
		Check: sprint1Step3,
	})

	idx, sprint1SecretPart := getSecretSprint1()
	SprintFS.AddStep(Step{
		ID: 5,
		Body: fmt.Sprintf(`[%s %s/%d] Вы получили фрагмент сообщения: '%s'
	- Найдены недостающие части у ваших коллег и выполнено условие сообщения`,
			prettyTitle("Group"), idx, len(sprint1Secret1Parts),
			prettyParam(sprint1SecretPart)),
		Check: sprint1StepGroup,
	})

	// SPRINT 3
	SprintGrep.AddStep(Step{
		ID: 1,
		Body: fmt.Sprintf(`Скачан файл %s и лежит по пути %s
	- Используйте curl или wget (man curl)`,
			prettyParam("http://192.168.213.84/sprints/3/task1.txt"),
			prettyParam("/tmp/task1.txt"),
		),
		Check: sprint3Step1,
	})

	SprintGrep.AddStep(Step{
		ID: 2,
		Body: fmt.Sprintf(`Оригинальный файл клонирован по пути %s и изменено следующее
	- Все фразы %s заменены на %s
	- Используйте утилиту %s`,
			prettyParam("/tmp/task1_sed.txt"),
			prettyParam("The liga"),
			prettyParam("From sprint 3"),
			prettyParam("sed"),
		),
		Check: sprint3Step2,
	})

	SprintGrep.AddStep(Step{
		ID: 3,
		Body: fmt.Sprintf(`Оригинальный файл клонирован по пути %s и изменено следующее
	- Файл содержит только предложения с фразой %s
	- Файл отсортирован
	- Используйте утилиты %s и %s`,
			prettyParam("/tmp/task1_sort.txt"),
			prettyParam("The liga"),
			prettyParam("grep"),
			prettyParam("sort"),
		),
		Check: sprint3Step3,
	})

	SprintGrep.AddStep(Step{
		ID: 4,
		Body: fmt.Sprintf(`На все файлы %s примемены политики безопасности:
	- %s`,
			prettyParam("/tmp/task1*.txt"),
			prettyParam("rwxrwxrwx"),
		),
		Check: sprint3Step4,
	})

	SprintGrep.AddStep(Step{
		ID: 5,
		Body: fmt.Sprintf(`Создан пользователь %s:
	- Он также имеет свою домашнюю директорию %s`,
			prettyParam("visiter"),
			prettyParam("/home/visiter"),
		),
		Check: sprint3Step5,
	})

	SprintGrep.AddStep(Step{
		ID: 6,
		Body: fmt.Sprintf(`Групповое задание (*):
	- Новый пользовтаель %s имеет пароль %s
	- На вашу ВМ разрешен вход по паролю пользователю %s
	- Пользователь %s не имеет админских прав и не может их получить
	- Спишитесь с коллегами, узнайте у кого получилось сделать все шаги
	- Попробуйте подключиться к ним и оставить 
	дружественное послание в файлике %s`,
			prettyParam("visiter"),
			prettyParam("1234"),
			prettyParam("visiter"),
			prettyParam("visiter"),
			prettyParam("/home/visiter/visiter.txt"),
		),
		Check: constTrue,
	})

	// SPRINT 4
	SprintLVM.AddStep(Step{
		ID: 1,
		Body: fmt.Sprintf(`Установлен пакет %s для python3 (pip)`,
			prettyParam("requests"),
		),
		Check: sprint4Step1,
	})

	SprintLVM.AddStep(Step{
		ID: 2,
		Body: fmt.Sprintf(`Существует блочник %s`,
			prettyParam("lv_lesson"),
		),
		Check: sprint4Step2,
	})

	SprintLVM.AddStep(Step{
		ID: 3,
		Body: fmt.Sprintf(`Файловая система %s смонтирована`,
			prettyParam("/mnt/lesson4"),
		),
		Check: sprint4Step3,
	})

	SprintLVM.AddStep(Step{
		ID: 4,
		Body: fmt.Sprintf(`Сервер python-http запущен на порту %s`,
			prettyParam("8080"),
		),
		Check: sprint4Step4,
	})

	SprintLVM.AddStep(Step{
		ID: 5,
		Body: fmt.Sprintf(`Групповое задание (*)
	- Предоставьте права пользователю %s на %s
	- Зайдите под пользователем %s на любую чужую ВМ
	- Сохраните свой ключ ssh на чужой вм
	- Скачайте файлы из %s к себе на ВМ`,
			prettyParam("visiter"),
			prettyParam("/mnt/lesson4/"),
			prettyParam("visiter"),
			prettyParam("/mnt/lesson4/"),
		),
		Check: constTrue,
	})

	SprintDeamon.AddStep(Step{
		ID: 1,
		Body: fmt.Sprintf(`Сервис %s запущен`,
			prettyParam("zabbix-agent"),
		),
		Check: sprint5Step1,
	})

	SprintDeamon.AddStep(Step{
		ID: 2,
		Body: fmt.Sprintf(`Zabbix-Agent. Параметры настроены:
	- %s
	- %s`,
			prettyParam("Hostname"),
			prettyParam("Server"),
		),
		Check: sprint5Step2,
	})

	SprintDeamon.AddStep(Step{
		ID: 3,
		Body: fmt.Sprintf(`Сервис %s запущен и в статусе %s`,
			prettyParam("nginx"),
			prettyParam("enabled"),
		),
		Check: sprint5Step3,
	})

	SprintDeamon.AddStep(Step{
		ID: 4,
		Body: fmt.Sprintf(`Сервис %s настроен на: 
	- Рестарт после сбоя программы`,
			prettyParam("nginx"),
		),
		Check: sprint5Step4,
	})

	SprintDeamon.AddStep(Step{
		ID: 5,
		Body: fmt.Sprintf(`Групповое задание (*)
	- Договорись с коллегой или зайди на те ВМ которые уже заходил
	- Сделай так, чтобы сервис %s не мог запуститься)))`,
			prettyParam("nginx"),
		),
		Check: constTrue,
	})

	SprintVLAN.AddStep(Step{
		ID:    1,
		Body:  fmt.Sprintf("%s существует", prettyParam("VLAN 500")),
		Check: sprint6step1,
	})

	SprintVLAN.AddStep(Step{
		ID: 2,
		Body: fmt.Sprintf(`Групповое задание (*)
	- С помощью %s запрети заходить по SSH на %s
	- Удостоверься, что ВМ соседей пингуется по новому %s
	- Попробуй зайти на ВМ соседей по новому %s
	- Если зашел, удали %s и его конфигурационные файлы`,
			prettyParam("iptables"), prettyParam("VLAN 500"),
			prettyParam("VLAN 500"), prettyParam("VLAN 500"),
			prettyParam("VLAN 500")),
		Check: constTrue,
	})

	SprintOps.AddStep(Step{
		ID:    1,
		Body:  fmt.Sprintf("Существует файл %s", prettyParam("/var/log/messages-debug")),
		Check: sprint7step1,
	})

	SprintOps.AddStep(Step{
		ID:    2,
		Body:  fmt.Sprintf("Запущена служба %s", prettyParam("rsyslog")),
		Check: sprint7step2,
	})

	SprintOps.AddStep(Step{
		ID: 3,
		Body: fmt.Sprintf(`Существует тестовое сообщение %s
	- Проверка в файле: %s`, prettyParam("This is a debug message"), prettyParam("/var/log/messages-debug")),
		Check: sprint7step3,
	})

	SprintOps.AddStep(Step{
		ID:    4,
		Body:  fmt.Sprintf("Установлена и запущена служба %s", prettyParam("Keepalived")),
		Check: sprint7step4,
	})

	SprintOps.AddStep(Step{
		ID: 5,
		Body: fmt.Sprintf(`Файл конфигурации существует 
	- %s`, prettyParam("/etc/keepalived/keepalived.conf")),
		Check: sprint7step5,
	})

	SprintOps.AddStep(Step{
		ID: 6,
		Body: fmt.Sprintf(`Установлены и запущены сервисы: 
	- %s
	- %s`,
			prettyParam("haproxy"), prettyParam("nginx")),
		Check: sprint7step6,
	})

	SprintOps.AddStep(Step{
		ID: 7,
		Body: fmt.Sprintf(`Существуют сертификат и ключ
	- %s
	- %s`,
			prettyParam("/etc/ssl/nginx/nginx.crt"), prettyParam("/etc/ssl/nginx/nginx.key")),
		Check: sprint7step7,
	})

	SprintOps.AddStep(Step{
		ID:    8,
		Body:  "Выполнить групповое задание из презентации (*)",
		Check: constTrue,
	})

}
