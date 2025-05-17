package apartmentModel

import (
	"backend/db"
	"backend/schemas"
	"context"
)

func GetApartamento(page int, pageSize int) ([]schemas.ApartamentoWithId, error) {
	offset := (page - 1) * pageSize

	sql := `SELECT * FROM apartamentos LIMIT $1 OFFSET $2 `

	rows, err := db.Connection.Query(context.Background(), sql, pageSize, offset)

	if err != nil {
		return []schemas.ApartamentoWithId{}, err
	}
	defer rows.Close()

	var rowSlice []schemas.ApartamentoWithId = []schemas.ApartamentoWithId{}

	for rows.Next() {
		var r schemas.ApartamentoWithId
		err := rows.Scan(&r.Id, &r.Numero, &r.Bloco)

		if err != nil {
			return []schemas.ApartamentoWithId{}, err
		}

		rowSlice = append(rowSlice, r)
	}

	return rowSlice, nil
}
