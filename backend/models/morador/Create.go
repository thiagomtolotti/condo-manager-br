package moradorModel

import (
	"backend/db"
	"backend/errs"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"
)

func Create(data schemas.Morador) *errs.AppError {
	query, err := utils.LoadSQL("morador/create.sql")
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading create morador SQL file: %w", err))
		return err
	}

	_, err = db.Connection.Exec(
		context.Background(),
		query,
		data.Cpf,
		data.Apartamento_id,
		data.Nome,
		data.Telefone,
		data.Responsavel,
		data.Proprietario,
	)
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("creating morador query: %w", err))
		return err
	}

	return nil
}
