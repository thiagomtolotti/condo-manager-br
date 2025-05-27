package apartmentoModel

import (
	"backend/db"
	"backend/errs"
	"backend/schemas"
	"backend/utils"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func FindById(id uuid.UUID) (*schemas.ApartamentoWithId, *errs.AppError) {
	var apartamento schemas.ApartamentoWithId

	query, err := utils.LoadSQL("apartamento/find_by_id.sql")
	if err != nil {
		return nil, errs.Unexpected(fmt.Errorf("reading find apartamento by id SQL file: %w", err))
	}

	err = db.Connection.QueryRow(context.Background(), query, id.String()).Scan(&apartamento.Id, &apartamento.Numero, &apartamento.Bloco)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, errs.Unexpected(fmt.Errorf("querying apartamento: %w", err))
	}

	return &apartamento, nil
}
