package cli

import (
	"fmt"
	"os"
	"strings"
)

const (
	NOTOK             string = ColorRed + "NOT OK" + ColorReset
	OK                string = ColorGreen + "OK" + ColorReset
	sprint1Secret1Str string = "HELLO LIGA"
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
