package data

import (
	"fmt"

	"github.com/pkg/errors"
)

func (s *Service) SaveBatch(ids []int64) error {
	stmt, err := s.db.Prepare("update task set is_deleted='Y',last_modified_at=datetime() where id=?")
	if err != nil {
		return errors.Wrapf(err, "failed to prepare statement")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrapf(err, "failed to begin transaction")
	}

	for _, id := range ids {
		_, err = tx.Stmt(stmt).Exec(id)
	}

	if err != nil {
		fmt.Println("doing rollback")
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return nil
}
