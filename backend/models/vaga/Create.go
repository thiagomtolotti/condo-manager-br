package vagaModel

import (
	"backend/db"
	"backend/schemas"
	"context"

	"github.com/google/uuid"
)

func Create(apartamento_id uuid.UUID, body schemas.Vaga) error {
	const query = `
        INSERT INTO vagas (
            apartamento_id, numero
        ) VALUES ($1, $2)`

	_, err := db.Connection.Exec(
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
