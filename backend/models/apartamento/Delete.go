package apartmentModel

import (
	"backend/db"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func Delete(id uuid.UUID) error {
	query := `DELETE FROM apartamentos WHERE id=$1`

	result, err := db.Connection.Exec(context.Background(), query, id)

	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		// TODO: Return with status code for proper treatment
		return fmt.Errorf("apartamento not found")
	}

	return nil
}
