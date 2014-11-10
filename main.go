package acom

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
)

func sendUdp(addr, content string) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalf("xxx")
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("xxx")
	}
	defer conn.Close()

	_, err = conn.Write([]byte(content))
	if err != nil {
		log.Fatalf("xxx")
	}
}

func genLogPath(format string, t time.Time) string {
	m := t.Minute()
	q := (m / 15) + 1

	logf := fmt.Sprintf(format, t.Year(), t.Month(), t.Day(),
		t.Year(), t.Month(), t.Day(), t.Hour(), q)

	return logf
}

func Logger(logFileFormat, contentFormat string, v ...interface{}) {
	var buffer bytes.Buffer

	t := time.Now()
	logFH := genLogPath(logFileFormat, t)

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
