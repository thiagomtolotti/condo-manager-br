package apartmentoModel

import (
	"backend/db"
	"backend/errs"
	"backend/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func Delete(id uuid.UUID) *errs.AppError {
	query, err := utils.LoadSQL("apartamento/delete.sql")
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading delete apartamento SQL file: %w", err))
		return err
	}

	result, err := db.Connection.Exec(context.Background(), query, id)
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("deleting apartamento in DB: %w", err))
		return err
	}

	if result.RowsAffected() == 0 {
		var err = errs.BadRequest("Nenhum apartamento com o id foi encontrado", nil)
		return err
	}

	return nil
}
