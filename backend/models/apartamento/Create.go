package apartmentModel

import (
	"backend/db"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"
)

func CreateApartamento(apartamento schemas.Apartamento) error {
	// TODO: Validate length of blobo in controller
	if len(apartamento.Bloco) > 10 {
		return fmt.Errorf("apartment block must be max 10 characters long")
	}

	sql, err := utils.LoadSQL("apartamento/create.sql")
	if err != nil {
		return fmt.Errorf("error reading create apartamento sql file :%v", err)
	}

	_, err = db.Connection.Exec(context.Background(), sql, apartamento.Numero, apartamento.Bloco)

	if err != nil {
		return fmt.Errorf("failed to insert apartamento: %w", err)
	}

	return nil
}
