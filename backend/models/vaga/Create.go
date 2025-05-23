package vagaModel

import (
	"backend/db"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func Create(apartamento_id uuid.UUID, body schemas.Vaga) error {
	query, err := utils.LoadSQL("vaga/create.sql")
	if err != nil {
		return fmt.Errorf("error reading create vaga sql: %w", err)
	}

	_, err = db.Connection.Exec(
		context.Background(),
		query,
		apartamento_id,
		body.Numero,
	)

	if err != nil {
		return err
	}

	return nil
}
