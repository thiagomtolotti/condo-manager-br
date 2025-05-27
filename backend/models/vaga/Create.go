package vagaModel

import (
	"backend/db"
	"backend/errs"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func Create(apartamento_id uuid.UUID, body schemas.Vaga) (uuid.UUID, *errs.AppError) {
	var id uuid.UUID
	query, err := utils.LoadSQL("vaga/create.sql")

	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading create vaga SQL file: %w", err))
		return uuid.UUID{}, err
	}

	err = db.Connection.QueryRow(
		context.Background(),
		query,
		apartamento_id,
		body.Numero,
	).Scan(&id)
	if err != nil {
		return uuid.UUID{}, errs.Unexpected(fmt.Errorf("creating vaga in DB: %w", err))
	}

	return id, nil
}
