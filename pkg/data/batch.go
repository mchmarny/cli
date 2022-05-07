package data

import (
	"time"

	"github.com/pkg/errors"
)

func SaveBatch(ids []int64) error {
	if db == nil {
		return errors.New("database not initialized")
	}

	stmt, err := db.Prepare("INSERT INTO sample (id, date) VALUES (?, ?")
	if err != nil {
		return errors.Wrapf(err, "failed to prepare batch statement")
	}

	tx, err := db.Begin()
	if err != nil {
		return errors.Wrapf(err, "failed to begin transaction")
	}

	for _, id := range ids {
		_, err = tx.Stmt(stmt).Exec(id, time.Now().UTC().Unix())
		if err != nil {
			if err = tx.Rollback(); err != nil {
				return errors.Wrapf(err, "failed to rollback transaction")
			}
			return errors.Wrapf(err, "failed to execute batch statement")
		}
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrapf(err, "failed to commit transaction")
	}

	return nil
}
