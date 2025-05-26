package apartmentoModel

import (
	"backend/db"
	"backend/errs"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func CreateApartamento(apartamento schemas.Apartamento) (uuid.UUID, *errs.AppError) {
	// TODO: Validate length of bloco in controller
	if len(apartamento.Bloco) > 10 {
		var err = errs.BadRequest("Apartamento block must be max 10 characters long", nil)
		return uuid.UUID{}, err
	}

	var id uuid.UUID

	sql, err := utils.LoadSQL("apartamento/create.sql")
	if err != nil {
		err = fmt.Errorf("loading create apartamento SQL file: %w", err)
		return uuid.UUID{}, errs.Unexpected(err)
	}

	err = db.Connection.QueryRow(context.Background(), sql, apartamento.Numero, apartamento.Bloco).Scan(&id)
	if err != nil {
		err = fmt.Errorf("Inserting new apartamento: %w", err)
		return uuid.UUID{}, errs.Unexpected(err)
	}

	return id, nil
}
