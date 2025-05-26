package moradorService

import (
	"backend/db"
	"backend/errs"
	"backend/utils/cpf"
	"context"
	"fmt"
)

func Validate(cpf cpf.CPF) *errs.AppError {
	// TODO: Remove query from service
	const query = `SELECT 1 FROM moradores WHERE cpf=$1`
	var exists int

	err := db.Connection.QueryRow(context.Background(), query, cpf).Scan(&exists)
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("querying morador from CPF: %w", err))
		return err
	}

	if exists == 0 {
		return errs.BadRequest("Não há morador com esse CPF", nil)
	}

	return nil
}
