package moradorModel

import (
	"backend/db"
	"backend/utils"
	"backend/utils/cpf"
	"context"
	"fmt"
)

func Delete(cpf cpf.CPF) (bool, error) {
	query, err := utils.LoadSQL("morador/delete.sql")
	if err != nil {
		return false, fmt.Errorf("error reading delete morador sql file: %w", err)
	}

	result, err := db.Connection.Exec(context.Background(), query, cpf.Value)

	if err != nil {
		return false, err
	}

	return result.RowsAffected() != 0, nil
}
