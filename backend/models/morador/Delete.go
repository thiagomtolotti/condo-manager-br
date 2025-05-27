package moradorModel

import (
	"backend/db"
	"backend/errs"
	"backend/utils"
	"backend/utils/cpf"
	"context"
	"fmt"
)

func Delete(cpf cpf.CPF) *errs.AppError {
	query, err := utils.LoadSQL("morador/delete.sql")
	if err != nil {
		return errs.Unexpected(fmt.Errorf("reading delete morador SQL file: %w", err))
	}

	result, err := db.Connection.Exec(context.Background(), query, cpf.Value)
	if err != nil {
		return errs.Unexpected(fmt.Errorf("deleting morador from DB: %w", err))
	}

	if result.RowsAffected() == 0 {
		return errs.BadRequest("Não há moradores com o CPF informado", nil)
	}

	return nil
}
