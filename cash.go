package iptec

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger/v4"
)

func NewCach() *cash {
	opt := badger.DefaultOptions(filepath.Join(os.TempDir(), "iptec"))
	db, err := badger.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	return &cash{
		db: db,
	}
}

type cash struct {
	db *badger.DB
}

func (c *cash) Close() {
	c.db.Close()
}

type CashMixinInterface interface {
	cashInitialization(string, *cash)
}

type CashMixin struct {
	prefix string
	cash   *cash
}

func (m *CashMixin) cashInitialization(prefix string, cash *cash) {
	m.prefix = prefix
	m.cash = cash
}

func (m *CashMixin) CashGet(key string) ([]byte, error) {
	var value []byte
	return value, m.cash.db.View(
		func(tx *badger.Txn) error {
			item, err := tx.Get(m.prepareKey(key))
			if err != nil {
				return fmt.Errorf("getting value: %w", err)
			}
			value, err = item.ValueCopy(nil)
			if err != nil {
				return fmt.Errorf("copying value: %w", err)
			}
			return nil
		})
}

func (m *CashMixin) CashSet(key string, data []byte) error {
	return m.cash.db.Update(
		func(txn *badger.Txn) error {
			return txn.Set(m.prepareKey(key), data)
		})
}

func (m *CashMixin) CashExist(key string) (bool, error) {
	var exists bool
	err := m.cash.db.View(
		func(tx *badger.Txn) error {
			if val, err := tx.Get(m.prepareKey(key)); err != nil {
				return err
			} else if val != nil {
				exists = true
			}
			return nil
		})
	if errors.Is(err, badger.ErrKeyNotFound) {
		err = nil
	}
	return exists, err
}

func (m *CashMixin) CashDelete(key string) error {
	return m.cash.db.Update(
		func(txn *badger.Txn) error {
			return txn.Delete(m.prepareKey(key))
		})
}

func (m *CashMixin) prepareKey(key string) []byte {
	return []byte(fmt.Sprintf("%s:%s", m.prefix, key))
}
