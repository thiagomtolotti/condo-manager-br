package vagaModel

import (
	"backend/db"
	"backend/errs"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"
)

func Get(page int, pageSize int) ([]schemas.VagaWithApartment, *errs.AppError) {
	offset := (page - 1) * pageSize

	query, err := utils.LoadSQL("vaga/list.sql")
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading list vagas SQL file: %w", err))
		return []schemas.VagaWithApartment{}, err
	}

	rows, err := db.Connection.Query(context.Background(), query, pageSize, offset)
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("querying vagas: %w", err))
		return []schemas.VagaWithApartment{}, err
	}
	defer rows.Close()

	var rowsSlice []schemas.VagaWithApartment = []schemas.VagaWithApartment{}

	for rows.Next() {
		var row schemas.VagaWithApartment

		err := rows.Scan(&row.Id, &row.Numero, &row.Apartamento_id)

		if err != nil {
			var err = errs.Unexpected(fmt.Errorf("reading row with id %s: %w", row.Id, err))
			return []schemas.VagaWithApartment{}, err
		}

		rowsSlice = append(rowsSlice, row)
	}

	return rowsSlice, nil
}
