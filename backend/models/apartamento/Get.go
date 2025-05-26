package apartmentoModel

import (
	"backend/db"
	"backend/errs"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"
)

func GetApartamento(page int, pageSize int) ([]schemas.ApartamentoWithId, *errs.AppError) {
	offset := (page - 1) * pageSize

	sql, err := utils.LoadSQL("apartamento/list.sql")
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading list apartamentos SQL file: %w", err))
		return []schemas.ApartamentoWithId{}, err
	}

	rows, err := db.Connection.Query(context.Background(), sql, pageSize, offset)
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("in list apartamentos query %w:", err))
		return []schemas.ApartamentoWithId{}, err
	}
	defer rows.Close()

	var rowSlice []schemas.ApartamentoWithId = []schemas.ApartamentoWithId{}
	for rows.Next() {
		var r schemas.ApartamentoWithId
		err := rows.Scan(&r.Id, &r.Numero, &r.Bloco)

		if err != nil {
			var err = errs.Unexpected(fmt.Errorf("reading row %v: %w", r.Id, err))
			return []schemas.ApartamentoWithId{}, err
		}

		rowSlice = append(rowSlice, r)
	}

	return rowSlice, nil
}
