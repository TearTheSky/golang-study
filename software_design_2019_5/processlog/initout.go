package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func initOut(logDir, commandName string) (stdout, stderr io.Writer, err error) {
	if logDir == "" {
		stdout = os.Stdout
		stderr = os.Stderr
	} else {
		ts := time.Now().Unix()

		stdoutFileName := fmt.Sprintf("%s-%v-stdlout.log", commandName, ts)
		stdoutFile, err := os.Create(filepath.Join(logDir, stdoutFileName))
		if err != nil {
			return nil, nil, err
		}
		stdout = io.MultiWriter(os.Stdout, stdoutFile)

		stderrFilename := fmt.Sprintf("%s-%v-stdlerr.log", commandName, ts)
		stderrFile, err := os.Create(filepath.Join(logDir, stderrFilename))
		if err != nil {
			return nil, nil, err
		}
		stderr = io.MultiWriter(os.Stderr, stderrFile)
	}
	return
}
