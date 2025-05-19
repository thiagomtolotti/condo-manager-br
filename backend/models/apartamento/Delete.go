package apartmentModel

import (
	"backend/db"
	"context"

	"github.com/google/uuid"
)

func Delete(id uuid.UUID) (bool, error) {
	query := `DELETE FROM apartamentos WHERE id=$1`

	result, err := db.Connection.Exec(context.Background(), query, id)

	if err != nil {
		return false, err
	}

	return result.RowsAffected() != 0, nil
}
