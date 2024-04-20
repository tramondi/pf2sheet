package pg

import (
	"errors"
	"fmt"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/jackc/pgx/v5"
)

func wrapQueryError(err error, msgFormat string, args ...any) error {
	if err == nil {
		return nil
	}

	msg := fmt.Sprintf(msgFormat, args...)

	if errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("error %w: %s", contract.ErrNotFound, msg)
	}

	return fmt.Errorf("unexpected error %s: %s", err.Error(), msg)
}
