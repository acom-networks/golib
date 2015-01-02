package acom

import (
	"os"
	"testing"
)

var MdbTests = []struct {
	bucket string
	key    string
	value  string
}{
	{"MAC", "test", "123"},
	{"AC", "192.168.0.1", "aaa.host.name"},
	{"SMP", "192.168.0.2", "smp.hostname"},
}

func TestMdbWrapperFail(t *testing.T) {
	_, err := NewMdb("/tmp/xxxx/123")
	if err == nil {
		t.Fatalf("Should be failed, but open successed\n")
	}

	path := "./test"
	bdb, _ := NewMdb(path)
	defer bdb.Close()
	defer os.Remove(path)
	defer os.Remove(path + "-lock")

	bdb.Put("MAC", []byte("test"), []byte("123"))
	x, _ := bdb.Get("MAC", "", []byte("test1"))
	if x != "" {
		t.Fatalf("expected: \"\", but got %s\n", x)
	}

	x, _ = bdb.Get("MAC", "test1", []byte("test1"))
	if x != "test1" {
		t.Fatalf("expected: \"test1\", but got %s\n", x)
	}
}

func TestBdbWrapper(t *testing.T) {
	path := "./test"

	bdb, err := NewMdb(path)
	if err != nil {
		t.Fatalf("Fail to open mdb: %s\n", err)
	}

	defer bdb.Close()
	defer os.Remove(path)
	defer os.Remove(path + "-lock")

	for i, tt := range MdbTests {
		x, _ := bdb.Get(tt.bucket, "", []byte(tt.key))
		if x != "" {
			t.Fatalf("ERROR: test %d, expected: \"\", got: \"%s\"\n", i, x)
		}
		bdb.Put(tt.bucket, []byte(tt.key), []byte(tt.value))

		y, _ := bdb.Get(tt.bucket, "", []byte(tt.key))
		if y != tt.value {
			t.Fatalf("ERROR: test %d, expected: \"%s\", got: \"%s\"\n", i, tt.value, y)
		}
	}
}
