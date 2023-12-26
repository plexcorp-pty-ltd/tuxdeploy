package tasks

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

func GetTmpPath(processType string) string {
	cwd, _ := os.Getwd()
	h := md5.New()
	h.Write([]byte(cwd))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	tmpPath := "/tmp/tuxdeploy_" + processType + "_" + hash + ".log"

	return tmpPath
}

func GetLastStepFromTmpFile(tmpPath string) (bool, int) {
	tmpPathExists := false
	lastStep := 0

	if _, err := os.Stat(tmpPath); err == nil {
		tmpPathExists = true
		data, err := os.ReadFile(tmpPath)
		if err == nil {
			res, err := strconv.Atoi(string(data))
			if err == nil {
				lastStep = res
			}
		}
	}

	return tmpPathExists, lastStep
}

func WriteStepToTmpFile(tmpPath string, tmpPathExists bool, stepNum int) {
	if tmpPathExists {
		os.Truncate(tmpPath, 0)
	}
	os.WriteFile(tmpPath, []byte(fmt.Sprintf("%d", stepNum)), 0755)
}
