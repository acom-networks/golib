package acom

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Logger(genLogPath func(t time.Time) string, contentFormat string, v ...interface{}) {
	var buffer bytes.Buffer

	t := time.Now()
	logFH := genLogPath(t)

	dirname := filepath.Dir(logFH)
	if _, err := os.Stat(dirname); err != nil {
		os.MkdirAll(dirname, 0644)
	}

	f, err := os.OpenFile(logFH, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("%s", err)
	}

	buffer.WriteString("[")
	buffer.WriteString(t.String())
	buffer.WriteString("] ")
	buffer.WriteString(contentFormat)

	if len(v) > 0 {
		fmt.Fprintf(f, buffer.String(), v...)
	} else {
		fmt.Fprintf(f, buffer.String())
	}
	f.Close()
}
