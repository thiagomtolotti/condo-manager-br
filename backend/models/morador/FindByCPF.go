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

func FindByCPF(cpf cpf.CPF) (*schemas.Morador, *errs.AppError) {
	var morador schemas.Morador

	query, err := utils.LoadSQL("morador/find_by_cpf.sql")
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading find morador by CPF SQL file: %w", err))
		return nil, err
	}

	err = db.Connection.QueryRow(context.Background(), query, cpf).Scan(&morador)
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("querying morador: %w", err))
		return nil, err
	}

	return &morador, nil
}
