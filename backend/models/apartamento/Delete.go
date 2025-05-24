package apartmentoModel

import (
	"backend/db"
	"backend/utils"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("No apartment with the given id was found")

func Delete(id uuid.UUID) error {
	query, err := utils.LoadSQL("apartamento/delete.sql")
	if err != nil {
		return fmt.Errorf("error reading delete apartamento sql file: %w", err)
	}

	result, err := db.Connection.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() != 0 {
		return ErrNotFound
	}

	return nil
}
