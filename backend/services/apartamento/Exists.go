package apartamentoService

import (
	"backend/db"
	"context"

	"github.com/google/uuid"
)

func Exists(id uuid.UUID) (bool, error) {
	const query = `SELECT 1 FROM apartamentos WHERE id=$1`

	var exists int
	err := db.Connection.QueryRow(context.Background(), query, id).Scan(&exists)

	if exists == 0 {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
