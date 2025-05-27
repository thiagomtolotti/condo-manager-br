package moradorModel

import (
	"backend/db"
	"backend/errs"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"
)

func Get(page int, pageSize int) ([]schemas.Morador, *errs.AppError) {
	offset := (page - 1) * pageSize

	query, err := utils.LoadSQL("morador/list.sql")
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading list moradores SQL file: %w", err))
		return []schemas.Morador{}, err
	}

	rows, err := db.Connection.Query(context.Background(), query, pageSize, offset)
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("querying list moradores: %w", err))
		return []schemas.Morador{}, err
	}
	defer rows.Close()

	var rowSlice []schemas.Morador = []schemas.Morador{}

	for rows.Next() {
		var r schemas.Morador

		err := rows.Scan(&r.Cpf, &r.Apartamento_id, &r.Nome, &r.Telefone, &r.Responsavel, &r.Proprietario)

		if err != nil {
			var err = errs.Unexpected(fmt.Errorf("Reading row %s: %w", r.Cpf, err))
			return []schemas.Morador{}, err
		}

		rowSlice = append(rowSlice, r)
	}

	return rowSlice, nil
}
