package vagaModel

import (
	"backend/db"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func Create(apartamento_id uuid.UUID, body schemas.Vaga) (uuid.UUID, error) {
	var id uuid.UUID
	query, err := utils.LoadSQL("vaga/create.sql")

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error reading create vaga sql: %w", err)
	}

	err = db.Connection.QueryRow(
		context.Background(),
		query,
		apartamento_id,
		body.Numero,
	).Scan(&id)

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
