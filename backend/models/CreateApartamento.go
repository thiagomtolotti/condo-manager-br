package models

import (
	"backend/db"
	"backend/schemas"
	"context"
	"fmt"
)

func CreateApartamento(apartamento schemas.Apartamento) error {
	if len(apartamento.Bloco) > 10 {
		return fmt.Errorf("apartment block must be max 10 characters long")
	}

	sql := `INSERT INTO apartamentos (numero, bloco) VALUES ($1, $2)`
	_, err := db.Connection.Exec(context.Background(), sql, apartamento.Numero, apartamento.Bloco)

	if err != nil {
		return fmt.Errorf("failed to insert apartamento: %w", err)
	}

	return nil
}
