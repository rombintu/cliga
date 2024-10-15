package cli

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"net"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

const (
	NOTOK             string = ColorRed + "NOT OK" + ColorReset
	OK                string = ColorGreen + "OK" + ColorReset
	sprint1Secret1Str string = "HELLO LIGA"
	OrigHASHfileTask1 string = "2d0d5ae879a784fd97e836867cfa0614"
	HASHfileTask2     string = "1cbcf0d448fb645cabd3fcfffb6507b8"
	HASHfileTask3     string = "1e8b315686070dfa270c7d9ca82404e8"
	HASHfileTask3v2   string = "1617de7bc198f162e0b31fe41a8c9e74"
)

var sprint1Secret1Parts = []string{
	"0KHQvtC30LTQsNC50YLQtSDRhNCw0LnQuw==",
	"L3RtcC9saWdhLnR4dA==",
	"0YEg0YHQvtC00LXRgNC20LjQvNGL0Lw=",
	"SEVMTE8gTElHQQ==",
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func dirNotExists(path string) bool {
	f, err := os.Stat(path)
	return os.IsNotExist(err) || !f.IsDir()
}

func hashFileIs(path string, hashFile string) bool {
	file, err := os.Open(path)
	if err != nil {
		// printAgentWarn("Ошибка при открытии файла", err, false)
		return false
	}
	defer file.Close()

	// Создаем новый хешер SHA-256
	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		// printAgentWarn("Ошибка при чтении файла", err, false)
		return false
	}
	// Получаем хеш в виде байтового массива
	hashInBytes := hash.Sum(nil)

	// Преобразуем хеш в строку в шестнадцатеричном формате
	hashString := hex.EncodeToString(hashInBytes)

	return hashString == hashFile
}

func filePermission(path string, perm fs.FileMode) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		// printAgentError("Ошибка при получении информации о файле", errNone, true)
		return false
	}
	mode := fileInfo.Mode()
	return mode.Perm() == perm
}

func userExists(username string) bool {
	u, err := user.Lookup(username)
	if err != nil {
		return false
	}
	homeDir := u.HomeDir
	return homeDir != ""
}

func ExecAndFindIsNotEmpty(c string, args []string, exists string) bool {
	cmd := exec.Command(c, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return strings.Contains(strings.TrimSpace(string(output)), exists)
}

func ExecAndFind(c string, args []string, exists string) bool {
	cmd := exec.Command(c, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 1 {
				printAgentWarn(
					fmt.Sprintf(
						"Команда %s %s выполнилась с ошибкой",
						c, strings.Join(args, " "),
					), err, false)
				return false
			}
		}
		printAgentWarn("Неизвестная ошибка", err, false)
		return false
	}
	return strings.Contains(string(output), exists)
}

func checkPortIsAvail(port string) bool {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return true
	}
	l.Close()
	return false
}

// Sprint 2
func sprint1Step0() bool {
	for i := 1; i < 8; i++ {
		if dirNotExists(fmt.Sprintf("/opt/sprint%d", i)) {
			return false
		}
	}
	return true

}

func sprint1Step1() bool {
	return fileExists("/opt/sprint1/sprint1.sh")
}

func sprint1Step2() bool {
	return fileExists("/tmp/dir/subdir/file.txt")

}

func sprint1Step3() bool {
	for i := 1; i < 100; i++ {
		if dirNotExists(fmt.Sprintf("/tmp/gendir%d", i)) {
			return false
		}
	}
	return true

}

func sprint1StepGroup() bool {
	filePath := "/tmp/liga.txt"
	if fileExists(filePath) {
		b, err := os.ReadFile(filePath)
		if err != nil {
			printAgentError("file error", err, true)
		}
		if strings.TrimSpace(string(b)) == sprint1Secret1Str {
			return true
		}
	}
	return false
}

func sprint3Step1() bool {
	return fileExists("/tmp/task1.txt") && hashFileIs("/tmp/task1.txt", OrigHASHfileTask1)
}

func sprint3Step2() bool {
	return fileExists("/tmp/task1_sed.txt") && hashFileIs("/tmp/task1_sed.txt", HASHfileTask2)
}

func sprint3Step3() bool {
	return fileExists("/tmp/task1_sort.txt") && hashFileIs("/tmp/task1_sort.txt", HASHfileTask3) || hashFileIs("/tmp/task1_sort.txt", HASHfileTask3v2)
}

func sprint3Step4() bool {
	return fileExists("/tmp/task1.txt") &&
		fileExists("/tmp/task1_sed.txt") &&
		fileExists("/tmp/task1_sort.txt") &&
		filePermission("/tmp/task1.txt", 0777) &&
		filePermission("/tmp/task1_sed.txt", 0777) &&
		filePermission("/tmp/task1_sort.txt", 0777)
}

func sprint3Step5() bool {
	return userExists("visiter")
}

func sprint4Step1() bool {
	return ExecAndFindIsNotEmpty("python3", []string{"-m", "pip", "show", "requests"}, "requests")
}

func sprint4Step2() bool {
	return ExecAndFindIsNotEmpty("lsblk", nil, "lv_lesson")
}

func sprint4Step3() bool {
	return ExecAndFindIsNotEmpty("findmnt", []string{"/mnt/lesson4"}, "/mnt/lesson4")
}

func sprint4Step4() bool {
	return checkPortIsAvail("8080")
}

func sprint5Step1() bool {
	return ExecAndFindIsNotEmpty("systemctl", []string{"is-active", "zabbix-agent"}, "active")
}

func sprint5Step2() bool {
	return ExecAndFindIsNotEmpty("cat", []string{"/etc/zabbix_agentd.conf"}, "Hostname") &&
		ExecAndFindIsNotEmpty("cat", []string{"/etc/zabbix_agentd.conf"}, "Server")
}

func sprint5Step3() bool {
	return ExecAndFindIsNotEmpty("systemctl", []string{"is-active", "nginx"}, "active") &&
		ExecAndFindIsNotEmpty("systemctl", []string{"is-enabled", "nginx"}, "enabled")
}

func sprint5Step4() bool {
	return ExecAndFindIsNotEmpty("systemctl", []string{"show", "nginx"}, "Restart=on-failure")
}
