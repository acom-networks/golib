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

func sendUdp(content string) {
	udpAddr, err := net.ResolveUDPAddr("udp", cfg.Lcp.Collector)
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

func Logger(format string, v ...interface{}) {
	var buffer bytes.Buffer

	t := time.Now()
	logf := genLogPath(*logFormat, t)

	dirname := filepath.Dir(logf)
	if _, err := os.Stat(dirname); err != nil {
		os.MkdirAll(dirname, 0644)
	}

	f, err := os.OpenFile(logf, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("%s", err)
	}

	buffer.WriteString("[")
	buffer.WriteString(t.String())
	buffer.WriteString("] ")
	buffer.WriteString(format)

	if len(v) > 0 {
		fmt.Fprintf(f, buffer.String(), v...)
	} else {
		fmt.Fprintf(f, buffer.String())
	}
	f.Close()
}
