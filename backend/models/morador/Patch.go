package moradorModel

import (
	"backend/db"
	"backend/errs"
	"backend/schemas"
	"backend/utils"
	"backend/utils/cpf"
	"context"
	"fmt"
)

func Patch(cpf cpf.CPF, data schemas.MoradorWithoutCPF) *errs.AppError {
	query, err := utils.LoadSQL("morador/patch.sql")
	if err != nil {
		return errs.Unexpected(fmt.Errorf("reading update morador SQL file: %w", err))
	}

	_, err = db.Connection.Exec(
		context.Background(),
		query,
		data.Apartamento_id,
		data.Nome,
		data.Telefone,
		data.Responsavel,
		data.Proprietario,
		cpf.Value,
	)
	if err != nil {
		return errs.Unexpected(fmt.Errorf("updating morador: %w", err))
	}

	return nil
}
