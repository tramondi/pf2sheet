package pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
)

func wrapQueryError(err error, msgFormat string, args ...any) error {
	if err == nil {
		return nil
	}

	msg := fmt.Sprintf(msgFormat, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("error %w: %s", contract.ErrNotFound, msg)
	}

	return fmt.Errorf("unexpected error %s: %s", err.Error(), msg)
}

func sqlNullInt64(value *int64) sql.NullInt64 {
	sqlValue := sql.NullInt64{Valid: value != nil}
	if value != nil {
		sqlValue.Int64 = *value
	}

	return sqlValue
}

func sqlNullInt32(value *int32) sql.NullInt32 {
	sqlValue := sql.NullInt32{Valid: value != nil}
	if value != nil {
		sqlValue.Int32 = *value
	}

	return sqlValue
}

func sqlNullInt16(value *int16) sql.NullInt16 {
	sqlValue := sql.NullInt16{Valid: value != nil}
	if value != nil {
		sqlValue.Int16 = *value
	}

	return sqlValue
}

func sqlNullString(value *string) sql.NullString {
	sqlValue := sql.NullString{Valid: value != nil}
	if value != nil {
		sqlValue.String = *value
	}

	return sqlValue
}
