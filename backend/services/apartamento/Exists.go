package apartamentoService

import (
	"backend/db"
	"backend/errs"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func Exists(id uuid.UUID) *errs.AppError {
	const query = `SELECT 1 FROM apartamentos WHERE id=$1`

	var exists int
	err := db.Connection.QueryRow(context.Background(), query, id).Scan(&exists)

	if err != nil {
		return errs.Unexpected(fmt.Errorf("reading apartamento query: %w", err))
	}

	if exists == 0 {
		return errs.BadRequest("Não há apartamento com este id", nil)
	}

	return nil
}
