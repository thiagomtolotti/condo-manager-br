package apartmentModel

import (
	"backend/db"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func CreateApartamento(apartamento schemas.Apartamento) (uuid.UUID, error) {
	// TODO: Validate length of bloco in controller
	if len(apartamento.Bloco) > 10 {
		return uuid.UUID{}, fmt.Errorf("apartment block must be max 10 characters long")
	}

	var id uuid.UUID

	sql, err := utils.LoadSQL("apartamento/create.sql")
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error reading create apartamento sql file: %w", err)
	}

	err = db.Connection.QueryRow(context.Background(), sql, apartamento.Numero, apartamento.Bloco).Scan(&id)

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to insert apartamento: %w", err)
	}

	return id, nil
}
