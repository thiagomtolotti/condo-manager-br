package vagaModel

import (
	"backend/db"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"
)

func Get(page int, pageSize int) ([]schemas.VagaWithApartment, error) {
	offset := (page - 1) * pageSize

	query, err := utils.LoadSQL("vaga/list.sql")
	if err != nil {
		return []schemas.VagaWithApartment{}, fmt.Errorf("error reading list vagas sql: %v", err)
	}

	rows, err := db.Connection.Query(context.Background(), query, pageSize, offset)

	if err != nil {
		return []schemas.VagaWithApartment{}, err
	}
	defer rows.Close()

	var rowsSlice []schemas.VagaWithApartment = []schemas.VagaWithApartment{}

	for rows.Next() {
		var row schemas.VagaWithApartment

		err := rows.Scan(&row.Id, &row.Numero, &row.Apartamento_id)

		if err != nil {
			return []schemas.VagaWithApartment{}, err
		}

		rowsSlice = append(rowsSlice, row)
	}

	return rowsSlice, nil
}
