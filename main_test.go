package acom

import (
	"testing"
	"time"
)

func TestGenLogPath(t *testing.T) {
	var ts int64
	template := "%04d-%02d-%02d/lcp_sniffer-%04d-%02d-%02d-%02d-%d.log"

	ts = 1234567890
	ti := time.Unix(ts, 0)
	logf := genLogPath(template, ti)
	if logf != "2009-02-14/lcp_sniffer-2009-02-14-07-3.log" {
		t.Error("wrong output")
	}

	ts = 1235567890
	ti = time.Unix(ts, 0)
	logf = genLogPath(template, ti)
	if logf != "2009-02-25/lcp_sniffer-2009-02-25-21-2.log" {
		t.Error("wrong output")
	}
}
