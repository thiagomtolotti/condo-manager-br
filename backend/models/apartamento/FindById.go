package apartmentoModel

import (
	"backend/db"
	"backend/errs"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func FindById(id uuid.UUID) (*schemas.ApartamentoWithId, *errs.AppError) {
	var apartamento schemas.ApartamentoWithId

	query, err := utils.LoadSQL("apartamento/find_by_id.sql")
	if err != nil {
		return nil, errs.Unexpected(fmt.Errorf("reading find apartamento by id SQL file: %w", err))
	}

	err = db.Connection.QueryRow(context.Background(), query, id).Scan(&apartamento)
	if err != nil {
		return nil, errs.Unexpected(fmt.Errorf("querying apartamento: %w", err))
	}

	return &apartamento, nil
}
