package moradorModel

import (
	"backend/db"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"
)

func Get(page int, pageSize int) ([]schemas.Morador, error) {
	offset := (page - 1) * pageSize

	query, err := utils.LoadSQL("morador/list.sql")
	if err != nil {
		return []schemas.Morador{}, fmt.Errorf("error reading list moradores sql: %w", err)
	}

	rows, err := db.Connection.Query(context.Background(), query, pageSize, offset)

	if err != nil {
		return []schemas.Morador{}, err
	}
	defer rows.Close()

	var rowSlice []schemas.Morador = []schemas.Morador{}

	for rows.Next() {
		var r schemas.Morador

		err := rows.Scan(&r.Cpf, &r.Apartamento_id, &r.Nome, &r.Telefone, &r.Responsavel, &r.Proprietario)

		if err != nil {
			return []schemas.Morador{}, err
		}

		rowSlice = append(rowSlice, r)
	}

	return rowSlice, nil
}
