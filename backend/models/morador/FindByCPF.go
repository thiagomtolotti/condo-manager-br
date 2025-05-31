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

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

func FindByCPF(cpf cpf.CPF) (*schemas.Morador, *errs.AppError) {
	var morador schemas.Morador

	query, err := utils.LoadSQL("morador/find_by_cpf.sql")
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading find morador by CPF SQL file: %w", err))
		return nil, err
	}

	err = pgxscan.Get(context.Background(), db.Connection, &morador, query, cpf.Value)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("querying morador: %w", err))
		return nil, err
	}

	return &morador, nil
}
