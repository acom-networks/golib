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

func SendUdp(logFileFormat, addr string, payload []byte) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		Logger(logFileFormat, "Fail to resolv udp address %s (%s)\n", addr, err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		Logger(logFileFormat, "Fail to connect to udp address %s (%s)\n", addr, err)
	}
	defer conn.Close()

	_, err = conn.Write(payload)
	if err != nil {
		Logger(logFileFormat, "Fail to write to udp address %s (%s)\n", addr, err)
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
