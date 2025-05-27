package moradorModel

import (
	"backend/db"
	"backend/errs"
	"backend/schemas"
	"backend/utils"
	"backend/utils/cpf"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func FindByCPF(cpf cpf.CPF) (*schemas.Morador, *errs.AppError) {
	var morador schemas.Morador

	query, err := utils.LoadSQL("morador/find_by_cpf.sql")
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading find morador by CPF SQL file: %w", err))
		return nil, err
	}

	// TODO: Use pgxscan to scan all fields at once
	err = db.Connection.QueryRow(context.Background(), query, cpf.Value).Scan(&morador.Cpf, &morador.Apartamento_id, &morador.Nome, &morador.Telefone, &morador.Responsavel, &morador.Proprietario)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("querying morador: %w", err))
		return nil, err
	}

	return &morador, nil
}
