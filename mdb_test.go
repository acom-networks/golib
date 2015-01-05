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
	mdb, _ := NewMdb(path)
	defer mdb.Close()
	defer os.Remove(path)
	defer os.Remove(path + "-lock")

	mdb.Put("MAC", []byte("test"), []byte("123"))
	x, _ := mdb.Get("MAC", "", []byte("test1"))
	if x != "" {
		t.Fatalf("expected: \"\", but got %s\n", x)
	}

	x, _ = mdb.Get("MAC", "test1", []byte("test1"))
	if x != "test1" {
		t.Fatalf("expected: \"test1\", but got %s\n", x)
	}
}

func TestMdbWrapper(t *testing.T) {
	path := "./test"

	mdb, err := NewMdb(path)
	if err != nil {
		t.Fatalf("Fail to open mdb: %s\n", err)
	}

	defer mdb.Close()
	defer os.Remove(path)
	defer os.Remove(path + "-lock")

	for i, tt := range MdbTests {
		x, _ := mdb.Get(tt.bucket, "", []byte(tt.key))
		if x != "" {
			t.Fatalf("ERROR: test %d, expected: \"\", got: \"%s\"\n", i, x)
		}
		mdb.Put(tt.bucket, []byte(tt.key), []byte(tt.value))

		y, _ := mdb.Get(tt.bucket, "", []byte(tt.key))
		if y != tt.value {
			t.Fatalf("ERROR: test %d, expected: \"%s\", got: \"%s\"\n", i, tt.value, y)
		}
	}
}

func TestMdbWrapperStat(t *testing.T) {
	path := "./test"

	mdb, err := NewMdb(path)
	if err != nil {
		t.Fatalf("Fail to open mdb: %s\n", err)
	}

	defer mdb.Close()
	defer os.Remove(path)
	defer os.Remove(path + "-lock")

	mdb.Put("test", []byte("test"), []byte("test"))

	stat, err := mdb.Stat()
	if err != nil {
		t.Fatalf("ERROR: expected null error, got: %s\n", err)
	}

	if stat.Entries != 1 {
		t.Fatalf("ERROR: expected 1 entry, got: %d\n", stat.Entries)
	}

	mdb.Put("test", []byte("test1"), []byte("test1"))

	stat, err = mdb.Stat()
	if err != nil {
		t.Fatalf("ERROR: expected null error, got: %s\n", err)
	}

	if stat.Entries != 2 {
		t.Fatalf("ERROR: expected 2 entry, got: %d\n", stat.Entries)
	}
}
