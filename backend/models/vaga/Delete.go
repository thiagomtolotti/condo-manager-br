package vagaModel

import (
	"backend/db"
	"context"

	"github.com/google/uuid"
)

func Delete(id uuid.UUID) (bool, error) {
	const query = `DELETE FROM vagas WHERE id = $1`

	result, err := db.Connection.Exec(context.Background(), query, id)

	if err != nil {
		return false, err
	}

	return result.RowsAffected() != 0, nil
}
