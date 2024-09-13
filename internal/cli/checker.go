package cli

import (
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

// Sprint 2. Step 1
func sprint1Step1() bool {
	return fileExists("/opt/sprint1/sprint1.sh")
}

func sprint1Step2() bool {
	return fileExists("/tmp/dir/subdir/file.txt")

}

func sprint1Step3() bool {
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
