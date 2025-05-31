package vagaModel

import (
	"backend/db"
	"backend/errs"
	"backend/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func Delete(id uuid.UUID) *errs.AppError {
	query, err := utils.LoadSQL("vaga/delete.sql")
	if err != nil {
		return errs.Unexpected(fmt.Errorf("reading delete vaga SQL file: %w", err))
	}

	result, err := db.Connection.Exec(context.Background(), query, id.String())
	if err != nil {
		return errs.Unexpected(fmt.Errorf("deleting vaga in DB %w", err))
	}
	if result.RowsAffected() == 0 {
		return errs.BadRequest("nenhuma vaga com o id encontrada", nil)
	}

	return nil
}
