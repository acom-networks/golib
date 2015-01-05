package acom

import (
	"bytes"
	"github.com/szferi/gomdb"
)

type MdbWrapper struct {
	env *mdb.Env
}

func NewMdb(filename string) (*MdbWrapper, error) {
	env, err := mdb.NewEnv()
	if err != nil {
		return nil, err
	}

	env.SetMapSize(100 * 1048576)

	err = env.Open(filename, mdb.NOSUBDIR, 0644)
	if err != nil {
		env.Close()
		return nil, err
	}

	return &MdbWrapper{env: env}, nil
}

func (mw *MdbWrapper) Put(bucket string, key, value []byte) error {
	txn, err := mw.env.BeginTxn(nil, 0)
	if err != nil {
		return err
	}

	dbi, err := txn.DBIOpen(nil, 0)
	if err != nil {
		txn.Abort()
		return err
	}

	var buf bytes.Buffer
	buf.WriteString(bucket)
	buf.WriteString("-")
	buf.Write(key)

	err = txn.Put(dbi, buf.Bytes(), value, 0)
	if err != nil {
		txn.Abort()
		return err
	}

	txn.Commit()
	return nil
}

func (mw *MdbWrapper) Get(bucket, defval string, key []byte) (string, error) {
	txn, err := mw.env.BeginTxn(nil, 0)
	if err != nil {
		return defval, err
	}
	defer txn.Abort()

	dbi, err := txn.DBIOpen(nil, 0)
	if err != nil {
		return defval, err
	}

	var buf bytes.Buffer
	buf.WriteString(bucket)
	buf.WriteString("-")
	buf.Write(key)

	value, err := txn.GetVal(dbi, buf.Bytes())
	if err != nil {
		return defval, err
	}

	return value.String(), nil
}

func (mw *MdbWrapper) Close() error {
	err := mw.env.Close()
	if err != nil {
		return err
	}

	mw.env = nil
	mw = nil

	return nil
}

func (mw *MdbWrapper) Stat() (*mdb.Stat, error) {
	return mw.env.Stat()
}

func (mw *MdbWrapper) Del(bucket string, key []byte) error {
	txn, err := mw.env.BeginTxn(nil, 0)
	if err != nil {
		return err
	}

	dbi, err := txn.DBIOpen(nil, 0)
	if err != nil {
		txn.Abort()
		return err
	}

	var buf bytes.Buffer
	buf.WriteString(bucket)
	buf.WriteString("-")
	buf.Write(key)

	err = txn.Del(dbi, buf.Bytes(), nil)
	if err != nil {
		txn.Abort()
		return err
	}

	txn.Commit()
	return nil
}
