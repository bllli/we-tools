package persistence

import "github.com/jmoiron/sqlx"

type UnitOfWork struct {
	db *sqlx.DB
}

type UnitOfWorkFunc func(tx *sqlx.Tx) error

// NewUnitOfWork creates a new UnitOfWork
func NewUnitOfWork(db *sqlx.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

func (u *UnitOfWork) Execute(f UnitOfWorkFunc) error {
	tx, err := u.db.Beginx()
	if err != nil {
		return err
	}

	err = f(tx)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return rollBackErr
		}
		return err
	}
	return tx.Commit()
}
