package apartmentModel

import (
	"backend/db"
	"backend/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func Delete(id uuid.UUID) (bool, error) {
	query, err := utils.LoadSQL("apartamento/delete.sql")
	if err != nil {
		return false, fmt.Errorf("error reading delete apartamento sql file: %v", err)
	}

	result, err := db.Connection.Exec(context.Background(), query, id)

	if err != nil {
		return false, err
	}

	return result.RowsAffected() != 0, nil
}
