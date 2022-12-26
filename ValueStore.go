package vaultstore

import (
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/gouniverse/uid"
)

// ValueStore creates a new vault entry and returns the ID
func (st *Store) ValueStore(value string, password string) (id string, err error) {
	encoded := encode(value, password)
	var newEntry = Vault{
		ID:        uid.HumanUid(),
		Value:     encoded,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var sqlStr string
	sqlStr, _, _ = goqu.Dialect(st.dbDriverName).Insert(st.vaultTableName).Rows(newEntry).ToSQL()

	if st.debugEnabled {
		log.Println(sqlStr)
	}

	_, err = st.db.Exec(sqlStr)
	if err != nil {
		if st.debugEnabled {
			log.Println(err.Error())
		}
		return "", err
	}
	return newEntry.ID, nil
}
