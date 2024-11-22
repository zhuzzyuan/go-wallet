package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"go-wallet/util/color"
	"go-wallet/util/log"

	"github.com/lib/pq"
)

// IsDuplicateEntryError checks if duplicated record error occurred or not.
func IsDuplicateEntryError(err error) bool {
	if postgresErr, ok := err.(*pq.Error); ok {
		return postgresErr.Code.Name() == "23505"
	}

	return false
}

// IsRecordNotFoundError checks if no rows was returned or not.
func IsRecordNotFoundError(err error) bool {
	return err == sql.ErrNoRows
}

// CheckIfRowsNotAffected checks if sql execution affects no rows or not.
func CheckIfRowsNotAffected(result sql.Result, query []string) {
	affectedRows, err := result.RowsAffected()
	if err != nil || affectedRows == 0 {
		if err == nil {
			err = fmt.Errorf("\nSQL execution failed:\n%s", strings.Join(query, "\n"))
		}

		log.Panic(color.BRed(err.Error()))
	}
}
